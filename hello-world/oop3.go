package main

import (
	"fmt"
	"math"
	"reflect"
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

// reflection made of three concept: TYPE, KIND, VALUE
func CalculateArea(s interface{}) {
	// 1. TYPE (type & variables)
	switch tp := s.(type) {
	case TCircle:
		fmt.Printf("I am a circle and my value is %v\n", tp)
	case TRect:
		fmt.Printf("I am a rectangle and my value is %v\n?", tp)
	default:
		fmt.Println("Non of the above!")
	}

	//1. TYPE
	myType := reflect.TypeOf(s)
	// query the type
	fmt.Println("Reflect.Type=", myType.Name())
	fmt.Println("Reflect.Kind=", myType.Kind())
	// only Struct kind has fields
	fmt.Println(myType.Name(), "has", myType.NumField(), "fields")
	//fmt.Println(myType.Elem().Name())
	for i := 0; i < myType.NumField(); i++ {
		fmt.Println(i, ". Field=", myType.Field(i).Name, "of data type", myType.Field(i).Type)
	}
}

func main() {
	c1 := TCircle{4.5}
	r1 := TRect{5, 7}
	shapes := []IShape{c1, r1}

	for _, shape := range shapes {
		fmt.Println(shape.area())
	}

	CalculateArea(c1)
	CalculateArea(r1)
}
