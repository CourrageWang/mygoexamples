package main

import (
	"net/http"
	"strings"
	"fmt"
	"io/ioutil"
	"strconv"
	"net/url"
	"time"
	"net"
	"encoding/json"
)

/* =========================golanghttp请求方式================================
    https://studygolang.com/articles/4489
    https://blog.csdn.net/kenkao/article/details/47857757
   http 发送方式
 一、使用http.Newrequest()
 先生成http.client()->再生成http.request()之后提交请求；client.Do(request)->处理返回结果
 每一步的过程都可以甚至一些具体参数
 */
/*
	管理HTTP  客户端的头域，重定向策略和其他设置
        client := &http.Client{
	     CheckRedirect: redirectPolicyFunc,
      }
        resp, err := client.Get("http://example.com")
  // ...
     req, err := http.NewRequest("GET", "http://example.com", nil)
// ...
    req.Header.Add("If-None-Match", `W/"wyzzy"`)
     resp, err := client.Do(req)


要管理代理、TSL配置 ，keepalive 压缩和其他设置，创建一个Transport
tr := &http.Transport{
	TLSClientConfig:    &tls.Config{RootCAs: pool},
	DisableCompression: true,
}
client := &http.Client{Transport: tr}
resp, err := client.Get("https://example.com")
 */

var netTransport = &http.Transport{
	Dial: (&net.Dialer{
		Timeout:   10 * time.Second,
		KeepAlive: 30 * time.Second,
	}).Dial,
	TLSHandshakeTimeout:   5 * time.Second,
	ResponseHeaderTimeout: 10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

func httpClientTestGET() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			req.URL.String()
			return nil
		},
		Transport: netTransport,
	}
	// 生成要访问的url
	urls := "http://wthrcdn.etouch.cn/weather_mini?city=西安市"
	// 提交请求
	request, err := http.NewRequest("GET", urls, nil)
	if err != nil {
		fmt.Println("request failed ...", err)
	}
	resp, er := client.Do(request) // 处理返回的结果
	if er != nil {
		fmt.Println("Do ...", er)
	}
	defer resp.Body.Close()
	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Copy failed ...", err)
	}
	length := len(bodyByte)
	fmt.Println(string(bodyByte[:length]))
	var data map[string]interface{}
	json.Unmarshal(bodyByte, &data)

}

func httpClientTestPOST() {
	client := &http.Client{}
	// 生成要访问的url
	urls := "http://wthrcdn.etouch.cn/weather_mini/"
	params := make(url.Values) // 封装请求数据
	params.Add("name", "jackma")
	params.Add("age", "12")
	data := params.Encode()
	// 提交POST请求
	request, err := http.NewRequest("POST", urls, strings.NewReader(data))
	if err != nil {
		fmt.Println("request failed ...", err)
	}
	resp, er := client.Do(request) // 处理返回的结果
	if er != nil {
		fmt.Println("Do ...", er)
	}
	defer resp.Body.Close()
	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Copy failed ...", err)
	}
	length := len(bodyByte)
	fmt.Println(string(bodyByte[:length]))
	if status, _ := strconv.Atoi(resp.Status); status == 200 {
		fmt.Println("请求成功。。。")
	}

}

// http client Do
func clientDoUse() {
	client := &http.Client{}
	reqest, _ := http.NewRequest("GET", "http://www.baidu.com", nil)

	reqest.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	reqest.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	reqest.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	reqest.Header.Set("Cache-Control", "max-age=0")
	reqest.Header.Set("Connection", "keep-alive")
	response, _ := client.Do(reqest)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body);
		var data map[string]interface{}
		json.Unmarshal([]byte(bodystr), &data)

		fmt.Println("data.", data)
	}
}

func main() {
	//clientDoUse()
	httpClientTestGET()
	//httpDo()
}
