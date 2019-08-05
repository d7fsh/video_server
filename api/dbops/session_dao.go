package dbops

import (
	"database/sql"
	"fmt"
	"strconv"
	"sync"
	"video_server/api/defs"
)

// 1. 向session表中插入数据
func InsertSession(sid string, ttl int64, loginName string) error {
	// 1. 将ttl由int64转换为十进制, 然后将这个10进制数字转换为字符串
	ttlStr := strconv.FormatInt(ttl, 10)
	// 2. 通过dbConn准备执行sql语句, 插入数据库
	stmt, err := dbConn.Prepare("INSERT INTO sessions (session_id,ttl,login_name) VALUES (?,?,?)")
	defer stmt.Close()
	if err != nil {
		return err
	}
	// 3. 判断插入结果
	_, err = stmt.Exec(sid, ttlStr, loginName)
	if err != nil {
		// 3.2 插入失败, 返回err
		return err
	}
	// 3.1 插入成功, 返回nil
	return nil
}

// 2, 删除session表中数据
func DeleteSession(sid string) error {
	stmt, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id = ?")
	defer stmt.Close()
	if err != nil {
		return err
	}
	// 执行删除操作
	_, err = stmt.Query(sid)
	if err != nil {
		return err
	}
	return nil
}

// 3. 根据session_id查询session表中的数据, 将查到的数据, 用SimpleSession存储
func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmt, err := dbConn.Prepare("SELECT ttl,login_name from sessions where session_id = ?")
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	var ttl string   // 字符串, 记录查询得到的时间戳
	var uname string // 用户名, 记录查询得到的用户名

	// 根据sid参数的值, 触发sql语句
	err = stmt.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	// ttl 转换成int64, res int64

	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = res
		ss.Username = uname

	} else {
		return nil, err
	}
	return ss, nil
}

// 4. 查询session表中所有的数据
func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmt, err := dbConn.Prepare("SELECT * FROM sesssions")
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var (
			id         string
			ttlStr     string
			login_name string
		)
		if err := rows.Scan(&id, &ttlStr, &login_name); err != nil {
			fmt.Println(err)
			break
		}

		if ttl, err := strconv.ParseInt(ttlStr, 10, 64); err == nil {
			ss := &defs.SimpleSession{Username: login_name, TTL: ttl}
			// 将ss根据唯一性的标识, 存储在m中
			m.Store(id, ss)
		}
	}
	return m, nil
}
