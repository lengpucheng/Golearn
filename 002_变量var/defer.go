package main

import "fmt"

func fun1() {
	fmt.Println("this is fun1")
}

func fun2() {
	fmt.Println("this is fun2")
}

func fun3() {
	fmt.Println("this is fun3")
}

// defer 关键字 在 {} 内其他语句执行完毕后调用结束前 生效
// 多个 defer 将 按照 先进后出执行  （先声明的后执行，压栈）
func main() {
	defer fmt.Println("this is go 1")
	fmt.Println("this is go 2")
	defer fmt.Println("this is go 3")
	fmt.Println("this is go 4")

	defer fun1()
	defer fun2()
	defer fun3()
}
