package main

import "fmt"

func fSlice(c *[]string) {
	(*c)[len((*c))-1] = "Let''s me wrap this up"
	*c = append(*c, "Good bye")
}

func main() {
	var comments = []string{}
	comments = append(comments, "Hello, 世界")
	comments = append(comments, "I have no idea what to say!")
	fSlice(&comments)
	fmt.Println(comments)
}
