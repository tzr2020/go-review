/*
	延迟（defer）
	defer用于延迟执行函数
	延迟语句为顺序执行，延迟函数为逆序执行
*/

package main

import "fmt"

var a, b, c int = 1, 2, 3

func main() {
	defer fmt.Println(a)
	a = 0
	defer fmt.Println(b)
	defer fmt.Println(c)
}
