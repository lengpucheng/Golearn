package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Println("new Goroutine i=", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// 创建一个  goroutine
	go newTask()
	i := 0
	// main 退出 所有的go 全部goroutine都会退出
	for {
		i++
		fmt.Println("main go i=", i)
		time.Sleep(1 * time.Second)
	}
}
