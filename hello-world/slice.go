package main

import "fmt"

func test() {
	fmt.Println("test me")
}

func greet(someone string) {
	fmt.Printf("Hello %s\n", someone)
}

// pass function as parameter
func test2(myFunc func(x int, y int) int) {
	fmt.Println(myFunc(7, 3))
}

// you can return multiple result func add(x, y int) (string, int) {
// return "try me", x+y
// }
func add(x int, y int) int {
	return x + y
}

func tellMe(first, last string) (fullname string) {
	fullname = first + ", " + last
	return
}
func main() {
	var a []int = []int{1, 3, 4, 56, 7, 12, 4, 6}
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	for i, element := range a {
		fmt.Printf("%d: %d\n", i, element)
	}

	// _ is anonymous variable
	for _, element := range a {
		fmt.Printf("%d\n", element)
	}

	// reference of where value of variable x is stored
	x := 7
	// this show the memory address where value x stored
	fmt.Println(&x)
	// now y refer to variable x pointer
	y := &x
	*y = *y * 2
	fmt.Println(x)

	// maps also mean key value pair
	// string key point to integer value
	// "apple":1
	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}
	// or
	mp = make(map[string]int)
	//access the value of "apple"
	fmt.Println(mp["apple"])
	//or change the value
	mp["apple"] = 99
	// add new
	mp["pineapple"] = 80
	// delete the key pair
	delete(mp, "apple")
	// check the key exist
	val, ok := mp["apple"]
	if ok {
		fmt.Println(val)
	}

	fmt.Println(mp)
	test()
	greet("wei Ming")
	p := greet
	p("Angie")
	add(1, 3)

	inner := func() {
		fmt.Println("Inner function")
	}
	inner()

	test2(add)
}
