package main

import "fmt"

func main() {
	// 数组的声明与初始化
	// var a [3]int
	// a = [3]int{1, 2, 3}
	// var a [3]int = [3]int{1, 2, 3}
	// a := [3]int{1, 2, 3}
	a := [...]int{1, 2, 3}

	fmt.Printf("a, value: %v, type: %T\n", a, a)

	// 遍历数组
	// for i := 0; i < len(a); i++ {
	// 	fmt.Printf("index: %v, value: %v\n", i, a[i])
	// }

	for k, v := range a {
		fmt.Printf("index: %v, value: %v\n", k, v)
	}

	// 数组是值类型
	a2 := [...]int{1, 2, 3}
	fmt.Println(a == a2)
}
