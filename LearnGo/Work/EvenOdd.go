package main

import "fmt"

func evenSum(arr []int, evenChan chan<- int) {
	sum := 0
	for _, num := range arr {
		if num%2 == 0 {
			sum += num
		}
	}
	evenChan <- sum
}

func oddSum(arr []int, oddChan chan<- int) {
	sum := 0
	for _, num := range arr {
		if num%2 != 0 {
			sum += num
		}
	}
	oddChan <- sum
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenChan := make(chan int)
	oddChan := make(chan int)

	go evenSum(arr, evenChan)
	go oddSum(arr, oddChan)

	evenSum := <-evenChan
	oddSum := <-oddChan

	fmt.Println("evenSum:", evenSum)
	fmt.Println("oddSum:", oddSum)

}
