package main

import (
	"example.com/backend"
)

func main() {

	// go mod edit -replace example.com/backend=../backend
	a := backend.App{}
	a.Port = ":3030"
	a.Initialize()
	a.Run()
}
