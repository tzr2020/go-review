/*
	函数值和函数类型
	函数是一种类型，可以作为值
*/
package main

import "fmt"

// 函数类型
type opt func(x, y int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func cala(x, y int, o opt) int {
	return o(x, y)
}

// func cala(x, y int, o func(int, int) int) int {
// 	return o(x, y)
// }

func main() {
	// 函数值
	f := add
	fmt.Println(f(1, 2))

	f = sub
	fmt.Println(f(1, 2))
	fmt.Printf("f, type: %T\n", f) // f, type: func(int, int) int

	// 函数值赋值给函数类型的变量
	var o opt
	fmt.Printf("o, value: %v, type: %T\n", o, o) // o, value: <nil>, type: main.opt

	o = add
	fmt.Println(o(1, 2))
	fmt.Printf("o, value: %v, type: %T\n", o, o) // o, value: 0x229460, type: main.opt

	// 函数值作为实际参数传入函数
	fmt.Println(cala(1, 2, add))
}
