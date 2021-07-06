package main

import "fmt"

func main() {
	// 1. 声明map的key类型和value类型
	var kv map[string]string
	if kv == nil {
		fmt.Println("the map is null")
		kv = make(map[string]string, 10)
	}

	kv["name"] = "yhlu"
	kv["age"] = "22"
	kv["gender"] = "boy"
	// 无序的 哈希排序
	fmt.Println(kv)

	// 2. make 推导
	kv2 := make(map[string]string, 5)
	fmt.Println(kv2)

	// 3. 直接赋值
	kv3 := map[int]string{
		1: "one",
		2: "two",
		3: "there",
	}
	// 哈希排序 数字的哈希为本身
	fmt.Println(kv3)

}
