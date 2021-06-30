/*
	循环接收通道数据
	3种方式
*/

package main

import "fmt"

func main() {
	c := make(chan int)
	go sendData(c)

	// for {
	// 	data := <-c
	// 	if data == 0 { // 通道关闭时，接收的数据是零值
	// 		break
	// 	}
	// 	fmt.Println(data)
	// }

	// for {
	// 	data, ok := <-c
	// 	if !ok { // 通过ok判断通道是否已经关闭
	// 		break
	// 	}
	// 	fmt.Println(data)
	// }

	for data := range c { // 自动判断通道是否已经关闭
		fmt.Println(data)
	}
}

func sendData(c chan int) {
	defer close(c)
	for i := 1; i <= 3; i++ {
		c <- i
	}
}
