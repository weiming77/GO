package utils

import (
	"booking-app/global"
	"bufio"
	"fmt"
	"os"
	"time"
)

func Prompt4Names(firstName *string, lastName *string) {
	*firstName = ""
	*lastName = ""
	prompt := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("What is your name? ")
		prompt.Scan()
		*firstName = prompt.Text()
		if len(*firstName) > 0 {
			break
		}
	}

	for {
		fmt.Print("What is your last name? ")
		prompt.Scan()
		*lastName = prompt.Text()
		if len(*lastName) > 0 {
			break
		}
	}
}

func Prompt4Destinations() (iFrom int, iTo int) {
	const FROM = 0
	const TO = 1
	var i int = 0
	FromTo := [2]int{0, 0}
	for (FromTo[FROM] == 0) || (FromTo[TO] == 0) || (FromTo[FROM] == FromTo[TO]) {
		if i < len(FromTo) {
			if FromTo[i] == 0 {
				fmt.Println("Flight plan")
				for x, city := range global.Destinations {
					if x > 0 {
						fmt.Printf("%v. %v\n", x, city)
					}
				}
				switch i {
				case FROM:
					fmt.Print("Which city the flight departed from: ")
				case TO:
					fmt.Print("Which city the flight arrived at: ")
				}
				fmt.Scan(&FromTo[i])
				if FromTo[i] < 1 || FromTo[i] > len(global.Destinations)-1 || FromTo[FROM] == FromTo[TO] {
					if FromTo[FROM] == FromTo[TO] {
						FromTo[FROM] = 0
						FromTo[TO] = 0
					} else {
						FromTo[i] = 0
					}
				}
			}
		}
		i++
		if i >= len(FromTo) {
			i = 0
		}
	}
	return FromTo[FROM], FromTo[TO]
}

func Prompt4NoofTickets(iFrom, iTo int) uint {
	var Tickets uint = 0
	for (Tickets == 0) || ((Tickets > 0) && (Tickets+global.BookedTickets > global.ConferenceTickets)) {
		Tickets = 0
		fmt.Printf("How many %v-%v tickets do you want? ", global.Destinations[iFrom], global.Destinations[iTo])
		fmt.Scan(&Tickets)
	}
	return Tickets
}

//func PrintBookingList(list []string) {
//func PrintBookingList(list []map[string]string) {
func PrintBookingList(list []global.TBookData) {
	fmt.Println("Booking List:")
	//for i := 0; i < len(list); i++ {
	//	fmt.Printf("%v. %v\n", i+1, list[i])
	/*
		var iFR, iTO, iNUM int64
		for i, m := range list {
			iFR, _ = strconv.ParseInt(m["fromCity"], 10, 10)
			iTO, _ = strconv.ParseInt(m["toCity"], 10, 10)
			iNUM, _ = strconv.ParseInt(m["userTickets"], 10, 10)
			fmt.Printf("%v. %v, %v booked %v tickets from %v to %v\n", i+1, m["firstName"], m["lastName"], iNUM, global.Destinations[iFR], global.Destinations[iTO])
	*/
	for i, m := range list {
		fmt.Printf("%v. %v, %v booked %v tickets from %v to %v\n", i+1, m.FirstName, m.LastName, m.UserTickets, global.Destinations[m.FromCity], global.Destinations[m.ToCity])
	}
}

func PrintBookingReceipt(bookData global.TBookData) {
	fmt.Println("Receipt is being printing...")
	time.Sleep(10 * time.Second)
	sReceipt := fmt.Sprintf("\nBooking Receipt\n") +
		fmt.Sprintln("***************") +
		fmt.Sprintf("Customer: %v, %v\n", bookData.FirstName, bookData.LastName) +
		fmt.Sprintf("Departure from: %v\n", global.Destinations[bookData.FromCity]) +
		fmt.Sprintf("Arrive at: %v\n", global.Destinations[bookData.ToCity]) +
		fmt.Sprintf("Tickets: %v\n", bookData.UserTickets) +
		fmt.Sprintln("******END******")

	fmt.Print(sReceipt)
	fmt.Println("Thank you and please come again.")
	// Sync done: Decrements the WaitGroup counter by 1
	// So this is called by the gorountine to indicate that the job is finished
	global.WG.Done()
}
