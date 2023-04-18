package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/weiming77/GO/microservice/handlers"
)

type customHandler struct {
	Message string
}

func (c customHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Hi I am a custom handler func"))
	// or you can display the content from customHandler
	// w.Write([]byte(c.Message))
	j, _ := json.Marshal(&c)
	w.Header().Set("content-type", "application/json")
	w.Header().Add("Tester", "Wei Ming")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handlers.Hello(l)
	gh := handlers.GoodBye(l)

	var c customHandler
	c.Message = "Hello Handler"
	// Register the two new handler functions and corresponding URL patters with
	// the servemux, in exactly the same way that we did before
	mux := http.NewServeMux()
	// converting the function into a handler type then
	// register it to a thing call MUX, server multiplexer
	mux.Handle("/", hh)
	mux.Handle("/goodbye", gh)
	mux.Handle("/getData", c)
	//mux.Handle("/getData", customHandler{Message: "This is a bloody testing for handler"})
	err := http.ListenAndServe(":9090", mux)
	l.Fatal(err)
}
