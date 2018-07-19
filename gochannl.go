package main

import (
	"fmt"
	"time"
)

func sum(s[] int, c chan int )  {
	sum := 0
	for _,v:=range s{
       sum  +=v
	}
	c <-sum
}
func main1()  {
	s:=[]int{2,4,6,8,9,-1,2}
	c:=make(chan int)// no buffer
	go sum(s[:len(s)/2],c)
	go sum(s[len(s)/2:],c)
	x,y:=<-c,<-c
	fmt.Println(x,y,x+y)
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func main2() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func main3()  {
 c:= make(chan int ,2)// has buffer
 //defer close(c)
 c<-1
 c<-2
 i ,ok :=<-c
 fmt.Println(i , ok)

}
// worker 完成任务后发送数据通知到channel通知main gorotuine 完成任务
func worker(done chan bool){
	time.Sleep(time.Second*2)
	fmt.Println("done the task...")
	//通知任务已经完成
	done<- true
}
func main()  {
  done:=make( chan bool , 1 )
  go worker(done)
  status,_ :=<-done
	if status {
		fmt.Println("has done ....")
	}



}