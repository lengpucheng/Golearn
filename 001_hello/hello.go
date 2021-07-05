package main // 包名

// 多个包可以用括号
import (
	"fmt"
	"time"
)

// main 函数
func main() { // 括号必须和函数名同行
	// 加不加分号都可以
	fmt.Println("hello Golang")
	time.Sleep(1 * time.Second)
	fmt.Println("wait 1 second")

}
