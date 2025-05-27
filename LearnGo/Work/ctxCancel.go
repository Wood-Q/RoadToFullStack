package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// // 创建一个带有超时功能的上下文
	// // context.WithTimeout 返回一个新的Context和一个CancelFunc。
	// // 在5秒后，或者调用cancel()函数时，新Context的Done通道会被关闭。
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	// defer cancel() // 确保在main函数退出时调用cancel，释放资源

	// catchan := make(chan struct{}, 1)
	// dogchan := make(chan struct{}, 1)
	// fishchan := make(chan struct{}, 1)

	// go func(ctx context.Context) {
	// 	defer fmt.Println("cat goroutine exited")
	// 	for {
	// 		select {
	// 		case <-catchan:
	// 			fmt.Println("cat")
	// 			time.Sleep(time.Millisecond * 500)
	// 			dogchan <- struct{}{}
	// 		case <-ctx.Done(): // 监听上下文的Done通道
	// 			return // 收到取消信号，退出循环
	// 		}
	// 	}
	// }(ctx) // 将上下文传递给goroutine

	// go func(ctx context.Context) {
	// 	defer fmt.Println("dog goroutine exited")
	// 	for {
	// 		select {
	// 		case <-dogchan:
	// 			fmt.Println("dog")
	// 			time.Sleep(time.Millisecond * 500)
	// 			fishchan <- struct{}{}
	// 		case <-ctx.Done():
	// 			return
	// 		}
	// 	}
	// }(ctx)

	// go func(ctx context.Context) {
	// 	defer fmt.Println("fish goroutine exited")
	// 	for {
	// 		select {
	// 		case <-fishchan:
	// 			fmt.Println("fish")
	// 			time.Sleep(time.Millisecond * 500)
	// 			catchan <- struct{}{}
	// 		case <-ctx.Done():
	// 			return
	// 		}
	// 	}
	// }(ctx)

	// catchan <- struct{}{} // 启动打印序列

	// // main goroutine会阻塞在这里，直到ctx.Done()通道关闭 (5秒后)
	// <-ctx.Done()
	// fmt.Println("\nMain goroutine exiting. Context cancelled.")

	// // 可以选择性地添加一个短暂的等待，确保所有goroutine有时间退出
	// time.Sleep(time.Millisecond * 100)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	catChan := make(chan struct{}, 1)
	dogChan := make(chan struct{}, 1)
	fishChan := make(chan struct{}, 1)

	go func(ctx context.Context) {
		defer fmt.Println("cat goroutine exited")
		for {
			select {
			case <-catChan:
				println("cat")
				dogChan <- struct{}{}
			case <-ctx.Done():
				return
			}
		}
	}(ctx)
	go func(ctx context.Context) {
		defer fmt.Println("dog goroutine exited")
		for {
			select {
			case <-dogChan:
				println("dog")
				fishChan <- struct{}{}
			case <-ctx.Done():
				return
			}
		}
	}(ctx)
	go func(ctx context.Context) {
		defer fmt.Println("fish goroutine exited")
		for {
			select {
			case <-fishChan:
				println("fish")
				catChan <- struct{}{}
			case <-ctx.Done():
				return
			}
		}
	}(ctx)

	catChan <- struct{}{}

	<-ctx.Done()

	fmt.Println("Main goroutine exiting. Context cancelled.")

	time.Sleep(time.Millisecond * 100)
}
