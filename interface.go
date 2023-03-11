package main

import "fmt"

// An interface is an ABSTRACT type
type IBackEndDeveloper interface {
	ProgramInGo()
	BuildBackEnd()
	TrainNewEmployees()
}

// A Struct is a CONCRETE type
type TSoftwareEngineer struct {
	ProgrammingInPython()
	BuildBackEnd()
	TrainNewEmployees()
}

func (Sender TSoftwareEngineer) ProgrammingInPython() {
	fmt.Println("Programming in Python")
}

func (Sender TSoftwareEngineer) BuildBackEnd() {
	fmt.Println("Building backend Application server")
}

func (Sender TSoftwareEngineer) TrainNewEmployees() {
	fmt.Println("Training New Employee")
}

func main() {

}