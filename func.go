package main

import (
	"fmt"
	"strconv"
)

type TOPFunc func(int, int) int

func add(i int, j int) int { return i + j }

func sub(i int, j int) int { return i - j }

func mul(i int, j int) int { return i * j }

func div(i int, j int) int { return i / j }

func anonymous_func(cnt int) {
	for i := 0; i < cnt; i++ {
		func(j int) {
			fmt.Println("Printing", j, "from inside the anonymous function")
		}(i)
	}
}

func intSeq() func() int {
	i := 7
	return func() int {
		i++
		return i
	}
}

//var opMap := map[string]func(int, int) int{
var opMap = map[string]TOPFunc{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func main() {
	expressions := [][]string{
		[]string{"2", "+", "9"},
		[]string{"4", "-", "3"},
		[]string{"6", "*", "5"},
		[]string{"8", "/", "4"},
		[]string{"two", "+", "three"},
	}
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("Invalid expression", expression)
			continue
		}
		op := expression[1]
		opfunc, ok := opMap[op]
		if !ok {
			fmt.Println("Unsupported operation", op)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opfunc(p1, p2)
		fmt.Println(result)
	}

	anonymous_func(5)

	closure_func := intSeq()
	fmt.Println(closure_func())
}
