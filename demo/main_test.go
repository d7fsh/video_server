package demo

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

// 1. 创建文件时, 文件名要以_test结尾
// 2. 在此时方法签名加上Test关键字
// 3. 方法中必须传递t *testing.T
func TestDBConnection(t *testing.T) {
	// 1. 尝试连接数据库
	dbConn, err := sql.Open("mysql", "root:I&My_Dog7@tcp(127.0.0.1:3306)/video_server?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	defer dbConn.Close()
	err = dbConn.Ping()
	if err != nil {
		panic(err.Error())
	}
	stmt, err := dbConn.Prepare("INSERT INTO users (login_name,pwd) VALUES (?,?);")
	defer stmt.Close()
	if err != nil {
		fmt.Println(err)
	}
	//3，判断插入结果
	_, err = stmt.Exec("tom", "12345")
	if err != nil {
		fmt.Println(err)
	}
}
