package main

import "fmt"

type inface interface {
	facebook()
}

type book struct{}

func (book book) facebook() {
	for a := 0; a < 5; a++ {
		fmt.Println(a)
	}
}

type probook struct{}

func (probook probook) facebook() {
	fmt.Println("nil")
}

func main() {
	var inface inface
	inface = new(probook)
	inface.facebook()

	inface = new(book)
	inface.facebook()
}
