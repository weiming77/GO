package greetings

import "fmt"

// hello returns a greeting for the named person.
func Hello(name string) string {
	// return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Merry Christmas!", name)
	//fmt.Println(message)
	return message
}
