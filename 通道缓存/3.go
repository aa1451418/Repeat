package main

import "fmt"

func main() {
	asg := make(chan int, 4)
	asg <- 1
	asg <- 2
	asg <- 3
	asg <- 4
	fmt.Println(asg)
}
