package main

import "fmt"

type R interface {
	show()
}

type W interface {
	input(val string)
}

type File struct {
	value string
}

func (this *File) input(val string) {
	this.value = val
}

func (this *File) show() {
	fmt.Println("this is value is  ", this.value)
}

func main() {
	var a string
	// pair<type:string,value:the>
	a = "the"
	var obj interface{}

	obj = a

	val, _ := obj.(string)
	fmt.Println(val)

	// pair<type File,value="nil">
	file := File{"nil"}
	var r R
	// pair<type File,value="nil">
	r = &file
	r.show()

	var w W
	// pair<type File,value="nil">
	// 断言如果是true可以将对象进行转换 （返回1为val 2为断言是否成功）,此处省略二返回  同理 _
	w = r.(W)
	w.input("this is file")

	// pair<type File,value="this is file">
	r.show()

	// 接口间的上下转型不允许 pair的改变 只要pair一样 指向的内容就是一样

}
