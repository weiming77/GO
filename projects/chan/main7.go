package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	iTimer       int = 0
	iTotal       int = 0
	iAccumulated int = 0
)

func Countdown(i int) {
	iTimer += i
}

func main() {
	now := time.Now()
	ch := make(chan int, 5) // channel with buff, sender can drop the value and leave without waiting for receiver receive the value in channel
	done := make(chan int)  // channel without buff, system will block sender from leaving until receiver receive the value in channel
	//done := make(chan struct{}) // channel without buff, system will block sender from leaving until receiver receive the value in channel

	for i := 0; i < 50; i++ {
		go func(v int) {
			// pretend as an busy go-routine
			iSec := rand.Intn(5) + 1
			time.Sleep(time.Duration(iSec) * time.Nanosecond) //
			fmt.Printf("Task %d will take %v secs for processing...\n", v, iSec)
			//done <- struct{}{}
			done <- iSec
			ch <- v
		}(i + 1)
	}

	for {
		select {
		case iTask := <-ch:
			iTotal += 1
			Countdown(-1)
			fmt.Printf("Task %d is completed.\n", iTask)
		case iSec := <-done:
			iAccumulated += iSec
			Countdown(1)
		default:
			if iTotal > 0 && iTimer == 0 {
				fmt.Printf("Total %d tasks with total processing %d secs completed in %v, leaving sender goroutine\n", iTotal, iAccumulated, time.Since(now))
				close(ch)
				return
			}
		}
	}
}
