package main

import (
	"fmt"
	"time"
)

func main() {
	//select
	//作用就是监听多个channel，类似于poll，epoll，可以监听多个channel的读写情况

	done1 := make(chan struct{})
	done2 := make(chan struct{})

	go func() {
		time.Sleep(time.Second * 1)
		done1 <- struct{}{}
	}()

	go func() {
		time.Sleep(time.Second * 2)
		done2 <- struct{}{}
	}()
	//随机执行，防止饥饿，如果done1和done2都阻塞，则随机选择一个执行
	//应用场景
	timer:=time.NewTimer(5*time.Second)
	for{
		select {
		case <-done1:
			fmt.Println("g1 done")
		case <-done2:
			fmt.Println("g2 done")
		case<-timer.C:
			fmt.Println("default")
			return
		}
	}
}
