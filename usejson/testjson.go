package main

import (
	"encoding/json"
	"fmt"
)
/**  golang 解析json时，如果按照interface{}方式接收数据的话，会按照如下方式进行

        bool, for JSON booleans

        float64, for JSON numbers

        string, for JSON strings

       []interface{}, for JSON arrays

       map[string]interface{}, for JSON objects

       nil for JSON null
 */
/**
    解析json数据的时候，结构体中定义的数据类型需要和json数据的类型一致否则会引起转换错误 如下：
    cannot unmarshal number into Go struct field Result.code of type string
 */
type Result struct {
	//Code    string `json:"code"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Version string `json:"version"`
}

func main() {
	str := `{"code":0,"message":"success!","version":"v 1.3.2"}`
	result := &Result{}
	err := json.Unmarshal([]byte(str), result)
	if err != nil {
		fmt.Println("there has error", err)
	}
	if 0 == result.Code {
		fmt.Println("数据返回正确")
	}
	fmt.Println(result.Code)
	// method 2
	parseJson := make(map[string]interface{})
	err2 := json.Unmarshal([]byte(str), &parseJson)
	if err2 != nil {
		fmt.Println(err2)
	}
	code :=int(parseJson["code"].(float64))
	if 0==code {
		fmt.Println("数据返回成功")
	}

	// 以上两种解析方式，无论使用哪种去解析，关注json原始数据的类型，至关重要，如果知道类型的话，使用第一种会更优


}
