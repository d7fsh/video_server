package session

import (
	"sync"
	"time"
	"video_server/api/dbops"
	"video_server/api/defs"
	"video_server/api/util"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

// 1. 生成一条session用于记录用户的登录状态
// 1.1 获取用户名
func GenerateNewSessionId(username string) string {
	// 1.2 指定时间戳(session有效时长) 拿到执行这行代码的当前时间+预设的有效时长
	currentTime := nowInMilli()
	ttl := currentTime + 30*60*1000
	// 1.3 session_id需要随机生成
	id, _ := util.NewUUID()

	session := &defs.SimpleSession{Username: username, TTL: ttl}
	// 1.4 将session存储到内存中
	sessionMap.Store(id, session)
	// 1.5 将session插入到数据库中
	dbops.InsertSession(id, ttl, username)
	return id
}

func nowInMilli() int64 {
	return time.Now().UnixNano() / 100000
}

// 2. 删除session
func DeleteExpiredSession(sid string) {
	// 2.1 从内存的map中删除session
	sessionMap.Delete(sid)
	// 2.2 从数据库中删除session
	dbops.DeleteSession(sid)
}

// 3. 判断session是否过期
// 返回值一: 有效session指向的用户名称
// 返回值二: true(过期), false(没过期, 可以)
func IsSessionExpired(sid string) (string, bool) {
	// 3.1 从内存中读取session(为了让内存中的session和数据库中的session保持同步)
	ss, ok := sessionMap.Load(sid) // 返回interface{}类型, ss需要类型断言
	if ok {
		// 3.2 拿ss中存储的ttl和此代码的时间戳进行比较
		ct := nowInMilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			// ttl< 当前读取时间(session过期), 内存中删除session, 数据库中删除session
			DeleteExpiredSession(sid)
			return "", true

		}
		// ttl> 当前读取时间(session有效), 返回用户名和false
		return ss.(*defs.SimpleSession).Username, false
	}
	return "", true
}

// 4. 从数据库中读取所有session存储到map中
func LoadSessionFromDB() {
	sessions, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}
	sessionMap = sessions
}
