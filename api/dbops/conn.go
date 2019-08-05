package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:MyNewPass4!@tcp(127.0.0.1:3306)/video_server?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		log.Println("mysql conn err: ", err)
		return
	}
	//defer dbConn.Close()
}
