/*
Go program <- Channel Communication -> Go Runtime <- System Calls -> OS Kernel
Go program <- Goroutine Management -> Go Runtime <- OS Threads -> OS Kernel
Concurrency vs parallelism
a) Parallelism
- Parallel events or tasks execute simultaneously and independently.
- True parallel events require multiple CPUs.
- Multiple tasks can start at the same time but with espensive costs of CPU hardware

b) Concurrency
- Concurrent tasks or events are interleaving and can happen in any given order.
- It is a non-deterministic way of achieving multiple tasks.
- ie
func (wg *WaitGroup) Add(delta int)
func (wg *WaitGroup) Done()
func (wg *WaitGroup) Wait()

Important Note:
* Never use sleep in Production, as it does not provide any concurrency gurantees
* All of the types in the sync package should be passed by pointer for functions.
*/
package main

import (
	"fmt"
	"sync"
)

func hello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello, World")
}

func goodbye() {
	fmt.Println("Good Bye!")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go hello(&wg)
	wg.Wait()
	goodbye()
}
