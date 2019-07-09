package main

import "fmt"

func main() {
	for i := 0; i < 30; i++ {
		i = i + 1 - 2 + 3 - 4 + 6
		fmt.Println(i)
	}
}
