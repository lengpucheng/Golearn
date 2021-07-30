package main

import "fmt"

func fibonacie(c, quit chan int) {
	x, y := 1, 1

	for {
		// select 可以同时监控多条channel的可读状态
		select {
		// 当一个可以执行的时候就可以去监听某一个通道
		case c <- x:
			x, y = y, x+y
		case <-quit:
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacie(c, quit)
}
