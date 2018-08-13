package main

/**
  默认的信道（无缓存的信道）的存取消息都是阻塞的，其会在存取消息时挂起当前的线程。

 */
var ch chan int = make(chan int)

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
	c, qut := make(chan int), make(chan int)
	go func() {
		c <- 1 // func 被c阻塞 （c等待数据输出）
		qut <- 0
	}()
	<-qut // 主线程中qut等大数据写入

}
func mychan() { // deadlock
	ch = make(chan int)
	<-ch //  阻塞当前线程 ，信道被锁
}
