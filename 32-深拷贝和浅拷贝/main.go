// 深拷贝和浅拷贝

package main

import "fmt"

type Person struct {
	Name string
	Age  int8
}

func main() {
	p := Person{
		Name: "zhangsan",
		Age:  11,
	}
	fmt.Printf("p, value: %+v, type: %T\n", p, p)

	// 深拷贝，另外分配内存
	p2 := p
	fmt.Printf("p2, value: %+v, type: %T\n", p2, p2)

	// 结构体是值类型
	p2.Name = "lisi"
	p2.Age = 12
	fmt.Printf("p2, value: %+v, type: %T\n", p2, p2)
	fmt.Printf("p, value: %+v, type: %T\n", p, p)

	// 浅拷贝，共享内存
	p3 := &p
	fmt.Printf("p3, value: %+v, type: %T\n", p3, p3)

	// 结构体指针是引用类型
	p3.Name = "wangwu"
	p3.Age = 13
	fmt.Printf("p3, value: %+v, type: %T\n", p3, p3)
	fmt.Printf("p, value: %+v, type: %T\n", p, p)
}
