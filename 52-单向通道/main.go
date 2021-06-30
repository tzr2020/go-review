/*
	单向通道
*/

package main

import "fmt"

func main() {
	c := make(chan int)

	go sendData(c)
	go receData(c)

	fmt.Scanln()
}

func sendData(c chan<- int) { // 只写通道
	defer close(c)
	for i := 1; i <= 3; i++ {
		c <- i
	}
}

func receData(c <-chan int) { // 只读通道
	for {
		data, ok := <-c
		if !ok {
			break
		}
		fmt.Println(data)
	}
}
