package main

import "fmt"

func main() {
	// 1. 推到声明,赋默认值,且固定其长度
	slice1 := []int{1, 2, 3}
	fmt.Printf("the slice1 is %v \n", slice1)

	// 2. 声明但不分配内存空间
	var slice2 []int
	fmt.Printf("the slice2 is %v \n", slice2)
	// 判空
	fmt.Println(slice2 == nil)
	// make 可以用来开辟内存空间（默认为0,二进制全0)
	slice2 = make([]int, 3)
	fmt.Printf("the slice2 is %v \n", slice2)

	// 3. 声明并开辟内存
	var slice21 []int = make([]int, 5)
	fmt.Printf("the slice2_1 is %v \n", slice21)

	// 4. 推导并开辟内存
	slice22 := make([]int, 5)
	fmt.Printf("the slice2_2 is %v \n", slice22)
}
