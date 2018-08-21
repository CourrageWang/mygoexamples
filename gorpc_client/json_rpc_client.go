package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

// go rpc json client
type Args struct {
	A, B int
}
type Quotient struct {
	Quo, Rem int
}

func main() {

	//if len(os.Args) != 2 {
	//	fmt.Println("Usage:", os.Args[0], "server:port")
	//	log.Fatal(1)
	//}
	//service := os.Args[1]
	service := "192.168.0.104:1234" //需要调用服务的地址。
	client, err1 := jsonrpc.Dial("tcp", service)
	if err1 != nil {
		log.Fatal("dialing:", err1)
	}
	args := Args{17, 8}
	var reply int
	err2 := client.Call("Arith.Multiply", args, &reply)
	if err2 != nil {
		log.Fatal("arith error :", err2)
	}
	fmt.Printf("Arith: %d *%d  = %d \n", args.A, args.B, reply)

	var quot Quotient
	err3 := client.Call("Arith.Divide", args, &quot)
	if err3 != nil {
		log.Fatal("arith error ", err3)
	}
	fmt.Printf("Arith : %d/%d = %d remainder %d \n", args.A, args.B, quot.Quo, quot.Rem)

}
