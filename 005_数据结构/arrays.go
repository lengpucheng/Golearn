package main

import "fmt"

// 固定数组传递参数 必须类型和长度一样  （值拷贝）
// 固定数组形参是实参副本 值传递
func array(per [4]int) {
	for value := range per {
		fmt.Println("value is ", value)
	}
}

func main() {
	// 固定长度数组
	var arrys [10]int

	for i := 0; i < 10; i++ {
		fmt.Println(arrys[i])
	}

	for i := 0; i < len(arrys); i++ {
		fmt.Println(arrys[i])
	}

	// 推导声明
	arrs2 := [4]int{2, 3, 4}
	// range 便利  (foreach)
	for index, value := range arrs2 {
		fmt.Println("index is ", index, " value=", value)
	}

	array(arrs2)
}
