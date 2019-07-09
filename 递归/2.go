package main

import "fmt"

func dikui(a int) int {
	for a < 10 {
		return 1
	}
	return a * dikui(a*123)
}

func main() {
	var a int = 10
	fmt.Println(a, dikui(a))
}
