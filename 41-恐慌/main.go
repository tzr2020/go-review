/*
	恐慌（panic）
	panic使程序进入恐慌，中断程序的运行，输出恐慌信息
	panic必须在函数中使用
*/

package main

import "fmt"

func f1() {
	fmt.Println("this is f1")
}

func f2() {
	fmt.Println("this is f2")
	panic("程序进入恐慌")
}

func f3() {
	fmt.Println("this is f3")
}

func main() {
	f1()
	f2()
	f3()
}
