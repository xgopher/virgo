package database

import (
	_ "github.com/go-sql-driver/mysql" // 导入 mysql 驱动
	"github.com/jinzhu/gorm"
	"os"
)

var DB *gorm.DB // db pool instance
var err error   // db err instance

func InitPool() {
	// Openning file
	conn := os.Getenv("MYSQL_CONN")
	DB, err = gorm.Open("mysql", conn)
	DB.LogMode(true)
	// Error
	if err != nil {
		panic(err)
	}

	DB.DB().SetMaxIdleConns(10)  // Golang 原生方法: 设置闲置的连接数
	DB.DB().SetMaxOpenConns(100) // Golang 原生方法: 最大打开的连接数，默认值为0表示不限制
	DB.DB().Ping()               // Golang 原生方法: ping

	// 表前辍
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "" + defaultTableName
	}
}
