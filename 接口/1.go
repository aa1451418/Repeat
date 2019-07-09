package main

import "fmt"

/* 定义接口 */
type jk interface {
	kj()
}

/* 定义结构体 */
type Jj struct{}

/* 创建方法 */
func (jj Jj) kj() {
	fmt.Println("interface one")
}

/* 定义结构体 */
type Lj struct{}

/* 创建方法 */
func (lj Lj) kj() {
	fmt.Println("interface tow")
}

/* main函数体 */
func main() {
	var jk jk
	jk = new(Lj)
	jk.kj()

	jk = new(Jj)
	jk.kj()
}
