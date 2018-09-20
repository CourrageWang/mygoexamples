package main

import (
	"crypto/tls"
	"net/http"
	"net/http/cookiejar"
	"fmt"
	"io/ioutil"
	"strings"
	"net/url"
)

func main() {

	//跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//http cookie接口
	cookieJar, _ := cookiejar.New(nil)

	c := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}
       //  https  get 请求
		str := "https://xxx.xxx.x.xxx:xxxx/api/v1/group/subscribe?access_token=.............&group_id=用户组"
		response, err := c.Get(str)
		if err != nil {
			fmt.Println("has error", err)
		}
		defer response.Body.Close()
		body, err2 := ioutil.ReadAll(response.Body)

		if err2 != nil {
			fmt.Println("has error", err2)
		}
		fmt.Println(string(body))


	/**
	     golang https post 请求
	 */
	postUrl := "https://xxx.xxx.x.xxx:xxxx/api/v1/groups"
	rebody := "application/x-www-form-urlencoded"
	reqm := make(url.Values)
	reqm.Add("access_token", "..................");
	reqm.Add("name", "xian")
	reqm.Add("parent_name", "srun")
	encode := reqm.Encode()
	response2, err3 := c.Post(postUrl, rebody, strings.NewReader(encode))
	if err != nil {
		fmt.Println("has error", err3)
	}
	defer response2.Body.Close()
	body, err4 := ioutil.ReadAll(response.Body)

	if err2 != nil {
		fmt.Println("has error", err4)
	}
	fmt.Println(string(body))

	//  https  del请求





}
