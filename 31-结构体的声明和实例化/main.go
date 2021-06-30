// 结构体的声明和实例化

package main

import "fmt"

// 声明结构体
type Person struct {
	Name string
	Age  int8
}

func main() {
	// 实例化结构体（分配内存），使用零值隐式初始化属性成员
	// var p Person
	p := Person{}
	fmt.Printf("p, value: %+v, type: %T\n", p, p)

	// 访问结构体属性成员
	p.Name = "zhangsan"
	p.Age = 11
	fmt.Printf("p, value: %+v, type: %T\n", p, p)

	// 实例化结构体，并显式初始化属性成员
	p2 := Person{
		Name: "lisi",
		Age:  12,
	}
	fmt.Printf("p2, value: %+v, type: %T\n", p2, p2)

	// 实例化结构体，使用零值隐式初始化属性成员，返回结构体指针
	p3 := new(Person)
	(*p3).Name = "wangwu"
	p3.Age = 13                                      // 语法糖
	fmt.Printf("p3, value: %+v, type: %T\n", p3, p3) // 语法糖
}
