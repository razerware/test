package main

import (
	"net/rpc"
	"log"
	"fmt"
	"net"
)

type args struct {
	A, B int
}

type quotient struct {
	Quo, Rem int
}

func main() {
	address, err := net.ResolveTCPAddr("tcp", "127.0.0.1:4200")
	if err != nil {
		panic(err)
	}
	conn, _ := net.DialTCP("tcp", nil, address)
	defer conn.Close()

	client := rpc.NewClient(conn)
	defer client.Close()

	args := &args{7, 8}
	// Synchronous call
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	// Asynchronous call
	quotient := new(quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done // will be equal to divCall
	if replyCall.Error != nil {
		log.Fatal("arith error:", replyCall.Error)
	}
	fmt.Printf("Arith: %d/%d=%d...%d", args.A, args.B, quotient.Quo, quotient.Rem)
	// check errors, print, etc.
}
