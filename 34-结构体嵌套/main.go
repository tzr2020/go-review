// 结构体嵌套

package main

import "fmt"

type Person struct {
	Name    string
	Age     int8
	Address Address // 聚合关系
}

type Address struct {
	City string
}

type Student struct {
	ID     int
	Person // 继承关系，通过结构体嵌套和匿名字段，来模拟面向对象的继承特性，子类可以使用父类的成员
}

func main() {
	p := Person{
		Name: "zhangsan",
		Age:  11,
		Address: Address{
			City: "guangzhou",
		},
	}
	fmt.Printf("p.Address.City, value: %v, type: %T\n", p.Address.City, p.Address.City) // 访问p.Address.City通过p.Address.City

	s := Student{
		ID: 1001,
		Person: Person{
			Name: "lisi",
			Age:  12,
			Address: Address{
				City: "beijing",
			},
		},
	}
	fmt.Printf("s.Person.Name, value: %v, type: %T\n", s.Name, s.Name) // 访问s.Person.Name通过s.Name，可以省略了Person
}
