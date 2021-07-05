// 切片添加元素和复制切片

package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	fmt.Printf("s, value: %v, type: %T\n", s, s)

	// 切片添加元素
	s = append(s, 4, 5)
	fmt.Printf("s, value: %v, type: %T\n", s, s)

	s2 := []int{6, 7, 8}
	s = append(s, s2[0:2]...)
	fmt.Printf("s, value: %v, type: %T\n", s, s)

	// 复制切片
	s3 := make([]int, len(s)) // 给新切片变量分配内存，不与原切片变量共享内存
	copy(s3, s)
	fmt.Printf("s3, addr: %p, value: %v, type: %T\n", s3, s3, s3)

	change(s3)
	fmt.Printf("s, addr: %p, value: %v, type: %T\n", s, s, s)
	fmt.Printf("s3, addr: %p, value: %v, type: %T\n", s3, s3, s3)
}

func change(s []int) {
	for i := 0; i < len(s); i++ {
		s[i] += 10
	}
}
