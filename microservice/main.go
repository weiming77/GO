package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// add a hello handler
func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello handler function...")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Opps", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello %s", d)
}

// add a goodbye handle
func goodbye(w http.ResponseWriter, r *http.Request) {
	log.Println("goodbye handler function...")
	w.Write([]byte("Good Bye and see ya"))
}

func main() {
	// Register the two new handler functions and corresponding URL patters with
	// the servemux, in exactly the same way that we did before
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/goodbye", goodbye)

	log.Println("Starting server on :9090")
	http.ListenAndServe(":9090", mux)
	// err := http.ListenAndServe(":9090", mux)
	// log.Fatal(err)
}
