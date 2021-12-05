package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")

	/*
		variable types:
		string
		 int,
		 float
		 boolean
	*/
	var name string = "Wei Ming"
	var temperature int = -14
	var age uint = 44 // unsign int aka cannot have -ve, only +ve
	var saving float64 = 1200.55

	adult := false // variable type declaration by initialization
	score := 0
	total := 0

	total++
	firstName := ""
	lastName := ""
	fmt.Printf("Enter your name: ")
	fmt.Scan(&firstName, &lastName)

	if (firstName + " " + lastName) == name {
		score += 1
	}

	total++
	var i uint
	fmt.Printf("\nHow old are you? ")
	fmt.Scan(&i)

	adult = (i >= 18) // == exact equal
	if adult {
		fmt.Printf("\n%s can proceed with game!", name)
		if i == age {
			score++
		}
	} else {
		fmt.Println("You're under age!")
		return // Exit function
	}

	total++
	fmt.Printf("\nWhat is temperature reading? ")
	ui := 0
	fmt.Scan(&ui)
	if ui == temperature {
		score++
	}

	total++
	fmt.Printf("\nHow much is your saving? ")
	f := 0.0000
	fmt.Scan(&f)
	if f == saving {
		score++
	}
	//	fmt.Printf("Name=%s", name)
	fmt.Printf("\nTemperature=%d", temperature)
	fmt.Printf("\nAge=%d\tadult is %v", age, adult)
	fmt.Printf("\nSaving=%f", saving)
	fmt.Printf("\n\nTotal score: %v%%", float64(score)/float64(total)*100)
}
