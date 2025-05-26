package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func doSomething(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("finish")
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("do nothing")
		}
	}
}

func main() {
	wg.Add(1)

	ctx, cancel := context.WithCancel(context.Background())
	go doSomething(ctx)
	time.Sleep(10 * time.Second)
	cancel()
	wg.Wait()
}
