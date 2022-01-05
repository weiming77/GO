package global

import "sync"

// Package level variables:
//They are accessible throughout the package unit
const ConferenceName = "Go Conference"
const ConferenceTickets = 50

var BookedTickets uint = 0
var UserTickets uint = 0

var Destinations [11]string = [11]string{"Unknown", "Kuala Lumpur", "Bangkok", "Singapore", "Tokyo", "ShangHai", "London", "Berlin", "Paris", "Reykjavík", "Amsterdam"}

//var Bookings = []string{} // slice aka dynamic array
//var Bookings = make([]map[string]string, 0) // empty slice of maps
type TBookData struct {
	FirstName   string
	LastName    string
	UserTickets uint
	FromCity    int
	ToCity      int
}

var Bookings = make([]TBookData, 0)

// Waitgroup: waits for the launched goroutine (separated thread from the main thread) to finish
// Package "sync" provides basic synchoronization functionality
var WG = sync.WaitGroup{}
