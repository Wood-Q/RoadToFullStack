package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	go func() {
		for i := 1; i <= 100; i += 2 {
			<-ch1
			fmt.Println("A", i)
			ch2 <- i
		}
	}()

	go func() {
		for i := 2; i <= 100; i += 2 {
			<-ch2
			fmt.Println("B", i)
			ch1 <- i
		}
	}()

	ch1 <- 1
	time.Sleep(time.Second * 1)
}
