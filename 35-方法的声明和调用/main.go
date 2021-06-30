/*
	方法的声明和调用
	方法是特殊的函数，是某个类型的行为成员，只有该类型的实例变量才能调用
*/

package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// 函数的声明
func Area(r Rectangle) float64 {
	return r.width * r.height
}

// 方法的声明，包含接受者
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	r := Rectangle{
		width:  5,
		height: 10,
	}
	fmt.Println(Area(r))
	fmt.Println(r.Area())

	c := Circle{
		radius: 10,
	}
	fmt.Println(c.Area())

	/*
		方法与函数的不同
		1.含义：方法是某个类型的行为成员，而函数是一组语句序列
		2.接受者：方法有接受者，只有接受者的类型的实例变量才能调用方法，而函数没有接受者
		3.名字：方法可以重名（接受者不同），而函数不可以重名
	*/
}
