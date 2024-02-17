package main

import (
	"fmt"
	"sync"
)

/*
A function that is responsible for managing ticket allocations
and responding to user requests.
It listens for incoming requests on a ticket channel and signals
on another Done Channel when it's time to stop.
*/
func manageTicket(ticketChan <-chan int, doneChan <-chan struct{}, tickets *int) {
	for {
		select {
		case userId := <-ticketChan:
			if *tickets > 0 {
				*tickets--
				fmt.Printf("User %d purchased a ticket. Tickets remainig: %d\n", userId, *tickets)
			} else {
				fmt.Printf("User %d found no tickets.\n", userId)
			}
		case <-doneChan:
			fmt.Printf("Tickets remaining: %d\n", *tickets)
		}
	}
}

/*
A function that simulate a user trying to buy a ticket.
It sends a request to the manageTicket goroutine through ticketChan.
*/
func buyTicket(wg *sync.WaitGroup, ticketChan chan<- int, userId int) {
	defer wg.Done()
	ticketChan <- userId
}

func main() {
	var wg sync.WaitGroup           // WaitGroup to wait for all goroutines to finish.
	tickets := 500                  // Total number of tickets available
	ticketChan := make(chan int)    // Channel for sending ticket purchase requests
	doneChan := make(chan struct{}) // Channel for signaling the stop

	go manageTicket(ticketChan, doneChan, &tickets)

	for userId := 0; userId < 2000; userId++ {
		wg.Add(1)
		go buyTicket(&wg, ticketChan, userId)
	}

	wg.Wait()
	doneChan <- struct{}{}
}
