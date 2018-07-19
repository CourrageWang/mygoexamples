package main

import (
	"fmt"
	"sort"
)

func main() {
	//testmap1()
	orderedRangeMap()
}

// map 原生range不保证有序
func testmap1() {
	rmap := make(map[int]interface{})
	rmap[0] = "this"
	rmap[1] = "is"
	rmap[2] = "test"
	for _, v := range rmap {
		fmt.Println(v) // test this is
	}
}

/**
if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		prefix := QueryEscape(k) + "="
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(prefix)
			buf.WriteString(QueryEscape(v))
		}
	}
	return buf.String()
 */
// map 有序
func orderedRangeMap() {

	rmap := make(map[string]interface{})
	rmap["A"] = "this"
	rmap["C"] = "test"
	rmap["B"] = "is"
	var keys []string
	for k, _ := range rmap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(rmap[k])
	}
	// slice
	s := make([]string, 5)
	s = append(append(append(s, "str"), "test"), "!!")
	fmt.Println(s)




}
