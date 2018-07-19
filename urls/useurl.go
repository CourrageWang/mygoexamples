package main

import (
	"net/url"
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
)

func main() {
	//Url1()
	useGet()
}

func Url1() {

	u, err := url.Parse("https://blog.csdn.net/wangshubo1989/article/details/75017632")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
	fmt.Println(u.User.Username())
	fmt.Println(u.Scheme)
	fmt.Println(u.Opaque)
	fmt.Println(u.Host)
	fmt.Println(u.Path)
	fmt.Println(u.RawPath)
	fmt.Println("u.Fragment=", u.Fragment)
}
func useGet() {
	rep, err := http.Get("https://www.cnblogs.com/jiangzhaowei/p/7881115.html")
	if err != nil {
		os.Exit(1)
	}
	b, err2 := ioutil.ReadAll(rep.Body)
	rep.Body.Close()
	if err2 != nil {
		os.Exit(1)
	}
	fmt.Printf("%s",b)

}
