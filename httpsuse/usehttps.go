package httpsuse

import (
	"net/url"
	"net/http"
	"crypto/tls"
	"net/http/cookiejar"
	"fmt"
	"io/ioutil"
	"strings"
)

/**
  封装https 的get  post 方法
 */

func Get(requestUrl string) string {

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
	response, err := c.Get(requestUrl)
	if err != nil {
		fmt.Println("has error", err)
	}
	defer response.Body.Close()
	body, err2 := ioutil.ReadAll(response.Body)

	if err2 != nil {
		fmt.Println("has error", err2)
	}
	return string(body)
}
func Post(requestUrl string, params url.Values) string {
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
	rebody := "application/x-www-form-urlencoded"
	encode := params.Encode()
	response2, err3 := c.Post(requestUrl, rebody, strings.NewReader(encode))
	if err3 != nil {
		fmt.Println("has error", err3)
	}
	defer response2.Body.Close()
	body, err4 := ioutil.ReadAll(response2.Body)

	if err3 != nil {
		fmt.Println("has error", err4)
	}
	return string(body)

}
