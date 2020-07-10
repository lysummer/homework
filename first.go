package main

import (
	"fmt"
)

func main() {
	mymap := make(map[string]int)
	mymap["李"] = 5
	mymap["张三"] = 1
	mymap["李四"] = 2

	for k, v := range mymap {
		fmt.Println(k, v)
	}
}
