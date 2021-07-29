package main

import "fmt"

// interface{} 表示万能类型  可以接受全部类型
func all(arg interface{}) {
	fmt.Println("this is all")
	fmt.Println(arg)

	// 断言是否为 某个 类型
	value, ok := arg.(string)
	if !ok {
		fmt.Println("this is age not is string")
	} else {
		fmt.Println("arg is string Type, value =", value)
		fmt.Printf("value type is %T \n", value)
	}

	fmt.Println("---------------------")
}

type Book struct {
	name string
}

func main() {
	all(Book{"my dream"})
	all(10)
	all("string")

}
