package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	ch := make(chan int, 10) // channel with buff, sender can drop the value and leave without waiting for receiver receive the value in channel

	go func() {
		defer close(ch)
		// inputs
		fmt.Println(time.Now(), "sending...")
		for i := 0; i < 300; i++ {
			ch <- i
		}
		fmt.Printf("%v all completed in %v, leaving goroutine\n\n", time.Now(), time.Since(now))
	}()

	// done is read only channel
	for v := range ch {
		fmt.Printf("%v: %v in the channel has received!\n", time.Now(), v)
	}
	fmt.Printf("All done in %v\n", time.Since(now))
}
