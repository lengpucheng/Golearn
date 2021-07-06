package main

import "fmt"

// 动态数组 是引用传递
func actions(arr []int) {
	for i, v := range arr {
		fmt.Println("this is ", i, " the value is ", v)
	}
}

// 不声明数组大小为动态数组 slice (切片)
func main() {
	action := []int{1, 2, 3, 4, 5}

	fmt.Printf("this is action is type of %T the len is %d", action, len(action))

	actions(action)
}
