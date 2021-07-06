package main

import (
	"fmt"
)

// 引用传递
func add(maps map[string]string) {
	for contry, copet := range maps {
		maps[contry] = "copet " + copet
	}
}

// 遍历显示 此处传递的是引用指针
func show(maps map[string]string) {
	for key, value := range maps {
		fmt.Println("the key is ", key, "value is ", value)
	}
}

func main() {
	city := make(map[string]string, 5)
	city["CN"] = "BJ"
	city["US"] = "DC"
	city["UK"] = "LD"

	// 显示
	show(city)
	// 引用传递
	add(city)
	show(city)

	// 1. 增加
	fmt.Println("add-----")
	city["JP"] = "TK"
	show(city)

	// 2. 修改
	fmt.Println("put-----")
	city["CN"] = "BeiJin"
	show(city)

	// 3. 删除
	fmt.Println("del-----")
	delete(city, "UK")
	show(city)
	delete(city, "HK")

	// 4. 查找
	fmt.Println("look-----")
	fmt.Println("the cn cop is", city["CN"])
	fmt.Printf("the TW is %v,%T", city["TW"], city["TW"])

}
