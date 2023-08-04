package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	ch := make(chan int, 5)     // channel with buff, sender can drop the value and leave without waiting for receiver receive the value in channel
	done := make(chan struct{}) // channel without buff, system will block sender from leaving until receiver receive the value in channel

	go func() {
		for i := 0; i < 50; i++ {
			go func(v int) {
				ch <- v
				// pretend as an busy go-routine
				time.Sleep(1 * time.Second)
			}(i + 1)
			fmt.Printf("%v: %d is sent!\n", time.Now(), i+1)
		}

		fmt.Printf("%v all values completed in %v, leaving sender goroutine\n", time.Now(), time.Since(now))
		done <- struct{}{}
	}()

	go func() {
		// xxx : This is overcomplicated becaue it only channel,
		// select only shines when using multiple channels
		for {
			select {
			case v, open := <-ch:
				if !open {
					close(done)
					break
				}
				fmt.Printf("%v: Value %d has been received!\n", time.Now(), v)
			}
		}
	}()

	<-done
	fmt.Printf("%v: Everything is done in %v.", time.Now(), time.Since(now))
}
