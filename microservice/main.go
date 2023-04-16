package main

import (
	"log"
	"microservice/handlers"
	"net/http"
	"os"
	// "github.com/weiming77/GO/tree/master/microservice/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.Hello(l)
	gh := handlers.GoodBye(l)

	// Register the two new handler functions and corresponding URL patters with
	// the servemux, in exactly the same way that we did before
	mux := http.NewServeMux()
	// converting the function into a handler type then
	// register it to a thing call MUX, server multiplexer
	mux.Handle("/", hh)
	mux.Handle("/goodbye", gh)

	log.Println("Starting server on :9090")
	http.ListenAndServe(":9090", mux)
	err := http.ListenAndServe(":9090", mux)
	log.Fatal(err)
}
