package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	// 获取字段类型
	fmt.Println("type: ", reflect.TypeOf(arg))

	// 获取字段的值
	fmt.Println("value: ", reflect.ValueOf(arg))
}

type User struct {
	Id   string
	Name string
	Age  int
}

func (t *User) Call() {
	fmt.Println("t is user call....")
}

func (t User) GetName() string {
	return t.Name
}

func (t User) GetID() string {
	return t.Id
}

func main() {
	float := 999.44
	reflectNum(float)

	user := User{"001", "lpc", 21}
	user.Call()

	reflectNum(user)
	// 获取类型
	theType := reflect.TypeOf(user)
	fmt.Println(theType.Name())
	// 获取值
	theVal := reflect.ValueOf(user)
	fmt.Println(theVal)

	// 获取字段数量
	for i := 0; i < theType.NumField(); i++ {
		// 获取第i个字段
		field := theType.Field(i)
		fmt.Println(field)
		fmt.Println(field.Name) // 字段的名称
		// 获取第i个字段的值
		val := theVal.Field(i)
		fmt.Println(val)
	}

	fmt.Println("===========")
	// 获取其中的方法函数
	for i := 0; i < theType.NumMethod(); i++ {
		// 获取方法
		method := theType.Method(i)
		fmt.Printf("the %d method is %T %v \n", i, method, method)
		// 获取方法名称等
		fmt.Println(method.Name)
		fmt.Println("----------")
	}

	fmt.Println("===========")

}
