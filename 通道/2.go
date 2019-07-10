package main

import "fmt"

func main() {
	borad := make(chan int)
	go func() { borad <- 707070 }()

	gosg := <-borad
	fmt.Println(gosg)

}
