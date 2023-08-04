// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan string, 3)
	go func() {
		var wg sync.WaitGroup
		defer close(ch)

		wg.Add(1)
		go func(str string) {
			wg.Done()
			ch <- str
		}("世界")

		wg.Add(1)
		go func(str string) {
			wg.Done()
			ch <- str
		}("Hello")

		wg.Wait()
	}()

	for s := range ch {
		fmt.Println(s)
	}
}
