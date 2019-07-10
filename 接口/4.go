package main

import "fmt"

type Twwiter interface {
	googole()
}

type single struct{}

func (single single) googole() {
	for q := 0; q < 4; q++ {
		if q != 0 {
			fmt.Println("q<0")
		}
	}
}

type riginel struct{}

func (riginel riginel) googole() {
	for q := 0; q < 4; q++ {
		if q > 0 {
			fmt.Println(q)
		}
	}
}

func main() {
	var Twwiter Twwiter
	Twwiter = new(single)
	Twwiter.googole()

	Twwiter = new(riginel)
	Twwiter.googole()
}
