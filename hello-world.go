package main

import (
	"bufio"
	"fmt"
	"os"
	greetings "prj/mod/greetings"
)

func main() {
	// or you can initial array: arr:= [3]int{4,5,6}
	var arr [5]int

	// we can know the size of array: len(arr)
	// multi dimensional arrays: arr2D:= [2][3]int{{1,2,3}, {3,4,5}}
	/* Slice is another data type on its own:
	ie var s[]int = arr[:] :<-- slice operator
	PS: notice we did not mention the number of elements
	or we can do
	var a[]int = []int{5,6,7,8,9}
	a:= append(a, 10)

	a:= make([]int, 5)
	*/

	iLoop := 0
	for iLoop < 5 {
		arr[iLoop] = iLoop + 1
		switch iLoop {
		case 1:
			fmt.Println("one")
		case 3:
			fmt.Println("Three")
		default:
			fmt.Println("May be Zero")
		}
		/*
			switch {
			case iLoop > 4:
				fmt.Println("Greater than 4")
			case iLoop > 2:
				fmt.Println("Greater than 2")
			case iLoop > 0:
				fmt.Println("Greater than 0")
			}
		*/
		fmt.Println("Welcome to my quiz game!")
		iLoop++
	}

	fmt.Println(arr)
	for iLoop := 0; iLoop < 5; iLoop++ {
		if iLoop == 0 {
			continue
		} else if iLoop == 4 {
			break
		}
		fmt.Println(iLoop)
	}

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
	fmt.Printf("what is your full name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name = scanner.Text()
	fmt.Println(greetings.Hello(name))

	firstName := ""
	lastName := ""
	fmt.Printf("Confirm your first name, last name: ")
	fmt.Scan(&firstName, &lastName)

	// if {} else if {} else {}
	if (firstName + " " + lastName) == name {
		score += 1
	}

	// fmt.Println(greetings.Hello(firstName + " " + lastName))

	total++
	var i uint
	//var sAge string
	//var errmsg string
	fmt.Printf("\nHow old are you? ")
	//sAge = scanner.Text()
	//i, errmsg = strconv.ParseInt(sAge, 10, 64)
	fmt.Scan(&i)

	adult = (i >= 18) // == exact equal
	if adult {
		fmt.Printf("\n%s can proceed with game!", name)
		// &&=AND; ||=OR; !=NOT
		if (i == age) && adult {
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
