// 类型别名和定义类型

package main

import "fmt"

func main() {
	// 类型别名
	type stringAlias = string
	var sa stringAlias = "Hello, world!"
	fmt.Printf("sa, value: %v, type: %T\n", sa, sa)

	// 定义类型
	type newInt int
	var ni newInt = 1
	fmt.Printf("ni, value: %v, type: %T\n", ni, ni)
}
