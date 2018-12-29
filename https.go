package main

import (
	"fmt"
	"net/url"
	"net/http"
	"crypto/tls"
	"net/http/cookiejar"
	"io/ioutil"
	"strings"
	"srun4kAuthIntf/src/floger"
	"encoding/json"
)

type Accesstoken struct {

	Message string `json:"message"`
	Version string `json:"version"`
	Code int `json:"code"`
	Data map[string]string `json:"data"`
}


func main() {
	str := Get("https://192.168.0.195:8001/api/v1/auth/get-access-token") // BcLf5QzbIlg9nNRyDlVwVyPExHM05IMQ
	fmt.Println(str)
	acc := &Accesstoken{}
	err := json.Unmarshal([]byte(str), acc)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println("---->",acc.Data["access_token"])

	/*requestParams := make(url.Values)
	requestParams.Add("access_token", "g8JpvEqJGfEbuxYU46dmdOd9pk2idf9E")
	requestParams.Add("user_name", "我是中文")
	requestParams.Add("user_real_name", "陕西人")
	requestParams.Add("user_password", "9766908")
	requestParams.Add("group_id", "test")
	requestParams.Add("products_id", "2")
	st := Post("https://192.168.0.195:8001/api/v1/users", requestParams)
	floger.Info(st)*/

}
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
	floger.Info("创建用户")
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
