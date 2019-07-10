package main

import "fmt"

func main() {
	a := make([]int, 5)

	b := map[int]string{1: "map one", 2: "map tow"}

	c := map[int]int{1: 1, 2: 2, 3: 4, 4: 5}
	fmt.Println(a, b, c)

}
