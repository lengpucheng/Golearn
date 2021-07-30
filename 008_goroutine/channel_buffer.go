package main

import (
	"fmt"
	"time"
)

func main() {
	// 带有缓冲区的channel
	c := make(chan int, 5)

	fmt.Println("len(c)=", len(c), "cap(c)=", cap(c))
	go func() {
		defer fmt.Println("go end")
		for i := 0; i < 7; i++ {
			c <- i
			fmt.Println("go ing i is ", i)
		}
	}()

	time.Sleep(2 * time.Second)

	// 当缓存区满了就会阻塞
	for i := 0; i < 7; i++ {
		num, ok := <-c
		fmt.Println("ok--", ok, "num is", num)
	}

	fmt.Println("main end")
}
