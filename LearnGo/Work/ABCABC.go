package main

import (
	"fmt"
	"sync"
)

func print(id string, ch <-chan struct{}, nextCh chan<- struct{}, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		<-ch
		fmt.Println(id)
		nextCh <- struct{}{}
	}
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)
	ch3 := make(chan struct{}, 1)

	go print("A", ch1, ch2, wg)
	go print("B", ch2, ch3, wg)
	go print("C", ch3, ch1, wg)

	ch1 <- struct{}{}
	wg.Wait()
}
