package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	//要监控多少个goroutine执行
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}

	wg.Wait()
}
