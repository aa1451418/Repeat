package main

import "fmt"

func main() {
	for b := 0; b < 13; b++ {
		b = b + 3
		fmt.Println(b)
	}
}
