package conf

// 此文件用于配置mysql数据库

const Drivename = "mysql"

type DBConfig struct {
	Host      string
	Port      int
	User      string
	Pwd       string
	Database  string
	IsRunning bool
}

var DBMasterList = []DBConfig{
	{
		Host:      "127.0.0.1",
		Port:      3306,
		User:      "root",
		Pwd:       "I&My_Dogs",
		Database:  "video_server",
		IsRunning: true,
	},
}

// 获取切片索引位置0的配置信息
var DBMaster DBConfig = DBMasterList[0]
