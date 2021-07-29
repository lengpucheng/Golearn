package main

import (
	"fmt"
	"reflect"
)

type Info struct {
	// 使用 `` 来定义标签Tag
	Name string `info:"name" doc:"名称"`
	Msg  string `info:"sex"`
}

// 解析结构体标签
func getTag(arg interface{}) {
	t := reflect.TypeOf(arg)

	for i := 0; i < t.NumField(); i++ {
		// 获取第i个字段的标签中  info 字段
		info := t.Field(i).Tag.Get("info")
		doc := t.Field(i).Tag.Get("doc")
		fmt.Println("info =", info)
		fmt.Println("doc =", doc)
		fmt.Println("---------------")
	}
}

func main() {
	info := Info{"info", "this is no msg"}
	getTag(info)
}
