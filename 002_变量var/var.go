package main

import "fmt"

// 声明全局变量
var ga int = 100

func main() {
	// 不赋值默均为0  （内存0填充）
	var a int
	fmt.Printf("type of a =%T\n", a)
	fmt.Println("a= ", a)

	// 赋值
	var b int = 100
	fmt.Printf("type of b =%T\n", b)
	fmt.Println("b= ", b)

	// 省去数据类型 通过值自动赋值类型
	var c = 100
	fmt.Printf("type of c =%T\n", c)
	fmt.Println("c= ", c)

	// 省去 var 关键字
	// 只能在函数体内使用  不能声明全局变量
	d := 100
	fmt.Printf("e =%d , type of e =%T\n", d, d)
	fmt.Printf("ga =%d , type of ga =%T\n", ga, ga)

	// 多变量声明
	var x1, x2 int = 100, 200
	println(x1, "  ---  ", x2)

	// 不同类型
	var y1, y2 = 100, "abcd"
	println(y1, " --- ", y2)

	// 多行
	var (
		z1 = 100
		z2 = true
		z3 = 3.14159
		z4 = "pai"
	)
	println("-z1- ", z1, " -z2- ", z2, "-z3- ", z3, "-z4- ", z4)
	var (
		f1 int
		f2 string
		f3 float32
		f4 bool
	)
	println(f1, f2, f3, f4)
}
