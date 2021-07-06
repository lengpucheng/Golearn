package main

import "fmt"

func defers() {
	fmt.Println(" this is defer")
}

func returns() int {
	fmt.Println("this is return")
	return 0
}

// 若 returns 和 defer 同时存在 return 先执行
func test() int {
	defer defers()
	return returns()
}

func main() {
	test()
}
