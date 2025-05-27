package main

import (
	"fmt"
	"time"
)

func main() {
	strChan := make(chan int, 1)
	numChan := make(chan int, 1)

	strChan <- 0

	go func() {
		for i := 65; i <= 90; i++ {
			<-strChan
			fmt.Println(string(rune(i)))
			numChan <- i
		}
	}()

	go func() {
		for i := 1; i <= 26; i++ {
			<-numChan
			fmt.Println(i)
			strChan <- 0
		}
	}()

	time.Sleep(time.Second * 1)
}
