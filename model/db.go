package model

import (
	"database/sql"
	"log"
)

/**
  规划数据裤的链接方式
 1 、全局变量  简单直接 设置一个指向数据库链接池的指针作为全局变量。
 2、使用初始化函数，数据库连接服务可以在其他包中完成。
 */

var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Print(err)
	}
	if err := db.Ping(); err != nil {
		log.Print(err)
	}
}
