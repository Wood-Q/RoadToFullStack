package main

import (
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}

	var n = 10

	wg.Add(n)

	for i := 1; i <= n; i++ {
		go func(i int) {
			defer wg.Done()
			println("goroutine", i, "done")
		}(i)
	}
	wg.Wait()
}
