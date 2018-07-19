package main

import (
	"github.com/astaxie/beego/config"
	"fmt"
)

const sourcepath = "/root/go/src/conf/conf.ini" // 文件路径
const localpath = "conf.ini"

/*
[demo]
key1 = "asta"
key2 = "xie"

可以通过 iniconf.String("demo::key2") 获取值





 */
func main() {
	//iniconf, err := config.NewConfig("ini", localpath)
	//if err != nil {
	//	log.Print("error...", err)
	//}
	//fmt.Println(iniconf.String("syslog::user")) // 从配置文件中读取
	////iniconf.Set("passWord", strconv.Itoa(23456))
	////fmt.Println(iniconf.String("passWord"))
	////iniconf.SaveConfigFile("conf.ini") // 存储到配置文件

	fmt.Println(getDBConf(localpath))

}

//user:password@tcp(127.0.0.1:3306)/hello
func getDBConf(path string) (string, error) {
	iniconf, err := config.NewConfig("ini", path)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s@%s(%s:%s)/%s", iniconf.String("syslog::user"), iniconf.String("syslog::password"), iniconf.String("syslog::protocol"),
		iniconf.String("syslog::ip"), iniconf.String("syslog::port"), iniconf.String("syslog::databaseName")), nil

}
