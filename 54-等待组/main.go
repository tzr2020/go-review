/*
	等待组（WaitGroup）
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// stop := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		time.Sleep(time.Microsecond * 500)
		fmt.Println("job1 done")

		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("job2 done")

		wg.Done()
	}()

	// go func() {
	// 	time.Sleep(time.Second * 3)
	// 	stop <- true
	// }()

	// fmt.Scanln()
	// <-stop
	// time.Sleep(time.Second * 3)
	wg.Wait()
}
