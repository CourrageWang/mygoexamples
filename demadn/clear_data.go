package main
//
//import (
//	"fmt"
//	"net/http"
//	"time"
//)
//
//func HelloResponse(rw http.ResponseWriter, request *http.Request) {
//
//	fmt.Println(11111)
//	fmt.Fprintf(rw, "OK !")
//}
//
//func main() {
//
//	//
//	//
//	//主线程不阻塞
//	timer2 := time.NewTimer(time.Second)
//
//	go func() {
//
//		//等触发时的信号
//
//		<-timer2.C
//
//		fmt.Println("Timer 2 expired")
//
//	}()
//
//	//由于上面的等待信号是在新线程中，所以代码会继续往下执行，停掉计时器
//
//	time.Sleep(time.Second*5)
//
//	http.HandleFunc("/", HelloResponse)
//	http.ListenAndServe(":3000", nil)
//}
