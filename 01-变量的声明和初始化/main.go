// 变量的声明和初始化

package main

import "fmt"

func main() {
	// var i int     // 声明并使用零值隐式初始化变量
	// i = 0		 // 变量赋值
	// var i int = 1 // 声明并使用字面值常量显式初始化变量
	// var i = 1     // 省略类型，会继续类型推断
	i := 1 // 简短声明语法，仅限函数体内使用

	fmt.Println(i)
}
