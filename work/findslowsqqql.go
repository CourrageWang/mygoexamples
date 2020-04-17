package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type queryInfo struct {
	queryUseTime   string
	queryDB        string
	queryTime      string
	queryUser      string
	queryStatement string
}

// 用户查询结构体
const (
	queryDB      string  = "DB"
	queryUser    string  = "User"
	queryUseTime string  = "Query_time"
	queryTime    string  = "Time"
	queryEndFlag string  = "Succ"
	timeOut      float64 = 0.00009 //  测试用 后期读配置文件
)

func initQueryInfo(info *queryInfo) {
	info.queryStatement = ""
	info.queryDB = ""
	info.queryUseTime = ""
	info.queryUser = ""
	info.queryTime = ""
}

func getSlowQueryInfo(path string) {
	f, err := os.Open(path)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 建立缓冲区，把文件内容存放到缓冲区
	rd := bufio.NewScanner(f)

	// 定义一个全局的map保存一个周期的结果
	var queryInfo queryInfo

	for rd.Scan() {
		line := rd.Text()

		/**
		按行读取，获取有效信息，当读到sUcc则完成一个块的读取。
		*/
		if strings.Contains(line, "#") {
			if strings.Contains(line, queryUseTime) {
				queryInfo.queryUseTime = line[len("# Query_time: "):]
			} else if strings.Contains(line, queryDB) {
				queryInfo.queryDB = line[len("# DB: "):]
			} else if strings.Contains(line, queryTime) {
				queryInfo.queryTime = line[len("# Time: "):]
			} else if strings.Contains(line, queryUser) {
				queryInfo.queryUser = line[len("# User: "):]
			}
		} else {
			queryInfo.queryStatement = line
		}

		if strings.Contains(line, queryEndFlag) {

			if queryInfo.queryStatement == "" { //去除掉sql语句为空的语句
				initQueryInfo(&queryInfo)
			} else {
				handleQueryInfo(&queryInfo)
			}
		}
	}
}

/**
  根据摘取的sql查询信息中的时间 判断时间是否超时，超时将做后续处理；
*/
func handleQueryInfo(info *queryInfo) {
	useTime, err := strconv.ParseFloat(strings.TrimSpace(info.queryUseTime), 64)
	if err!=nil {
		panic(err)
	}
	// 过滤出超时的查询信息
	if useTime>timeOut {
		fmt.Println(info)
	}
}

func main() {

	path := "/Users/yqwang/Workspace/gopath/src/github.com/CourrageWang/mygoexamples/work/a.log"
	getSlowQueryInfo(path)
}
