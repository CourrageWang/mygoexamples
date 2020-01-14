package main

import (
	"github.com/astaxie/beego/logs"
)

/***
  日志持久化
【 日志能够持久化 】
*/
func main() {
	log := logs.NewLogger(10000) // 创建一个日志记录器，参数为缓冲区的大小
	// 设置配置文件
	/*logConfig := `{
		   "filename" : "project.log",
		   "maxlines" : 0,
		   "maxsize"  : 0 ,
		   "daily"    :true,
		   "maxdays"   :3,
	       "level":7,
	       "color":true
		}`*/
	//err := log.SetLogger(logs.AdapterFile, logConfig) // 设置日志记录方式：本地文件记录
	err := log.SetLogger(logs.AdapterFile, `{
	   "filename" : "project.log",
	   "maxlines" : 0,
	   "maxsize"  : 0 ,
	   "daily"    :true,
	   "maxdays"   :3,
       "level":7,
       "color":true
	}`)                                // 设置日志记录方式：本地文件记录
	log.SetLogger(logs.AdapterConsole) // 设置日志记录方式：控制台记录
	if err != nil {
		log.Error("设置日期时出现异常", err)
	}
	log.SetLevel(logs.LevelDebug) // 设置日志写入缓冲区的等级
	log.EnableFuncCallDepth(true) // 输出log时能显示输出文件名和行号（非必须）

	log.Info("------开始程序--------")
	log.Debug("Ip地址为%s", "192.168.0.125")
	log.Debug("try to connect the sql:%s", "127.0.0.1")
	log.Error("error  has happend %s", "Ip Address has not correct")

	log.Flush() // 将日志从缓冲区读出，写入到文件
	log.Close()
}

/**
golang 日志框架beego logs 本地输出；
*/
func main2() {
	log := logs.NewLogger(10000)  // 创建一个日志记录器，参数为缓冲区的大小
	log.SetLogger("console", "")  // 设置日志记录方式：控制台记录
	log.SetLevel(logs.LevelDebug) // 设置日志写入缓冲区的等级：Debug级别（最低级别，所以所有log都会输入到缓冲区）
	log.EnableFuncCallDepth(true) // 输出log时能显示输出文件名和行号（非必须）

	log.Emergency("Emergency")
	log.Alert("Alert")
	log.Critical("Critical")
	log.Error("Error")
	log.Warning("Warning")
	log.Notice("Notice")
	log.Informational("Informational")
	log.Debug("Debug")
	log.Close()
}