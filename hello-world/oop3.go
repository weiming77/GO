package main

import (
	"fmt"
	"math"
)

type IShape interface {
	area() float64
}

type TCircle struct {
	radius float64
}

type TRect struct {
	width  float64
	height float64
}

func (sender TCircle) area() float64 {
	return math.Pi * sender.radius * sender.radius
}

func (sender TRect) area() float64 {
	return sender.width * sender.height
}

func main() {
	c1 := TCircle{4.5}
	r1 := TRect{5, 7}
	shapes := []IShape{c1, r1}

	for _, shape := range shapes {
		fmt.Println(shape.area())
	}
}
