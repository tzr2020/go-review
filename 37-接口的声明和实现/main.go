/*
	接口的声明和实现
	接口是一种类型，包含一组方法签名，用于定义某类对象的行为，等待其它类型去实现行为的细节
	类型都是隐式实现接口的，类型只要声明了某个接口的所有方法，那么该类型就实现了该接口
*/

package main

import "fmt"

type animal interface {
	call()
}

type dog struct{}

type cat struct{}

func (d dog) call() {
	fmt.Println("wangwang")
}

func (c cat) call() {
	fmt.Println("miaomiao")
}

func main() {
	d := dog{}
	c := cat{}

	var a animal
	fmt.Printf("a, type: %T\n", a)

	// 实现了某个接口的类型的实例变量，可以赋值给该接口变量，接口变量也就拥有了多种实现细节，这一特性体现了面向对象的多态特性
	a = d
	fmt.Printf("a, type: %T\n", a)
	a.call()

	a = c
	fmt.Printf("a, type: %T\n", a)
	a.call()
}
