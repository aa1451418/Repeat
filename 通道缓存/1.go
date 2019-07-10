package main

import "fmt"

func main() {
	adh := make(chan string, 4)
	adh <- "通道缓存"
	adh <- "Tow"
	adh <- "Tree"
	adh <- "for"
	fmt.Println(adh)
}
