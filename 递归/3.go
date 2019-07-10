package main

import "fmt"

func clodman(a int) int {
	for a < 100 {
		return 1
	}
	return a * clodman(a*12)
}

func main() {
	var a = 12
	fmt.Println(a, clodman(a))
}
