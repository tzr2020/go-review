// 映射的声明与初始化

package main

import (
	"fmt"
)

func main() {
	// 映射的声明与初始化
	// var m map[string]int // 未初始化，值为nil
	// fmt.Printf("m, value: %v, addr: %p\n", m, m)

	// m = make(map[string]int) // 通过make分配内存空间，值不为nil
	// fmt.Printf("m, value: %v, addr: %p\n", m, m)

	// var m map[string]int = map[string]int{
	// 	"zhangsan": 89,
	// 	"lisi":     56,
	// 	"wangwu":   66,
	// }

	// var m = map[string]int{
	// 	"zhangsan": 89,
	// 	"lisi":     56,
	// 	"wangwu":   66,
	// }

	m := make(map[string]int)
	m = map[string]int{
		"zhangsan": 89,
		"lisi":     56,
		"wangwu":   66,
	}

	fmt.Printf("m, value: %v, addr: %p\n", m, m)

	// 映射是引用类型
	m2 := m
	fmt.Printf("m2, value: %v, addr: %p\n", m2, m2)

	// 通过键访问值
	fmt.Println(m["zhangsan"])

	// 遍历映射
	for k, v := range m {
		fmt.Println(k, v)
	}

	// 查看键值对在映射是否存在
	value, ok := m["wangwu"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("not found")
	}

	value, ok = m["zhaoliu"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("not found")
	}

	// 删除键值对
	if _, ok := m["wangwu"]; ok {
		delete(m, "wangwu")
	} else {
		fmt.Println("not found")
	}

	fmt.Printf("m, value: %v, addr: %p\n", m, m)

	// 清空映射
	m = map[string]int{}
	fmt.Printf("m, value: %v, addr: %p\n", m, m)
}
