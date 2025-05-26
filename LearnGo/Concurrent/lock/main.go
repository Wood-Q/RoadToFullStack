package main

import (
	"fmt"
	"sync"
)

var total int
var wg sync.WaitGroup
// var lock sync.Mutex //互斥锁
var ch chan struct{} = make(chan struct{}, 1)

func add() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		<-ch
		total+=1
		ch<-struct{}{}
	}
}

func sub() {
	defer wg.Done()

	for i := 0; i < 100000; i++ {
		<-ch
		total-=1
		ch<-struct{}{}
	}
}

func main() {
	ch <- struct{}{}
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println("total:", total)
}
