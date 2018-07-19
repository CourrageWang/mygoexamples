package main

import "fmt"

type UserInfo struct {
	Name    string
	Pass    string
	Age     int
	Job     string
	Address string
}
type Point struct { //可以将 公共元素提取出来使用
	X, Y int
}
type Circle struct {
	Point // 匿名成员
	Radius int
}
type Wheel struct {
	Circle
	Spoke int
}

func InitWheel() {
	var w Wheel
	w.X = 10
	w.Y = 20
	w.Radius = 14
	w.Spoke = 2
}
func (u *UserInfo) SetNamePAss(name, pass string) *UserInfo {
	u.Name = name
	u.Pass = pass
	return u
}
func getMoreInfo() {
	m := make(map[string]UserInfo)
	u1 := UserInfo{"Test1", "123", 12, "worker", "bj"}
	u2 := UserInfo{"Test2", "456", 13, "worker", "bj1"}
	u3 := UserInfo{"Test3", "789", 14, "worker", "bj2"}
	m[u1.Name] = u1
	m[u2.Name] = u2
	m[u3.Name] = u3
	for k, v := range m {
		fmt.Println("k = ", k, ",v=", v)
	}
}

func main() {
	/*var uinfo UserInfo
	uinfo.SetNamePAss("test", "953284")
	fmt.Println(uinfo.Name, uinfo.Pass)*/
	getMoreInfo()
}
