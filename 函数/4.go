package main

import "fmt"

func funcone(a, b int) int {
	return a + b
}

func main() {
	repeat := funcone(123, 2)
	fmt.Println(repeat)
}
