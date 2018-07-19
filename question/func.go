package main

import (
	"fmt"
	"strconv"
	"time"
	"runtime"
	"runtime/debug"
	"log"
)

/**
 golang函数调用参数均为值传递，不是指针传递或者引用传递
 */
func main1() {
	m := make(map[int]string)
	m[0] = "test"
	fmt.Println(testmap(m))
	fmt.Println(m)

}
func testmap(m map[int]string) string {
	 m[0]="a"
	return m[0]
}


func modify(s []int) {
	fmt.Printf("%p \n", &s)
	s = []int{1,1,1,1}
	fmt.Println(s)
	fmt.Printf("%p \n", &s)
}

func main() {
	/*a := [5]int{1, 2, 3, 4, 5}
	s := a[:]
	fmt.Printf("%p \n", &s)
	modify(s)
	fmt.Println(s[3])*/
	TestGC()
}

func TestGC() {
	ss := make([]string, 100<<20)
	for i := range ss {
		ss[i] = strconv.Itoa(i)
	}
	log.Println("gc before1")
	time.Sleep(2 * time.Second)
	log.Println("gc before2")
	runtime.GC()
	debug.FreeOSMemory()
	log.Println("gc after1")
	ref := &ss[2]
	runtime.GC()
	debug.FreeOSMemory()
	log.Println("gc after2")
	time.Sleep(5 * time.Second)
	log.Printf("ref: %p", ref)
}
