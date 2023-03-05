package main

import "fmt"

type TStudent struct {
	name   string
	grades []int
	age    uint
}

func (sender *TStudent) getAge() int {
	return int(sender.age)
}
func (sender *TStudent) setAge(value int) {
	sender.age = uint(value)
}
func (sender *TStudent) getAverageGrade() float32 {
	sum := 0
	for _, v := range sender.grades {
		sum += v
	}
	return float32(sum) / float32(len(sender.grades))
}

func main() {
	s1 := TStudent{"Tim", []int{70, 80, 90, 99}, 19}
	s1.getAge()

	fmt.Println(s1.getAge())
	s1.setAge(7)
	fmt.Println(s1.getAge())
	fmt.Println(s1.getAverageGrade())
}
