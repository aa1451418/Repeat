package main

import "fmt"

func main() {

	var a map[string]string
	a = make(map[string]string)
	a["yuang"] = "this yuang"

	fmt.Println(a)
}
