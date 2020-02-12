package main

/***
  测试s48串口demo
*/

const (
	netWork = "tcp"
	add     = "192.168.1.19:5000"
	command ="d201015a015b"
)

func main() {
	//conn, err := net.DialTimeout(netWork, add, 5*time.Second)
	//
	//if err!=nil {
	//	println(err)
	//}
	//conn.Write([]byte(command))
  println(111)
}
