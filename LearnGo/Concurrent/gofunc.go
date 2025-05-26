package main

import (
	"fmt"
	"time"
)

func asyncPrint() {
	fmt.Println("hello")
}

func main() {
	go asyncPrint() //go关键字启动一个goroutine
	time.Sleep(1 * time.Second) //等待1秒，确保goroutine有足够的时间执行
	fmt.Println("world")
}
