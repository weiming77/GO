package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan bool)
	now := time.Now()

	go func(done chan bool) {
		fmt.Println("I am working on it")
		time.Sleep(3 * time.Second)
		//done <- true OR
		close(done)
	}(ch)

	<-ch
	fmt.Printf("My job is done in %v", time.Since(now))
}
