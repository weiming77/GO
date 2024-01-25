package main

import (
	"fmt"
	"time"
)

func main() {
	evilNinjas := []string{"Tommy", "Jonny", "Bobby", "Andy", "Anord", "Matrix", "Pirates"}
	noofkillers := 2 // limit only 2x killers in the worker pool
	limiter := make(chan struct{}, noofkillers)
	now := time.Now()

	for _, ninja := range evilNinjas {
		limiter <- struct{}{} // Acquire a token. Waits here for token releases from the limiter.
		go attack(ninja, limiter)
	}

	// Wait for all goroutines to complete
	for i := 0; i < cap(limiter); i++ {
		limiter <- struct{}{}
	}
	fmt.Printf("Mission accomplished in %v", time.Since(now))

}

func attack(target string, deceased <-chan struct{}) {
	defer func(ninja string) {
		<-deceased // Release the token
		fmt.Printf("%s got killed\n", ninja)
	}(target)

	fmt.Println("Throwing ninja stars at", target)
	time.Sleep(time.Second * 2)
}
