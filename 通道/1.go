package main

import "fmt"

func main() {

	asgs := make(chan string)

	go func() { asgs <- "通道" }()

	asgb := <-asgs
	fmt.Println(asgb)

}
