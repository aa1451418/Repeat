package main

import "fmt"

func dk(a int) int {
	for a < 10 {
		return 2
	}
	return a * dk(a*5)
}

func main() {
	a := 3
	fmt.Println(a, dk(a))
}
