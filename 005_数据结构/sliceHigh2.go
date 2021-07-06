package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	// 切片 左闭右开  第0到第一
	num1 := nums[0:2]
	fmt.Printf("the num1 is %v\n", num1)

	// 取左边开始都后边全部
	num2 := nums[0:]
	fmt.Printf("the [0:] is %v\n", num2)

	// 表示取右边开始到 4-1 为止
	num3 := nums[:4]
	fmt.Printf("the [:4] is %v\n", num3)

	// 表示取全部
	num4 := nums[:]
	fmt.Printf("the [:] is %v\n", num4)

	// 引用 指的是指针 会同时修改
	num4[0] = 0
	fmt.Printf("the num4 is %v\n", num4)
	fmt.Printf("the nums is %v\n", nums)

	// copy 可以直接拷贝 是值的拷贝
	num5 := make([]int, 5)
	copy(num5, nums)
	num5[1] = 0
	fmt.Printf("the num5 is %v\n", num5)
	fmt.Printf("the nums is %v\n", nums)
}
