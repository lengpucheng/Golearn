package libb

import "fmt"

func init() {
	fmt.Println("this is liba Two")
}

// 首字母大小表示对外可调用
// 首字母小写表示只能在包内调用
func LibaInc() {
	println("libA Inc A")
}
