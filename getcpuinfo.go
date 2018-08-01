package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
	"flag"
	"runtime"
)
//https://cloud.tencent.com/developer/section/1144254
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	log.Println("begin")
	flag.Parse()

		f, err := os.Create("cpuprofile.prof")
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()

	//heapProfile("cpu1.prof")
	for i := 0; i < 30; i++ {

		nums := fibonacci2(i)
		//heapProfile("cpu1.prof")
		fmt.Println(nums)
	}

		f2, err2 := os.Create("mem.prof")
		if err2 != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f2.Close()


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
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

}
