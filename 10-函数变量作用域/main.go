// 函数变量作用域

package main

import "fmt"

// 全局变量
var a int = 1
var b int = 2

/*
	x，y是形式参数，是局部变量
	res是函数内声明的变量，也是局部变量
	均不能被函数外部访问
*/
func incr(x, y int) int {
	res := x + y
	return res
}

func main() {
	// 函数可以访问全局变量
	fmt.Println(incr(a, b))
}
