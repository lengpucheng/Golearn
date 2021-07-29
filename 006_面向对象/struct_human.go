package main

import "fmt"

type Human struct {
	name   string
	gender string
}

func (this *Human) Eat() {
	fmt.Println("this is Human eating....")
}

func (this *Human) Walk() {
	fmt.Println("this is Human walking....")
}

type Chinese struct {
	Human // 继承了Human 全部
	age   int
}

// 重写父类方法
func (this *Chinese) Eat() {
	fmt.Println("this is Chinese eat")
}

func (this *Chinese) Talk() {
	fmt.Println("Chinese speak Chinese")
}

func main() {
	h := Human{"lpc", "man"}
	h.Eat()
	h.Walk()

	// 定义子类
	chinese := Chinese{Human{"luyh", "man"}, 10}
	chinese.Eat()
	chinese.Walk()
	chinese.Talk()
}
