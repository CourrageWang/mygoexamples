package main

import (
	"errors"
	"net/rpc"
	"net"
	"log"
	"net/rpc/jsonrpc"
)

//  go json  rpc  server
type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	if err != nil {
		log.Fatal("ResolveIPAddr  has an error ", err)
		return
	}

	listener, err2 := net.ListenTCP("tcp", tcpAddr)
	if err2 != nil {
		log.Fatal("ListenTCP has an error ", err2)
	}
	for {
		conn, err3 := listener.Accept()
		if err3 != nil {
			continue
		}

		jsonrpc.ServeConn(conn)

	}
}
