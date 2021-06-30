/*
	select分支选择语句
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "every 2s"
			time.Sleep(time.Second * 2)
		}
	}()

	// for {
	// 	fmt.Println(<-c1)
	// 	fmt.Println(<-c2)
	// }

	for {
		select {
		case data1 := <-c1:
			fmt.Println(data1)
		case data2 := <-c2:
			fmt.Println(data2)
		}
	}
}
