package main

import "fmt"

func main() {
	// make 第三个参数为容量
	number := make([]int, 4, 5)
	fmt.Printf("the number type is %T , len is %d, cap is %d \n", number, len(number), cap(number))

	// 在其中追加元素 长度会变
	number = append(number, 1)
	fmt.Printf("this number is %v ,len is %d , cap is %d\n", number, len(number), cap(number))

	// 继续追加 容量翻倍
	number = append(number, 1)
	fmt.Printf("this number is %v ,len is %d , cap is %d\n", number, len(number), cap(number))

	// 不声明容量时候 容量和长度相同
	num := make([]int, 5)
	fmt.Printf("this number is %v ,len is %d , cap is %d\n", num, len(num), cap(num))
	num = append(num, 1)
	fmt.Printf("this number is %v ,len is %d , cap is %d\n", num, len(num), cap(num))
}
