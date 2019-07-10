package main

import "fmt"

func main() {

	mooth := [12]int{1, 23, 4}

	for gate := range mooth {
		fmt.Println(mooth[gate])
	}
}
