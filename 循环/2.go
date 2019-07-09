package main

import "fmt"

func main() {
	for a := 0; a < 15; a++ {
		fmt.Println(a)
		a = a + 2
	}
}
