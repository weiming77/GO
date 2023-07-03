package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)     // channel with buff, sender can drop the value and leave without waiting for receiver receive the value in channel
	exit := make(chan struct{}) // channel without buff, system will block sender from leaving until receiver receive the value in channel

	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			fmt.Println(time.Now(), i, "sending...")
			ch <- i
			fmt.Println(time.Now(), i, "sent!")

			time.Sleep(1 * time.Second)
		}

		fmt.Println(time.Now(), "all completed, leaving goroutine")
	}()

	go func() {
		// xxx : This is overcomplicated becaue it only channel,
		// select only shines when using multiple channels
		/*
			for {
				select {
				case v, open := <-ch:
					if !open {
						close(exit)
						return
					}
					fmt.Println(time.Now(), "Value has been received!", v)
					//default:
					//	fmt.Println("Nothing is happenning")
				}
			}
		*/
		// XXX: In case where only ONE channel is used
		defer close(exit)
		for v := range ch {
			fmt.Println(time.Now(), "Value in the channel has received!", v)
		}
	}()

	fmt.Println(time.Now(), "Waiting for everything to be completed...")
	<-exit
	fmt.Println(time.Now(), "The END!")
}
