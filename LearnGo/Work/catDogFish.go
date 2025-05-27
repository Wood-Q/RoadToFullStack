package main

import (
	"fmt"
	"time"
)

func main() {
	// var wg sync.WaitGroup

	// wg.Add(3)

	// chcat := make(chan struct{})
	// chdog := make(chan struct{})
	// chfish := make(chan struct{})

	// go func() {
	// 	defer wg.Done()
	// 	for {
	// 		<-chcat
	// 		fmt.Println("cat")
	// 		time.Sleep(time.Millisecond * 500)
	// 		chdog <- struct{}{}
	// 	}
	// }()

	// go func() {
	// 	defer wg.Done()
	// 	for {
	// 		<-chdog
	// 		fmt.Println("dog")
	// 		time.Sleep(time.Millisecond * 500)
	// 		chfish <- struct{}{}
	// 	}
	// }()

	// go func() {
	// 	defer wg.Done()
	// 	for {
	// 		<-chfish
	// 		fmt.Println("fish")
	// 		time.Sleep(time.Millisecond * 500)
	// 		chcat <- struct{}{}
	// 	}
	// }()

	// chcat <- struct{}{}
	// wg.Wait()

	waitChan := make(chan struct{}, 1)

	catchan := make(chan struct{}, 1)
	dogchan := make(chan struct{}, 1)
	fishchan := make(chan struct{}, 1)

	go func() {
		for {
			<-catchan
			fmt.Println("cat")
			time.Sleep(time.Millisecond * 500)
			dogchan <- struct{}{}
		}
	}()

	go func() {
		for {
			<-dogchan
			fmt.Println("dog")
			time.Sleep(time.Millisecond * 500)
			fishchan <- struct{}{}
		}
	}()

	go func() {
		for {
			<-fishchan
			fmt.Println("fish")
			time.Sleep(time.Millisecond * 500)
			catchan <- struct{}{}
		}
	}()

	catchan <- struct{}{}
	<-waitChan
	<- time.Tick(time.Second * 5)
}
