package main

import "fmt"

// 使用const定义枚举
const (
	BJ = 0
	SH = 1
	WH = 2
)

//使用iota可以自动赋值累加 iota从0开始 每一行 +1
// 使用过iota的行后面的计算方法公式和iota行应用
const (
	CN = iota + 5
	USA
	JP
	FH
)

func main() {
	// 定义常量 只读
	const length = 10
	fmt.Println("length= ", length)
	fmt.Println("BJ= ", BJ)
	fmt.Println("WH= ", WH)
	fmt.Println("SH= ", SH)
	fmt.Println("CN= ", CN)
	fmt.Println("USA= ", USA)
	fmt.Println("JP= ", JP)
	fmt.Println("FH= ", FH)

}
