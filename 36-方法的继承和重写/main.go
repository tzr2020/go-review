// 方法的继承和重写

package main

import "fmt"

type Person struct {
	Name string
	Age  int8
}

type Student struct {
	ID     int
	Person // 继承关系，通过结构体嵌套和匿名字段，来模拟面向对象的继承特性，子类可以使用父类的成员
}

// 方法的继承
// 行为成员
func (p Person) Say() {
	fmt.Println("I am Person")
}

// 方法的重写，一个包含匿名字段的结构体也实现了该匿名字段实现的方法，子类也实现了与父类同名的方法
// func (s Student) Say() {
// 	fmt.Println("I am Student")
// }

func main() {
	s := Student{
		ID: 1001,
		Person: Person{
			Name: "lisi",
			Age:  12,
		},
	}
	s.Say()

	// s.Person.Say()
}
