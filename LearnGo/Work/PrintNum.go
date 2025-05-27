package main

import (
	"fmt"
	"time"
)

var n = 10

func main() {
	arr := make([]chan struct{}, n)
	for i := 0; i < n; i++ {
		arr[i] = make(chan struct{}, 1)
	}

	for i := 0; i < n; i++ {
		next := (i + 1) % n
		go printNum(i, arr[i], arr[next])
	}
	arr[0] <- struct{}{}
	time.Sleep(time.Second * 1)
}

func printNum(i int, ch, next chan struct{}) {
	for {
		<-ch
		fmt.Println(i)
		next <- struct{}{}
		return
	}
}
