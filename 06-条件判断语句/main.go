// 条件判断语句

package main

import "fmt"

func main() {
	score := 80

	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 60 {
		fmt.Printf("及格")
	} else {
		fmt.Printf("不及格")
	}
}
