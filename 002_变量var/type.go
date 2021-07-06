package main

import "fmt"

// type 用于声明类型来表示之后的某一个类型
// 声明num类型 表示 int
type num int

func main() {
	// 和普通类型一样使用
	var nums num = 10
	fmt.Println(nums)
}
