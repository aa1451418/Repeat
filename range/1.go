package main

import "fmt"

func main() {
	a := [12]int{1, 3, 4, 5, 6}

	for b := range a {
		fmt.Println(a[b])
	}

}
