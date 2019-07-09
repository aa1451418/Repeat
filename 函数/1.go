package main

import "fmt"

/* 定义hls函数 */
func hls(a, b int) int {
	return a * b
}

func main() {
	res := hls(123, 2435)
	fmt.Println(res)
}
