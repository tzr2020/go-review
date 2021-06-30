// 匿名函数

package main

import "fmt"

type opt func(x, y int) int

func cala(x, y int, o opt) int {
	return o(x, y)
}

func main() {
	// 定义匿名函数并直接调用
	fmt.Println(func(x, y int) int {
		return x + y
	}(1, 2))

	func() {
		fmt.Println("Hello, world!")
	}()

	// 匿名函数赋值给变量
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(cala(1, 2, add))

	f := func() {
		fmt.Println("Hello, world!")
	}
	f()

	// 匿名函数作为实际参数
	fmt.Println(cala(1, 2, func(x, y int) int {
		return x + y
	}))
}
