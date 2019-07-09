package main

import "fmt"

func main() {

	/* 整数型数组 */
	a := [2]int{1, 2}

	/* 字符串类型数组 */
	b := []string{1: "字符串类型数组"}

	/* make函数数组 */
	c := make([]int, 10)    //make整数型
	d := make([]string, 10) //make字符串型
	d[1] = "make func arry"
	fmt.Println(a, b, c, d)

}
