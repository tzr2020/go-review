/*
	恢复（recover）
	recover拦截恐慌，如果程序进入恐慌，返回恐慌信息并使程序恢复运行，如果程序没有进入恐慌则返回nil
	recover必须在延迟语句的函数中使用
*/

package main

import "fmt"

func f1() {
	fmt.Println("this is f1")
}

func f2() {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg)
			fmt.Println("程序恢复运行")
		}
	}()
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
