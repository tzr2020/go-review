// 常量和计数器

package main

import "fmt"

func main() {
	// const host string = "127.0.0.1"
	const host = "127.0.0.1"
	fmt.Printf("host, value: %v, type: %T\n", host, host)

	const (
		a = iota
		b
		c
	)
	fmt.Println(a, b, c)
}
