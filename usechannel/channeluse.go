package main

import (
	"fmt"
)
/**
  默认的信道（无缓存的信道）的存取消息都是阻塞的，其会在存取消息时挂起当前的线程。
  go 启动的所有的gorotuine里面的非缓冲县信道，一个一定要存， 一个一定要取要成对出现。
 =======================缓冲信道====================================
 缓冲信道不仅可以流通数据还可以缓存数据，当缓冲信道到达满的状态就会阻塞。
 换从信道是一个线程安全的队列。
  解决死锁问题：
1、把没取走的数据取走， 没放入数据的放入数据，因为无缓冲信道不能承载数据。
2、使用缓冲信道。放入数据后不会挂起当前协程，只有当缓存的数量到达上线后才会阻塞。
 */
var ch chan int = make(chan int)

var quit chan int

func too() {
	ch <- 2 //  存消息
}
func main() {

	/*data := make(chan int)
	exit := make(chan bool)
	go func() {
		for d := range data {
			fmt.Println(d)
		}
		fmt.Println("recv over .")
		exit <- true // 发送退出通知
	}()
	data <- 1
	data <- 2
	data <- 3
	close(data) // 关闭队列
	fmt.Println("send over .")
	<-exit*/

	/*go too()
	fmt.Println(<-ch)*/

	//mychan()

	/*c, qut := make(chan int), make(chan int)
	go func() {
		c <- 1 // func 被c阻塞 （c等待数据输出）
		qut <- 0
	}()
	<-qut // 主线程中qut等大数据写入*/

	//  缓冲信道
	/*ch := make(chan string, 3)
	ch <- "h"
	ch <- "e"
	ch <- "l"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)*/

	//======================信道数据读取和关闭============================
	ch2 := make(chan string, 3)
	ch2 <- "h"
	ch2 <- "e"
	ch2 <- "l"
	close(ch2) //显示关闭信道
	for str := range ch2 { //  range 不等到信道关闭是不会技术读取的，当信道干涸了， range就会阻塞当前的gorituine
		fmt.Println(str)
	}
	//  ======================等待多个gorotuine的方案==============================
	/**  1、使用单个无线缓冲信道阻塞程序
	     2、使用容量为goroutines数量的缓冲信道
	 */
	// 1
	count := 1000
	quit = make(chan int)
	for i := 0; i < count; i++ {
		go foo(i)
	}
	for i := 0; i < count; i++ {
		<-quit
	}

	// 2
	quit = make(chan int, 1000)
	for i := 0; i < count; i++ {
		go foo(i)
	}
	// 区别 ：无缓冲的信道是一批数据一个一个操作的（一个进一个出）
	//  缓冲则是一个一个存储然后一起流出。

}
func mychan() { // deadlock
	ch = make(chan int)
	<-ch //  阻塞当前线程 ，信道被锁
}

func foo(id int) {
	fmt.Println(id)
	quit <- 0 // 关闭通知
}
