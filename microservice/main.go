package main

import (
	"context"
	"log"
	"microservice/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
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

	// lets put everyhing here
	s := &http.Server{
		Addr:         ":9090",
		Handler:      mux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	log.Println("Starting server on :9090")
	// we can test the program at DOS like
	// C:\curl localhost:9090 -d 'someone'
	// C:\curl localhost:9090/goodbye

	// refactor and wrap it up in a go func therefore it will start serve service
	// and would not block the program. It also start a shutdown
	go func() {
		err := s.ListenAndServe()
		//err := http.ListenAndServe(":9090", mux)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// so what we can use the OS signal package and in the OS package
	// we can register for the notification that notify takes a channel.
	// Create a channel that going to receive the signal.
	sigChan := make(chan os.Signal)
	// signal broacast Interrupt signal to sigChan
	signal.Notify(sigChan, os.Interrupt)
	// follow by broacast kill command signal...
	signal.Notify(sigChan, os.Kill)

	// PS: Allow application to close any work that's going on in the handler. ie
	// Close the opened database connection,
	// Wait and finish up any large upload,
	// Finish any communication to another service.

	// block happens starts at here because reading from a channel will be blocked
	// until there's a message available to be consumed.
	signal := <-sigChan
	l.Println("Received terminate, graceful shutdown", signal)
	// Preparation for graceful shutdown, ceate an 30 secs timeout deadline context
	// meaning allow 30 secs attempt to gracefully shutdown all the working handlers
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	// gracefully shutdown will wait until the request that are currently
	// handled by the server have completed it, at the same time immediately
	// no longer accept anymore new request(s) but it will wait until everybody
	// finish their works then program exit without interruption.
	s.Shutdown(tc)
}
