package main

import (
	"fmt"
)

func connect() error {
	// attempt to connect to server
	if true {
		panic("Fail connect to server on purpose!")
	}

	return nil
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	err := connect()
	if err != nil {
		fmt.Println("Error connecting to server:", err)
	}

	fmt.Println("The job is done finally")
}
