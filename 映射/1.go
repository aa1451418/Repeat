package main

import "fmt"

func main() {
	var a map[string]string
	a = make(map[string]string)
	a["法国"] = "法国"

	for b := range a {
		fmt.Println(a[b])
	}

}
