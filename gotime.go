package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"bytes"
)

func main() {
	//fmt.Println(myjson())
	TrimTest()
	//EncodeTest()
}
func myjson() string {
	json_str := "{\"device\": \"1\",\"data\": [{\"humidity\": \"27\",\"time\": \"2017-07-03 15:23:12\"},{\"humidity\": \"2\",\"time\": \"2018-07-03 15:23:12\"}]}"

	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(json_str), &m)
	if err != nil {
		fmt.Println(err)
	} else {

		data := m["data"]
		if v, ok := data.([]interface{})[0].(map[string]interface{}); ok {
			//fmt.Println(ok, v["humidity"], v["time"].(string))
			return v["time"].(string)
		}
	}
	return ""
}

type Values map[string][]string

func TrimTest() {
	s := "{18829290974}"
	s2 := "1333333}"
	fmt.Println(strings.Trim(s, "{}"))
	fmt.Println(strings.Split(s2, "}")[0]) 
	fmt.Println(s[1:len(s)-1])



}

func EncodeTest() string {
	var buffer bytes.Buffer
	re := make(Values)
	re.Add("uer_name", "123456")
	re.Add("uer_phone", "18829290974")
	re.Add("uer_pass", "zxft")
	for ke, va := range re {
		buffer.WriteString(ke)
		buffer.WriteString("="+va[0]+"&")
	}
	fmt.Println(buffer.String()[0:len(buffer.String())-1])

	return buffer.String()

}

func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}
