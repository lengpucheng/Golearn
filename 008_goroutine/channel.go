package main

import "fmt"

func main() {
	// 定义channel 无缓冲
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine end..")
		fmt.Println("goroutine ing ")
		// 向channel 写入数据
		c <- 666
	}()

	// channel 会将管道两边的时间进行同步  在异步执行中若出现了不同步状态  读写先到的一边会被阻塞  直到另一边到达读取/写入点

	// 从c中读取数据
	num := <-c
	fmt.Println("num= ", num)
	fmt.Println("end")

}
