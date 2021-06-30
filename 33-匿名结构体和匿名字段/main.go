// 匿名结构体和匿名字段

package main

import "fmt"

// 已命名结构体类型
type Person struct {
	Name  string
	Age   int8
	uint8 // 匿名字段，使用类型名字段名
}

func main() {
	// 匿名结构体，未命名结构体类型
	p := struct {
		Name string
		Age  int8
		uint8
	}{
		Name:  "zhangsan",
		Age:   11,
		uint8: 0,
	}
	fmt.Printf("p, value: %+v, type: %T\n", p, p)

	p2 := Person{
		Name:  "lisi",
		Age:   12,
		uint8: 1,
	}
	fmt.Printf("p2, value: %+v, type: %T\n", p2, p2)
}
