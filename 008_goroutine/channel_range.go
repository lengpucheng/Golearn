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

	// 使用range 将c中的数据全部读出来   当c被关闭且无数据时候 data返回 否则会一直阻塞
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("main end...")

}
