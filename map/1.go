package main

import "fmt"

func main() {

	a := map[int]string{1: "map"}
	a[0] = "This is map arry"

	b := map[int]int{1: 1, 2: 43}

	fmt.Println(a, b)
}
