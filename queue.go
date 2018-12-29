package main

import (
	"time"
	"fmt"
	"os"
)

/**
  golang 任务队列
 */

var done chan bool = make(chan bool, 100)

func worker2(done chan bool) {
	time.Sleep(time.Millisecond * 500)
	fmt.Println("done the task...")
	f, _ := os.Create("chanen.txt")
	f.WriteString("I:")
	//通知任务已经完成
	done <- true
}
func tetstchannl(done chan bool)  {
	worker2(done)
}
func main() {
	/*//done:=make( chan bool , 1 )
	go worker2(done)
	//time.Sleep(time.Second*1)
	go worker2(done)*/

	for i := 0 ;i< 1000  ;i++  {
		go tetstchannl(done)
	}
	<-done
	//go worker(done)
	//status,_ :=<-done
	//if status {
	//	fmt.Println("has done ....")
	//}

}
