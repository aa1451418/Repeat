package main

import "fmt"

func main() {

	todao := make(chan string)

	go func() {
		todao <- "todao"
	}()

	daoto := <-todao
	fmt.Println(daoto)

}
