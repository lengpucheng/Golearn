package main

import "fmt"

// 单个返回
// 函数名 首字母 大写表示外包可以用  小写只能在包内调用
func fun(str string, num int) int {
	fmt.Println("str= ", str)
	fmt.Println("num= ", num)
	ans := 100 + num
	return ans
}

// 多个匿名返回
func funcs(str string, num int) (int, string) {
	fmt.Println("str= ", str)
	fmt.Println("num= ", num)
	ans := num * 10
	return ans, str
}

// 多个返回 给对应返回变量赋值即可
// 多个返回值相同 可以使用   （va,v2 int)
// 两个返回的值 实际上是形参 直接使用或返回也可以有默认值（0）
func funcsn(str string, num int) (ans1 int, ans2 string) {
	fmt.Println("str= ", str)
	fmt.Println("num= ", num)

	// 直接返回两个也可以
	//return 100,"ans"

	ans1 = num * 10
	ans2 = str + "a"
	return
}

func main() {
	fmt.Println("\nfun-------")
	c := fun("this is fun", 100)
	fmt.Println("c= ", c)

	fmt.Println("\nfuncs-------")
	ans, str := funcs("this funcs", 10)
	fmt.Println("ans= ", ans, "\nstr= ", str)

	fmt.Println("\nfuncsn-------")
	ans1, ans2 := funcsn("this funcsn", 15)
	fmt.Println("ans= ", ans1, "\nstr= ", ans2)
}
