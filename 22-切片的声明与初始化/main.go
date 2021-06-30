package main

import "fmt"

func main() {
	// 切片的声明与初始化
	// var s []int
	// s := make([]int, 3)
	// s = []int{1, 2, 3}
	// var s []int = []int{1, 2, 3}
	s := []int{1, 2, 3}

	fmt.Printf("s, value: %v, type: %T\n", s, s)

	// 遍历切片
	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("index: %v, value: %v\n", i, s[i])
	// }

	for k, v := range s {
		fmt.Printf("index: %v, value: %v\n", k, v)
	}

	// 切片是引用类型
	change(s)
	fmt.Printf("s, value: %v, type: %T\n", s, s)

	// 切片截取操作
	s2 := s[0:2]
	fmt.Printf("s2, value: %v, type: %T\n", s2, s2)

	a := [3]int{1, 2, 3}
	s3 := a[0:2]
	fmt.Printf("s3, value: %v, type: %T\n", s3, s3)
}

func change(s []int) {
	for i := 0; i < len(s); i++ {
		s[i] += 10
	}
}
