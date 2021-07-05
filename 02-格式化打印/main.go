// 格式化打印

package main

import "fmt"

func main() {
	// 所有类型通用
	s := "Hello, world!"
	fmt.Printf("s, value: %v\n", s)
	fmt.Printf("s, value: %#v\n", s)
	fmt.Printf("s, type: %T\n", s)

	// 布尔型
	t := true
	fmt.Printf("t, value: %t\n", t)

	// 整型
	i := 7
	fmt.Printf("i, value: %b\n", i) // 十进制值
	fmt.Printf("i, value: %d\n", i) // 二进制值

	// 浮点型
	f := 1.23456789
	fmt.Printf("f, value: %f\n", f)
	fmt.Printf("f, value: %.3f\n", f) // 限制3位小数
}
