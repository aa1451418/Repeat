package main

import "fmt"

func main() {
	a := 10

	switch a {
	case 1:
		fmt.Println("a不等于10")
	case 10:
		fmt.Println("a=10")
	}

}
