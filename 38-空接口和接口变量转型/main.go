/*
	空接口和接口变量转型
	空接口变量可以存储任意类型的值
*/

package main

import "fmt"

// 空接口，就是没有方法签名的接口
type empty interface{}

func main() {
	i := 1

	var e empty
	e = i
	fmt.Printf("e, value: %v type: %T\n", e, e)

	// 空接口类型变量不支持算术运算操作，需要转型
	// e = e + 2

	// 空接口变量转型方式一
	if instance, ok := e.(int); ok {
		fmt.Println("空接口变量的值是int类型")
		instance = instance + 2
		fmt.Printf("instance, value: %v type: %T\n", instance, instance)
	} else if instance, ok := e.(string); ok {
		fmt.Println("空接口变量的值是string类型")
		fmt.Printf("instance, value: %v type: %T\n", instance, instance)
	}

	// 空接口变量转型方式二
	switch instance := e.(type) {
	case int:
		fmt.Println("空接口变量的值是int类型")
		instance = instance + 2
		fmt.Printf("instance, value: %v type: %T\n", instance, instance)
	case string:
		fmt.Println("空接口变量的值是string类型")
		fmt.Printf("instance, value: %v type: %T\n", instance, instance)
	}
}
