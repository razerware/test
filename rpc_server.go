package main

import (
	"log"
	"runtime"
	"net/rpc"
	"test/service"
	"net"
	"fmt"
	"time"
)

//////////////////////////////////////////////////
//Server
func main() {
	runtime.GOMAXPROCS(4)

	arith := new(service.Arith)
	server := rpc.NewServer()
	log.Printf("Register service:%v\n", arith)
	server.Register(arith)

	log.Printf("Listen tcp on port %d\n", 4200)
	l, e := net.Listen("tcp", ":4200")

	if e != nil {
		log.Fatal("Listen error:", e)
	}
	log.Println("Ready to accept connection...")
	conCount := 0
	go func(){
		for {
			select{
			case c:=<-service.Ch:
				fmt.Printf("insert %d",c)
				time.Sleep(10*time.Second)
			}
		}

	}()
	for {
		conn, err := l.Accept()
		fmt.Println("444")
		if err != nil {
			log.Fatal("Accept Error:,", err)
			continue
		}
		conCount++
		log.Printf("Receive Client Connection %d\n", conCount)
		go server.ServeConn(conn)
	}
	time.Sleep(100*time.Second)
}
