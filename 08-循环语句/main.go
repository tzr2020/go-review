// 循环语句

package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// i := 0
	// for ; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// i := 0
	// for ; ; i++ {
	// 	if i > 9 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// i := 0
	// for {
	// 	if i > 9 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// 	i++
	// }

	// s := "abcde"
	// for i, v := range s {
	// 	fmt.Printf("index: %v, value: %v, type: %T\n", i, v, v)
	// }

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v*%v=%v ", i, j, i*j)
		}
		fmt.Println()
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
