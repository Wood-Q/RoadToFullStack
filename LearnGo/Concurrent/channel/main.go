package main

import (
	"fmt"
	"time"
)

/*
	go中channel的应用场景
	1. 消息传递，消息过滤
	2. 信号广播
	3. 事件订阅和广播
	4. 任务分发
	5. 并发控制
	6. 结果汇总
*/

func main() {
	//有缓冲channel适合消费者和生产者之间的通信
	ch := make(chan int, 1)
	ch <- 114514
	data := <-ch
	fmt.Println(data)
	//无缓冲channel适用于通知，B要第一时间知道A是否完成，用于一种信号通知事件
	ch1 := make(chan int)
	go func(ch1 chan int) {
		data := <-ch1
		fmt.Println(data)
	}(ch1)
	ch1 <- 1
	time.Sleep(time.Second)
	//for range遍历channel
	ch2 := make(chan int, 2)
	ch2 <- 2
	ch2 <- 3
	for num := range ch2 {
		fmt.Println(num)
	}
	close(ch2)

}
