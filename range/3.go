package main

import "fmt"

func main() {
	a := [12]int{1, 2, 4, 5, 6, 7, 8, 9, 10}

	for b := range a {
		fmt.Println(a[b])
	}
}
