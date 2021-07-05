package main

import "fmt"

// 形式参数存储 实际参数的 地址
func changValue(p *int) {
	*p = 10
}

func main() {
	a := 1
	fmt.Println(a)
	// & 表示获取地址
	fmt.Println(&a)
	// * 表示地址对应的变量
	fmt.Println(*&a)

	// 实际参数为地址
	changValue(&a)
	fmt.Println(a)

	point1 := &a
	// 可以获取二级指针
	point2 := &point1

	point1var := *point1
	point2var := **point2

	// *二级指针可以得出一级指针地址
	println(point1var, point2var, *point2)
}
