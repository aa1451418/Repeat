package main

import "fmt"

func main() {
	Holiday := map[int]int{1: 1, 2: 2, 3: 4}

	sensible := map[int]string{1: "map one", 2: "map tow"}

	cablecar := map[int]string{1: "Car"}

	fmt.Println(Holiday, sensible, cablecar)
}
