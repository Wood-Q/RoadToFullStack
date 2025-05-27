package main

import "fmt"

//type Animal interface {
//	speak() string
//}
//
//type Dog struct {
//}
//
//func (d Dog) speak() string {
//	return "Wolf"
//}
//
//type AnimalFactory interface {
//	CreateAnimal() Animal
//}
//type DogFactory struct {
//}
//
//func (d DogFactory) CreateAnimal() Animal {
//	return Dog{}
//}
//
//func main() {
//	var factory AnimalFactory
//	factory = DogFactory{}
//	dog := factory.CreateAnimal()
//	fmt.Println(dog.speak())
//}

type Animal interface {
	move() string
}

type Dog struct{}

func (d Dog) move() string {
	return "狗爬"
}

type AbstractFactory interface {
	Create() Animal
}

type DogFactory struct{}

func (df DogFactory) Create() Animal {
	return Dog{}
}

func main() {
	var factory AbstractFactory
	factory = DogFactory{}
	dog1 := factory.Create()
	fmt.Println(dog1.move())

}
