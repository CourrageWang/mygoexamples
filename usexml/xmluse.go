package main

import (
	"io/ioutil"

 ft "fmt"
"encoding/xml"
)
type XLReceive struct {
	Description string
	ServerPort  string
	ApiServerIp string
}

//
type PAReceive struct {
	Description string
	ServerPort  int
	ApiServerIp string
}

func main() {

	getXMlInfo()
}
func getXMlInfo() {
	content, err := ioutil.ReadFile("golang.xml")
	if err != nil {
		ft.Println("read file faiel ", err)
	}
	var result XLReceive
	err2 := xml.Unmarshal(content, &result)
	if err2 != nil {
		fmt.Println(err)
	}
	fmt.Println(result.ApiServerIp)

	//var paReceive PAReceive
	paReceive := &PAReceive{
		Description: "这是测试文件",
		ServerPort:  1280,
		ApiServerIp: "192.168.0.122",
	}

	//生成xml并追加
	data, err3 := xml.MarshalIndent(&paReceive, "", "  ")
	if err3 != nil {
		fmt.Println("marshal failed .. ", err3)
	}
	fmt.Println(string(data[:len(data)]))
	//ioutil.WriteFile("golang.xml", data, 0644)

}
