// pineline

package main

import "fmt"

func slice2Channel(nums []int) <-chan int {
	// return read only channel
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		// close the channel after the loop
		close(out)
	}()

	return out
}

// will take in the read only channel return from slice2Channel
func sq(in <-chan int) <-chan int {
	// return read only channel
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	//input
	nums := []int{2, 3, 4, 7, 11}

	// stage 1
	dataChannel := slice2Channel(nums)

	// stage 2
	finalChannel := sq(dataChannel)

	// stage 3
	for n := range finalChannel {
		fmt.Println(n)
	}
}
