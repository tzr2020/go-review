// 类型转换

package main

import "fmt"

func main() {
	i := 97

	f := float32(i)
	s := string(i)

	fmt.Printf("f, value: %v, type: %T\n", f, f)
	fmt.Printf("s, value: %v, type: %T\n", s, s)
}
