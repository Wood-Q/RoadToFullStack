package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "012"
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("转换失败")
	}
	fmt.Println(num)
	fmt.Println("转换成功")
}
