package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		// 关闭通道  只有当没有数据传输了才有需要关闭  如果关闭后发数据会直接报错，但关闭后若有缓存是可以依然接收数据
		close(c)
	}()

	for {
		// ok 为 true 表示没有关闭
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("main end...")

}
