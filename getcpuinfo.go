package main

import (



	"fmt"

	"log"

	"os"

	"runtime/pprof"

	"flag"
	"time"
)



var (

	//定义外部输入文件名字

	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file.")




)



func main() {

	log.Println("begin")

	flag.Parse()



	if *cpuprofile != "" {

		f, err := os.Create(*cpuprofile)

		if err != nil {

			log.Fatal(err)

		}



		pprof.StartCPUProfile(f)

		defer pprof.StopCPUProfile()

	}



	for i := 0; i < 30; i++ {

		nums := fibonacci2(i)
		heapProfile("cpu1.prof")

		fmt.Println(nums)

	}
	heapProfile("cpu1.prof")
	//cpuprofile()


}

//递归实现的斐波纳契数列

func fibonacci2(num int) int {

	if num < 2 {

		return 1

	}

	return fibonacci2(num-1) + fibonacci2(num-2)

}
func cpuProfile() {
	f, err := os.OpenFile("cpu.prof", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	log.Println("CPU Profile started")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	time.Sleep(60 * time.Second)
	fmt.Println("CPU Profile stopped")
}

// 生成堆内存报告
func heapProfile(path string) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	time.Sleep(30 * time.Second)

	pprof.WriteHeapProfile(f)
	fmt.Println("Heap Profile generated")
}