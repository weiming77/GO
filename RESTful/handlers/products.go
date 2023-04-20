package handlers

import (
	"log"
	"net/http"

	"github.com/weiming77/GO/RESTful/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// handler function return products' information
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}
	// otherwise catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// Internal function return products' information
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	// return the product list at JSON string
	lp := data.GetProducts()
	// d, err := json.Marshal(lp) // resource consuming
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
	// rw.Write(d)
}
