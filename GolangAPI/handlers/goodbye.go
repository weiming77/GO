package handlers

import (
	"log"
	"net/http"
)

type IGoodBye struct {
	l *log.Logger
}

func GoodBye(l *log.Logger) *IGoodBye {
	return &IGoodBye{l}
}

// add the interface
func (g IGoodBye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.l.Println("goodbye handler function...")
	w.Write([]byte("Good Bye and see ya"))
}
