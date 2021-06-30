/*
	互斥锁
*/

package main

import (
	"fmt"
	"sync"
)

var x int // 共享资源

var lock sync.Mutex // 互斥锁

func add() {
	for i := 1; i <= 5000; i++ {
		// lock.Lock()   // 上锁
		x = x + 1 // 临界区，非原子操作
		// lock.Unlock() // 解锁
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// 多个goroutine同时操作一个共享资源，发生竞态问题
	go func() {
		add()

		wg.Done()
	}()

	go func() {
		add()

		wg.Done()
	}()

	wg.Wait()

	// 理想情况下x=10000，但发生竞态问题，极有可能是x<10000
	// 可以通过互斥锁的上锁操作和解锁操作解决
	// 临界区的非原子操作之前上锁，之后解锁
	fmt.Println(x)
}
