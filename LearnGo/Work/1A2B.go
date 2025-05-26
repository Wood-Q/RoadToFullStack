package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	go func() {
		for i := 0; i < 10; i++ {
			<-ch1
			fmt.Println(i)
			ch2 <- 1
		}
		wg.Done()
	}()
	go func() {
		for i := 65; i < 75; i++ {
			<-ch2
			//转为大写字母
			fmt.Println(string(rune(i)))
			ch1 <- 1
		}
		wg.Done()
	}()
	ch1 <- 1
	wg.Wait()
}
