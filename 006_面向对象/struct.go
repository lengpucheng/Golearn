package main

import "fmt"

// 声明一个结构体
type stu struct {
	name  string
	age   int
	isBoy bool
}

// 值传递 实参的副本
func addAgeV(student stu) {
	student.age = student.age + 1
	fmt.Printf("-------------addAgeV------\n the  is %T the value is %v \n", student, student)
}

// 指针传递 实参的地址
func addAgeO(student *stu) {
	student.age = (student.age) + 1
	fmt.Printf("-------------addAgeO------\n the  is %T the value is %v \n", student, student)
}

func main() {
	// 声明
	var lpc stu
	fmt.Printf("the stu is %T the value is %v \n", lpc, lpc)
	// 赋值
	lpc.name = "lengpucheng"
	lpc.age = 21
	lpc.isBoy = true
	fmt.Printf("the stu is %T the value is %v \n", lpc, lpc)

	// 推到声明
	yj := stu{"gqy", 20, false}
	fmt.Printf("the stu is %T the value is %v \n", yj, yj)

	// 值传递
	addAgeV(lpc)
	fmt.Printf("the obj is %T the value is %v \n", lpc, lpc)

	// 引用传递
	addAgeO(&yj)
	fmt.Printf("the stu is %T the value is %v \n", yj, yj)

}
