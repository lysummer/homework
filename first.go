package main

import (
	"fmt"
)

func main() {
	//aaaa
	//aaaaaadadad
	//dadad
	mymap := make(map[string]int)
	mymap["李"] = 5
	mymap["张三"] = 1
	mymap["李四"] = 2
	mymap["王五"] = 3

	for k, v := range mymap {
		fmt.Println(k, v)
	}
}
