package main

import "fmt"

func main() {
	as := make(chan string)
	go func() { as <- "channel" }()

	gos := <-as
	fmt.Println(gos)
}
