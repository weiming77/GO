package global

import "sync"

// Package level variables:
//They are accessible throughout the package unit
const ConferenceName = "Go Conference"
const ConferenceTickets = 50

var BookedTickets uint = 0
var UserTickets uint = 0

var Destinations [11]string = [11]string{"Unknown", "Kuala Lumpur", "Bangkok", "Singapore", "Tokyo", "ShangHai", "London", "Berlin", "Paris", "Reykjavík", "Amsterdam"}

//var Bookings = []string{} // slice aka dynamic array or you can
//var Bookings = make([]string, 0)
//var Bookings = make([]map[string]string, 0) // empty slice of maps
// PS: MAKE function take parameters: <dataType>, <length>, <capacity>
// you can find out by len(Bookings) and cap(Bookings) respectively
type TBookData struct {
	FirstName   string
	LastName    string
	UserTickets uint
	FromCity    int
	ToCity      int
}

var Bookings = make([]TBookData, 0)

// many more shortcuts
var city = struct {
	name       string
	country    string
	population uint
}{
	name:       "Kuala Lumpur",
	country:    "Malaysia",
	population: 250000,
}

// Waitgroup: waits for the launched goroutine (separated thread from the main thread) to finish
// Package "sync" provides basic synchoronization functionality
var WG = sync.WaitGroup{}
