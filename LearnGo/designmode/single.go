package main

import (
	"fmt"
	"sync"
)

type Single struct {
	Name string
}

var (
	single *Single
	once   sync.Once
)

func GetInstance(name string) *Single {
	once.Do(func() {
		single = &Single{
			Name: name,
		}
	})
	return single
}

func main() {
	s1 := GetInstance("main")
	fmt.Println(s1.Name)
}
