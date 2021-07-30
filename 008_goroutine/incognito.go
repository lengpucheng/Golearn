package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 匿名go 函数 无参数匿名函数
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B defer")
			// 在函数中退出 goroutine
			runtime.Goexit()
			fmt.Println("B")
		}()
		fmt.Println("A")
	}() // ()为执行该函数

	// 带参数匿名函数
	go func(a int, b int) bool {
		fmt.Println("a is=", a, " b is =", b)
		return true
	}(1, 20)

	// 死循环避免main退出
	for {
		time.Sleep(1 * time.Second)
	}
}
