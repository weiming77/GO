package main

import (
	"fmt"
	"time"
)

func main() {
	evilNinjas := []string{"Tommy", "Jonny", "Bobby", "Andy"}
	smokeSignal := make(chan bool)
	now := time.Now()

	go func() {
		defer close(smokeSignal)
		for _, ninja := range evilNinjas {
			go attack(ninja)
		}
	}()

	for _, ninja := range evilNinjas {
		<-smokeSignal
		fmt.Println(ninja, "is down!")
	}

	fmt.Printf("Mission accomplished in %v", time.Since(now))

}

func attack(target string) {
	fmt.Println("Throwing ninja stars at ", target)
	time.Sleep(time.Second * 2)

}
