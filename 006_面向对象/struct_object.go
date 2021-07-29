package main

import "fmt"

// Hero /** 大写表示 public 对外可以访问 包括属性和方法  大写均表示 public  小写只能在本包内使用

type Hero struct {
	Name  string
	Ad    int
	Level int
}

func (this Hero) Show() {
	fmt.Println("name=", this.Name, " ad=", this.Ad, " Level=", this.Level)
}

func (this Hero) GetName() {
	fmt.Println("Name=", this.Name)
}

// SetName /** 此处是值传递  修改的副本 并不会修改原对象
func (this Hero) SetName(name string) {
	// this 是对象的拷贝
	this.Name = name
}

func (this *Hero) SetNameTrue(newName string) {
	this.Name = newName
}

func valTrans() {
	hero := Hero{"ZhangSan", 100, 1}
	hero.Show() // name is ZhangSan
	hero.SetName("LiShi")
	hero.Show() // name also is ZhangSan
	// this 是 对象的拷贝是值传递  需要使用指针

	hero.SetNameTrue("LiShi")
	hero.Show() // name also is ZhangSan

}

func main() {
	valTrans()
}
