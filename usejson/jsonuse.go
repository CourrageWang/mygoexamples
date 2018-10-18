package main

import (
	"encoding/json"
	"fmt"
)

// 因为有些结构体成员的名字并不相同，因此需要使用tag标签
type Movie struct {
	Title  string                        // `json:"-"` 解析时忽略该字段
	Year   int  `json:"released"`        // tag为标签 给结构体字段打上标签 冒号前是类型后面是标签名，解析时使用released而非Year
	Color  bool `json:"color,omitempty"` // tag里面加上omitempy，可以在序列化的时候忽略0值或者空值
	Actors []string
}

func TestJson() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
	}
	data, err := json.Marshal(movies)
	if err != nil {
		fmt.Println("JSON marshaling failed")
	}
	fmt.Printf("%s", data)

}

// 解析感兴趣的内容
func UnMarshalYouWant() {
	var Titles [] struct{ Title string }
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
	}
	data, err := json.Marshal(movies)
	if err != nil {
		fmt.Println("JSON marshaling failed")
	}
	if errs := json.Unmarshal(data, &Titles); errs != nil {
		fmt.Println("JSON unmarshaling failed")
	}
	fmt.Println(Titles)

}

// 使用map存储json
func UnMarshalToMap() {
	var dataMap []map[string]interface{} // 存储解析后的json对象
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman", "Will Smith", "Jim Carry"}},
	}
	data, err := json.Marshal(movies)
	if err != nil {
		fmt.Println("JSON marshaling failed")
	}
	errs := json.Unmarshal(data, &dataMap)
	if errs != nil {
		fmt.Println("JSON unmashaling  failed...", errs)
	}
	fmt.Println(dataMap)

}

func main() {
	//TestJson()
	//UnMarshalYouWant()
	UnMarshalToMap()

}
