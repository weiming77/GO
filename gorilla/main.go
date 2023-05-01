package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"gorilla/handlers"

	"github.com/gorilla/mux"
)

// Get the Bind address for the server from Environment variable
var bindAddress = os.Getenv("BIND_ADDRESS")

func main() {
	l := log.New(os.Stdout, "products-api ", log.LstdFlags)

	ph := handlers.NewProducts(l)

	//sm := http.NewServeMux()
	sm := mux.NewRouter()

	//sm.Handle("/", ph)
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MidlewareProductValidationProduct)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MidlewareProductValidationProduct)

	if len(bindAddress) == 0 {
		bindAddress = ":9090"
	}
	s := http.Server{
		Addr:         bindAddress,
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		l.Println("Starting server on port 9090...")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
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
