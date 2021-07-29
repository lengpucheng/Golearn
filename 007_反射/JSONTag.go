package main

import (
	"encoding/json"
	"fmt"
)

// tag 用于  json 或 orm
type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"ATM", 1999, 10, []string{"DG", "LINA", "YG"}}

	// 编码  结构体 --》 json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal fail")
		return
	}

	fmt.Printf("json= %s\n", jsonStr)

	// json --》 结构体
	myMove := Movie{}
	err = json.Unmarshal(jsonStr, &myMove)
	if err != nil {
		fmt.Println("encode to struct have err")
		return
	}
	fmt.Println(myMove)

}
