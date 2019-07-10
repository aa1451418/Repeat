package main

import "fmt"

func main() {
	a := []int{1, 2, 3}

	b := []string{1: "str"}

	c := make([]int, 3)
	d := make([]string, 3)
	fmt.Println(a, b, c, d)
}
