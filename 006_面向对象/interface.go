package main

import "fmt"

// 定义接口 接口实际上是一个指针 可以指向任何一个实现类
type Animal interface {
	Eat()
	GetColor() string
	GetType() string
}

// 重写了接口中的全部方法 即默认为实现了接口
type Cat struct {
	Color string
}

func (this *Cat) Eat() {
	fmt.Println("cat eating....")
}

func (this *Cat) GetColor() string {
	return this.Color
}
func (this *Cat) GetType() string {
	return "cat"
}

type Dog struct {
	Color string
}

func (this *Dog) Eat() {
	fmt.Println("dog eating....")
}

func (this *Dog) GetColor() string {
	return this.Color
}
func (this *Dog) GetType() string {
	return "dog"
}

func showAnimal(animal Animal) {
	fmt.Println("the Type is ", animal.GetType(), " the Color is ", animal.GetColor())
	animal.Eat()
}

func main() {
	animal := &Cat{"black"}
	animal.GetType()

	animal2 := &Dog{"yellow"}
	animal2.GetType()

	showAnimal(animal)
	showAnimal(animal2)
}
