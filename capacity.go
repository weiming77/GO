package main

import "fmt"

func main() {
	scores := make([]int, 0, 5)
	c := cap(scores)
	fmt.Printf("Capacity of %d\n", c)

	for i := 0; i < 25; i++ {
		scores = append(scores, i)

		// if our capacity has changed,
		// Go had to grow array to accomodate the necessary
		if cap(scores) != c {
			c = cap(scores)
			fmt.Printf("Capacity of %d\n", c)
		}
	}
}
