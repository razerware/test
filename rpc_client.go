package main

import (
	"net/rpc"
	"log"
	"fmt"
	"net"
)

type MonitorStats struct {
	CpuPercent string
	MemPercent string
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

	args := &MonitorStats{"7%", "8%"}

	res:=false
	for{
		lzycall:=client.Go("MonitorStats.Collect",args,&res,nil)
		relzycall:=<-lzycall.Done
		if relzycall.Error != nil {
			log.Fatal("arith error:", relzycall.Error)
		}
		fmt.Printf("lzy %v",res)
	}

}
