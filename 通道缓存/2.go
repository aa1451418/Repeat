package main

import "fmt"

func main() {
	pass1 := make(chan int, 5)
	pass1 <- 1
	pass1 <- 2
	pass1 <- 3
	pass1 <- 4
	pass1 <- 5
	fmt.Println(pass1)
}
