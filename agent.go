package main

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/cpu"
)

func main() {
	v, _ := mem.VirtualMemory()
	c, _ := cpu.Percent(0, false)
	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total/1024/1024/1024, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(c)
}
