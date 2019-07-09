package main

import "fmt"

func main() {

	a := make([]string, 12)
	a[1] = "醛固酮"
	a[2] = "醛固酮"
	a[3] = "醛固酮"

	for b := range a {
		fmt.Println(a[b])
	}
}
