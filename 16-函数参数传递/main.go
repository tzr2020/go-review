// 函数传参

package main

import "fmt"

var i int = 1

// 形式参数为值类型
func incr(i int) {
	i = i + 1
}

// 形式参数为引用类型
func incr2(i *int) {
	*i = *i + 1
}

func main() {
	// 值传递
	incr(i)
	// 函数外部的实际参数i没有发生改变
	fmt.Printf("i, value: %v, type: %T\n", i, i)

	// 引用传递
	incr2(&i)
	// 函数外部的实际参数i发生改变
	fmt.Printf("i, value: %v, type: %T\n", i, i)
}
