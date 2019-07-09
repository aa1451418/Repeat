package main

import "fmt"

type Jkg interface {
	telephone()
}

type Ja struct{}

func (ja Ja) telephone() {
	fmt.Println("185-2051-7543")
}

type Aj struct{}

func (aj Aj) telephone() {
	fmt.Println("1.2022501174")
}

func main() {
	var Jkg Jkg
	Jkg = new(Ja)
	Jkg.telephone()

	Jkg = new(Aj)
	Jkg.telephone()
}
