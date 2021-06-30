// 条件分支语句

package main

import (
	"fmt"
)

func main() {
	score := 55
	grade := ""

	switch {
	case score >= 90:
		grade = "A"
	case score >= 60:
		grade = "D"
	default:
		grade = "E"
	}

	switch grade {
	case "A":
		fmt.Println("优秀")
	case "D":
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}
}
