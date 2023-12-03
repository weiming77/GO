package main

import (
	"fmt"
	"time"
)

func main() {
	evilNinjas := []string{"Tommy", "Jonny", "Bobby", "Andy"}
	smokeSignal := make(chan string)
	completed := make(chan bool)
	now := time.Now()
	var iCount uint = 0

	go func() {
		for _, ninja := range evilNinjas {
			go attack(ninja, smokeSignal)
		}

		for {
			select {
			case deceased, ok := <-smokeSignal:
				if !ok {
					close(completed)
					return
				}
				fmt.Println("evil ninja", deceased, "is down!")
				iCount++
			default:
				if iCount == uint(len(evilNinjas)) {
					close(smokeSignal)
				}
			}
		}
	}()

	<-completed
	fmt.Printf("Mission accomplished in %v", time.Since(now))

}

func attack(target string, deceased chan<- string) {
	fmt.Println("Throwing ninja stars at", target)
	time.Sleep(time.Second * 2)
	deceased <- target
}
