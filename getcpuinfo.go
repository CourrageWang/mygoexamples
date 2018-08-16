package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
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
	for i := 0; i < 300; i++ {

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
func Heapfile() {
	Heap,err := os.Create("theap.prof")
	if err!=nil {
		return
	}
	defer  Heap.Close()
	pprof.WriteHeapProfile(Heap)
}

