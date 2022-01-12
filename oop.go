package main

import "fmt"

// custom type
type TPoint struct {
	x int32
	y int32
}

type TCircle struct {
	radius float64
	center *TPoint
}

func change(pt *TPoint) {
	pt.x = 100
}

func main() {
	var p1 TPoint = TPoint{1, 2}
	var p2 TPoint = TPoint{3, 4}
	fmt.Println(p1.x)
	fmt.Println(p1.y)
	p2.x = p2.x * 2
	p2.y = p2.y * p2.x
	fmt.Println(p2)

	// also you can set value any one of them
	p3 := TPoint{x: 3}
	fmt.Println(p3)

	fmt.Println(p1.x)
	change(&p1)
	fmt.Println(p1.x)

	c1 := TCircle{4.567, &TPoint{50, 100}}
	fmt.Println(c1.center.x)
}
