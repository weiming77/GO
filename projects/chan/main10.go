package main

import (
	"fmt"
	"math/rand"
)

func generatorFunc[T any, K any](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)
	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()

	return stream
}

func main() {
	done := make(chan int)
	defer close(done)

	randNumFetcher := func() int { return rand.Intn(500000000) }
	for rando := range generatorFunc(done, randNumFetcher) {
		fmt.Println(rando)
	}
}
