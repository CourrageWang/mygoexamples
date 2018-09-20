package gotime

import (
	"time"
	"fmt"

)

/**
   go lang 时间工具类
 */


 // golang 获取seconds秒后的日期

 func GetSecondsDate( seconds string)  {

	now := time.Now()
	//da :=now.Add(seconds*time.Second)
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	t,_:=time.ParseDuration(seconds+"s")
	fmt.Println(t)
	new:= now.Add(t)
	fmt.Println(new.Format("2006-01-02 15:04:05"))



}
