package main

import "fmt"

type Peo struct {
	Name   string
	Gender string
}

type City struct {
	Size int
}

type Chn struct {
	Peo
	City
	Local string
}

func (t *Peo) showPeo() {
	fmt.Println("this peo name is ", t.Name)
}

func show(c City) {
	fmt.Println("the city size is ", c.Size)
}

func main() {
	chinese := Chn{
		Peo:   Peo{"china", "♀"},
		City:  City{9600000},
		Local: "亚洲",
	}

	fmt.Println(chinese)
	chinese.showPeo()

	var i interface{}
	i = chinese
	city, ok := i.(City)
	if ok {
		show(city)
	} else {
		fmt.Println("is not ok")
		fmt.Printf("the val %v ", city)
	}
}
