/*
	通道（channel）
	通道用于协程之间通信，具有阻塞机制
*/

package main

import "fmt"

func main() {
	var c chan string
	fmt.Printf("c, value: %v, type: %T\n", c, c) // c, value: <nil>, type: chan string

	c = make(chan string) // 阻塞式发送数据，直到通道中的数据被接收
	// c = make(chan string, 1) // 具有缓冲的通道，发送数据可以暂时存储数据而不阻塞
	fmt.Printf("c, value: %v, type: %T\n", c, c) // c, value: 0xc00001a0c0, type: chan string

	defer close(c)

	go func() {
		c <- "hello"
	}()

	data := <-c // // 阻塞式接收数据，直到通道中的存在数据
	fmt.Println(data)

	fmt.Scanln()
}
