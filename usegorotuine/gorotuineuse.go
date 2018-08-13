package main

import (
	"runtime"
	"sync"
	"fmt"
	"sync/atomic"
)

var (
	counter int64          //所有协程需要修改的变量
	wg      sync.WaitGroup //等待线程结束
)
//gorotuine
func grontuine01() {
	runtime.GOMAXPROCS(1) // 分配一个逻辑处理器
	var wg sync.WaitGroup // 用来等待程序完成
	wg.Add(2)             // 计数器加2，表示等待两个协程
	fmt.Printf("start gorotuines...")
	go func() {
		defer wg.Done() // 函数退出时调用Done 通知main线程 工作已经完成。
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	go func() {
		defer wg.Done() // 函数退出时调用Done 通知main线程 工作已经完成。
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()
	wg.Wait() //等待gorotuine结束
}

// 竞争条件
func gorotuine02() {
	wg.Add(2)
	//go incCouter(1)
	//go incCouter(2)
	// 测试原子操作
	go useactomic(1)
	go useactomic(2)

	wg.Wait()
	fmt.Println("final counter is :", counter)

}
func incCouter(id int) { // 2  一个gorotuine 覆盖另一个的值。
	defer wg.Done()
	for count := 0; count < 2; count++ {
		value := counter
		runtime.Gosched() // 当前grotuine 从线程退出，并返回队列
		value ++          // 增加本地value变量的值
		counter = value   //将该值保存回counter

	}
}

/*  对于竞态资源 可以使用原子操作。
    原子函数能够已很底层的加锁机制来同步访问整型变量和指针
 */
func useactomic(id int) { // 4
	defer wg.Done()
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1) // 安全的为counter增加1
		runtime.Gosched()            // 当前grotuine 从线程退出，并返回队列

	}
}

/*  使用互斥锁
   保证同时只有一个gorotuine

 */

func main() {
	//grontuine01()
	gorotuine02()
}
