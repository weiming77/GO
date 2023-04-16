// add package
package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Create a struct which implements the interface HTTP handler
// Avoid create concrete object inside the handler, the reason
// testability (unit test) and use technique called dependency injection
type IHello struct {
	l *log.Logger
}

// return hello handler as a reference
func Hello(l *log.Logger) *IHello {
	return &IHello{l}
}

// add the interface that satisfies the HTTP handler interface
// PS: r is reference to an HTTP.Request
func (h *IHello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// here is what need to sactisfy the HTTP handler Interface
	h.l.Println("Hello handler function...")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Opps", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello %s", d)
}
