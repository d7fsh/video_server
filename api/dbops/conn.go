package dbops

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"lottery_project/conf"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	sourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
		conf.DBMaster.Host,
		conf.DBMaster.Pwd,
		conf.DBMaster.Host,
		conf.DBMaster.Port,
		conf.DBMaster.Database)

	dbConn, err = sql.Open(conf.DriverName, sourceName)
	if err != nil {
		log.Println("创建数据库连接失败, 请检查: ", err)
		return
	}
	//defer dbConn.Close()
}
