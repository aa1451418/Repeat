package main

import "fmt"

func main() {

	ddd := make(chan string, 5)
	ddd <- "chan one"
	ddd <- "chan tow"
	ddd <- "chan three"
	ddd <- "chan foru"
	ddd <- "chan five"
	fmt.Println(ddd)
}
