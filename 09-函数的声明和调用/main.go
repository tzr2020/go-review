// 函数的声明和调用

package main

import "fmt"

// 函数声明
func add(x, y int) int {
	res := x + y
	return res
}

func main() {
	// 函数调用
	res := add(1, 2)
	fmt.Println(res)
}
