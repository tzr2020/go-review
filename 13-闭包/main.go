/*
	闭包
	是附有数据的行为
	函数 + 引用环境 = 闭包
*/

package main

import "fmt"

func add(i int) int {
	sum := 0 // sum的生命周期与add相同，add返回后，sum的内存空间被释放
	sum += i
	return sum
}

// 闭包实现计数器
func adder() func(int) int {
	sum := 0 // sum的生命周期与adder返回的匿名函数相同，adder返回后，sum的内存空间不会被释放
	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(add(i))
	}

	fmt.Println("----------")

	f := adder()
	for i := 1; i <= 5; i++ {
		fmt.Println(f(i))
	}
}
