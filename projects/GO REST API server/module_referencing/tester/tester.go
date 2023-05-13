package main

import (
	"example.com/practice"
)

func main() {
	// Since we are yet to publish the practice pakcage, We will run into the error:
	// no required module provides package example.com/practice; to add it: go get example.com/practice
	// when we go run tester.go
	// Resolution is to give go.mod some helps via command
	// go mod edit -replace example.com/practice=../practice
	// go mod tidy
	practice.Test()
}
