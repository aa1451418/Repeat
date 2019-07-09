package main

import "fmt"

func main() {
	/* 整数型 */
	a := []int{1, 2, 3}

	/* 字符串类型 */
	b := []string{1: "2"}

	/* make函数型 */
	c := make([]int, 10)

	fmt.Println(a, b, c)
}
