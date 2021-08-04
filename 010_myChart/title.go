package main

import (
	"fmt"
	"time"
)

func TitleInfo() {
	info := "______________myChart_______________"
	info += "\nstart  go " + time.Now().String()
	info += "\n______________myChart_______________"
	fmt.Println(info)
}
