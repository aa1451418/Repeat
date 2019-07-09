package main

import "fmt"

func main() {
	var a map[string]string
	a = make(map[string]string)
	a["我是映射"] = "我是映射"
	fmt.Println(a)

}
