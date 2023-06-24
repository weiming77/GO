package handlers

import (
	"context"
	"fmt"
	"gorilla/data"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

/*
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// GET when read data
		// handler function return products' information
		// curl -v localhost:9090| jq
		p.getProducts(rw, r)
		return
	} else if r.Method == http.MethodPost {
		// POST when create new data
		// test at command prompt:
		// curl -v localhost:9090 -d '{"id":1, "name": "green tea", "description": "a cup of tea"}'
		p.addProduct(rw, r)
		return
	} else if r.Method == http.MethodPut {
		// PUT when update data and replace the entire data
		// expect the ID in the URL
		// curl -v localhost:9090/1 -XPUT -d '{"name": "green tea", "description": "a cup of green tea"}'
		reg := regexp.MustCompile("/([0-9]+)")
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "Extract product ID error", http.StatusBadRequest)
			return
		}
		// curl -v localhost:9090/1 -X PUT | jq
		//p.l.Println("Got ID:", id)

		p.updateProducts(id, rw, r)
		return
	} else if r.Method == http.MethodPatch {
		// PATCH when update a few data fields

	}
	// otherwise catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}
*/
// Internal function return products' information
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET products")
	// return the product list at JSON string
	lp := data.GetProducts()
	// d, err := json.Marshal(lp) // resource consuming
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
	// rw.Write(d)
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST products")
	// we are getting the product from Request.
	// The KeyProduct will return us an interface so we need to cast this, because
	// it is impossible for anything to get into this without going thru the middleware first
	// and middleware ensuring that is happening.
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	// create and add new product to product list
	// this block being replaced by the middleware above
	//prod := &data.Product{}
	//err := prod.FromJSON(r.Body)
	//if err != nil {
	//	http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
	//}
	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(&prod)
}

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable convert id: "+vars["id"], http.StatusBadRequest)
		return
	}
	p.l.Println("Handle PUT/Update products", id)

	// we are getting the product from Request.
	// The KeyProduct will return us an interface so we need to cast this, because
	// it is impossible for anything to get into this without going thru the middleware first
	// and middleware ensuring that is happening.
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	// update the product by ID
	// prod := &data.Product{}
	// PS: This is being replaced by the middleware above
	//err = prod.FromJSON(r.Body)
	//if err != nil {
	//	http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
	//	return
	//}
	//prod.ID = id

	p.l.Printf("Prod: %#v", prod)
	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(rw, "Unknown error", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct {
}

// MidlewareValidationProduct validate the product in the request and call the next if ok
func (p *Products) MidlewareProductValidationProduct(next http.Handler) http.Handler {
	// middleware will execute before the actual handler
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[Error] deserializing product", err)
			http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
			return
		}

		// validate the product by sanitize the inputs
		// curl -X POST localhost:9090 -d "{\"name\":\"\", \"description\":\"free sample\", \"price\":0.00, \"sku\":\"Na\"}"
		err = prod.Validate()
		if err != nil {
			p.l.Println("[Error] validating product", err)
			http.Error(rw, fmt.Sprintf("Error validating product: %s\n", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
