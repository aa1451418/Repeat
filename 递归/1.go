package main

import "fmt"

func techer(a int) int {
	for a < 10 {
		return 1
	}
	return a * techer(a*10)
}

func main() {
	var i = 12
	fmt.Println(i, techer(i))
}
