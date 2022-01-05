package main

import (
	"booking-app/global"
	"booking-app/utils"
	"fmt"
)

func main() {
	var firstName, lastName string = "", ""
	// var idx int
	var iFromCity, iToCity int
	var sInput string

	fmt.Printf("Welcome to %v booking system!\n", global.ConferenceName)
	fmt.Println("Get your ticket(s) here to attend")

	for {
		fmt.Printf("\nTotal tickets: %v\tAvailable: %v\n", global.ConferenceTickets, (global.ConferenceTickets - global.BookedTickets))
		utils.Prompt4Names(&firstName, &lastName)
		iFromCity, iToCity = utils.Prompt4Destinations()

		// we can loop array/slice using for index, value:= range bookings {}
		// "_" is an blank identifier; it's to ignore an variable you don't want to use
		global.UserTickets = utils.Prompt4NoofTickets(iFromCity, iToCity)

		fmt.Printf("Customer: %v, %v\tBooked tickets: %v From: %v to: %v", firstName, lastName, global.UserTickets, global.Destinations[iFromCity], global.Destinations[iToCity])
		fmt.Print("\nConfirm (Y/N): ")
		fmt.Scan(&sInput)

		if (sInput == "Y") || (sInput == "y") {
			global.BookedTickets += global.UserTickets
			/* using slice/dynamic array of string
			for i := 0; i < int(global.UserTickets); i++ {
				global.Bookings = append(global.Bookings, firstName+" "+lastName+" from "+global.Destinations[iFromCity]+" to "+global.Destinations[iToCity])
				idx += 1
			}
			*/
			/* All keys have the same data type. ie string
			// map[key]value
			var bookingData = make(map[string]string)
			bookingData["firstName"] = firstName
			bookingData["lastName"] = lastName
			bookingData["fromCity"] = strconv.FormatInt(int64(iFromCity), 10)
			bookingData["toCity"] = strconv.FormatInt(int64(iToCity), 10)
			bookingData["userTickets"] = strconv.FormatInt(int64(global.UserTickets), 10)
			global.Bookings = append(global.Bookings, bookingData)
			*/
			var bookingData = global.TBookData{
				FirstName:   firstName,
				LastName:    lastName,
				FromCity:    iFromCity,
				ToCity:      iToCity,
				UserTickets: global.UserTickets}

			global.Bookings = append(global.Bookings, bookingData)
			// Before the goroutine, need to call sync.add: Sets the number of
			// goroutine to wait for (increase the counter by the provided number )
			global.WG.Add(1)
			go utils.PrintBookingReceipt(bookingData) // separated thread
		}

		if global.BookedTickets >= global.ConferenceTickets {
			fmt.Println("Tickets sold out!")
			utils.PrintBookingList(global.Bookings)
			// Exit end of program
			return
		}
	}
	// sync.Wait(): Block until the WaitGroup counter is 0
	global.WG.Wait()
}
