// 格式化打印

package main

import "fmt"

func main() {
	s := "Hello, world!"
	fmt.Printf("s, value: %v\n", s)
	fmt.Printf("s, value: %#v\n", s)
	fmt.Printf("s, type: %T\n", s)

	t := true
	fmt.Printf("t, value: %t\n", t)

	i := 7
	fmt.Printf("i, value: %b\n", i)
	fmt.Printf("i, value: %d\n", i)

	f := 1.23456789
	fmt.Printf("f, value: %f\n", f)
	fmt.Printf("f, value: %.3f\n", f)
}
