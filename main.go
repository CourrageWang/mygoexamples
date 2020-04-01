package main

import (
	"fmt"
	"github.com/CourrageWang/mygoexamples/mode"
	"github.com/CourrageWang/mygoexamples/util"
	"net/http"
	"strings"
	"sync"
	"time"
)

func HelloResponse(rw http.ResponseWriter, request *http.Request) {
	fmt.Println(11111)
	fmt.Fprintf(rw, "OK !")
}

func main() {

	cleanStatus := util.Conf.DUMP.Status
	fmt.Println("清除转状态为：", cleanStatus)

	/**
	  定时任务 清除数据
	*/

	var wg sync.WaitGroup
	wg.Add(1)

	ticker1 := time.NewTicker(5 * time.Second)

	go func(t *time.Ticker) {
		defer wg.Done()
		for {
			<-t.C
			// 清除开关是否打开
			if cleanStatus != "" && strings.Compare("ON", cleanStatus) == 0 {
				// 获取清除日期间隔
				durationTime := util.Conf.DUMP.Space

				fmt.Println("space :------>", durationTime)

				fmt.Println("db info is:", util.Conf.DB.OLDTiDB_DSN)

				// 初始化数据库信息
				mode.InitoldDB(util.Conf.DB.OLDTiDB_DSN)
				mode.InitnewDB(util.Conf.DB.TiDBDSN)

				mode.CleanParks15DaysNotBindARM(durationTime)
				mode.CleanGroup15DaysNotBindParks(durationTime)
			} else {
				fmt.Println("清除数据功能暂未开启，请重新确认！！")
			}

		}
	}(ticker1)

	http.HandleFunc("/", HelloResponse)
	http.ListenAndServe(":3000", nil)

	wg.Wait()
}
