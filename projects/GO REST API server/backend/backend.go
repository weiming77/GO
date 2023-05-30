package backend

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	//use gorila mux when response differently based on API method call
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

const DBFILE = "../db/movies.db"

type App struct {
	DB     *sql.DB
	Port   string
	Router *mux.Router
}

func (a *App) Initialize() {
	DB, err := sql.Open("sqlite3", DBFILE)

	if err != nil {
		log.Fatal(err.Error())
	}

	a.DB = DB

	// Ideally we want differert API methods handled by different GO methods.
	// In other words we want different request handlers for one model ie products, orders
	// to accomplish this, we would use a router. a router keep track of what code to execute (controller) based on the API endpoint and method called.
	// I like to use router for better code layout, as I can group together API handlers base on the resource (model) they affect
	// use router so http server response according to the APIs method used.
	a.Router = mux.NewRouter()
	a.InitializeRouters()
}

type Movie struct {
	id       int64
	title    string
	director string
	year     int
}

// default root handle
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

// This is get handle - Data retrieve/view
func getRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a GET")
}

// This is a POST handle - New data
func postRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a POST")
}

// This is a PUT handle - Update several/partial data
func putRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a PUT")
}

// This is a PATCH handle - update several/partial data
func patchRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a PATCH")
}

// This is a DELETE handle - update several/partial data
func deleteRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a DELETE")
}

// Helper function with Messaage
func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(response))
}

// helper function with Error
func responseWithError(w http.ResponseWriter, code int, message string) {
	responseWithJSON(w, code, map[string]string{"error": message})
}

// This is GET handle - get all the products
func (a *App) getAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := getProducts(a.DB)
	if err != nil {
		fmt.Println("getProducts error: %s\n", err.Error())
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseWithJSON(w, http.StatusOK, products)
}

// this is GET handle - get a product by ID
func (a *App) fetchProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		fmt.Println("getProduct error: ID:%s\n", id)
		responseWithError(w, http.StatusInternalServerError, fmt.Errorf("Expecting product id but value %s detected!\n", id).Error())
		return
	}

	var p product
	p.ID, _ = strconv.Atoi(id)
	err := p.getProduct(a.DB)

	if err != nil {
		fmt.Println("getProduct error: %s\n", err.Error())
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseWithJSON(w, http.StatusOK, p)
}

func (a *App) newProduct(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var p product

	fmt.Printf("newProduct read error: %q\n", reqBody)
	if err := json.Unmarshal(reqBody, &p); err != nil {
		fmt.Println("newProduct read error: %s\n", err.Error())
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := p.createProduct(a.DB); err != nil {
		fmt.Println("newProduct error: %s\n", err.Error())
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseWithJSON(w, http.StatusOK, p)
}

func (a *App) deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	if id == "" {
		fmt.Println("deleteProduct error: ID:%s\n", id)
		responseWithError(w, http.StatusInternalServerError, fmt.Errorf("Expecting product id but value %s detected!\n", id).Error())
		return
	}

	var p product
	p.ID, _ = strconv.Atoi(id)
	err := p.deleteProduct(a.DB)

	if err != nil {
		fmt.Println("deleteProduct error: %s\n", err.Error())
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseWithJSON(w, http.StatusOK, p)
}

func (a *App) allOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := getOrders(a.DB)
	if err != nil {
		fmt.Printf("allOrders error: %s\n", err.Error())
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseWithJSON(w, http.StatusOK, orders)
}

func (a *App) fetchOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var o order
	o.ID, _ = strconv.Atoi(id)
	err := o.getOrder(a.DB)
	if err != nil {
		fmt.Printf("getOrder error: %s\n", err.Error())
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseWithJSON(w, http.StatusOK, o)
}

func (a *App) newOrder(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var o order
	json.Unmarshal(reqBody, &o)

	err := o.createOrder(a.DB)
	if err != nil {
		fmt.Printf("newOrder error: %s\n", err.Error())
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	for _, item := range o.Items {
		var oi orderItem
		oi = item
		oi.OrderID = o.ID
		err := oi.createOrderItem(a.DB)
		if err != nil {
			fmt.Printf("newOrder, newOrderItem error: %s\n", err.Error())
			responseWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		responseWithJSON(w, http.StatusOK, o)
	}
}

func (a *App) newOrderItems(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var ois []orderItem
	json.Unmarshal(reqBody, &ois)
	//fmt.Printf("newOrderItems %s", []byte(reqBody))

	for _, item := range ois {
		var oi orderItem
		oi = item
		//fmt.Printf("newOrderItem %d %d", oi.OrderID, oi.ProductID)
		err := oi.createOrderItem(a.DB)
		if err != nil {
			fmt.Printf("newOrderItem error: %s\n", err.Error())
			responseWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	responseWithJSON(w, http.StatusOK, ois)
}

// This is a POST handle - create new product
func (a *App) InitializeRouters() {

	// APPSERVER is returning "hello world" when request match the end point in regardless of different REST methods
	//http.HandleFunc("/", helloWorld)
	// curl -i http://localhost:3030
	a.Router.HandleFunc("/", getRequest).Methods(http.MethodGet)
	// curl -i -X POST http://localhost:3030
	a.Router.HandleFunc("/", postRequest).Methods(http.MethodPost)
	// curl -i -X PUT http://localhost:3030
	a.Router.HandleFunc("/", putRequest).Methods(http.MethodPut)
	// curl -i -X PATCH http://localhost:3030
	a.Router.HandleFunc("/", patchRequest).Methods(http.MethodPatch)
	// curl -i -X DELETE http://localhost:3030
	a.Router.HandleFunc("/", deleteRequest).Methods(http.MethodDelete)
	// curl -i -X GET localhost:3030/products
	a.Router.HandleFunc("/products", a.getAllProducts).Methods(http.MethodGet)
	// curl -i -X GET localhost:3030/product/5
	a.Router.HandleFunc("/product/{id}", a.fetchProduct).Methods(http.MethodGet)
	// curl -i -X POST localhost:3030/products -H "Content-Type: application/json" -d "{\"ProductCode\":\"PRTN\",\"Name\":\"Proton\",\"Inventory\":50,\"Price\":49.55,\"Status\":\"A\"}"
	a.Router.HandleFunc("/products", a.newProduct).Methods(http.MethodPost)
	// curl -i -X DELETE localhost:3030/product/5
	a.Router.HandleFunc("/product/{id}", a.deleteProduct).Methods(http.MethodDelete)
	// curl -i -X GET localhost:3030/orders
	a.Router.HandleFunc("/orders", a.allOrders).Methods(http.MethodGet)
	// curl -i -X GET localhost:3030/order/6
	a.Router.HandleFunc("/order/{id}", a.fetchOrder).Methods(http.MethodGet)
	// curl -X POST localhost:3030/orders -H "Content-Type: application/json" -d "{\"customerName\": \"Daisy Duck\", \"total\": 30, \"status\": \"Shipped\", \"items\": [{\"product_id\": 2, \"quantity\": 1}, {\"product_id\": 3, \"quantity\": 3}]}"
	a.Router.HandleFunc("/orders", a.newOrder).Methods(http.MethodPost)
	// curl -X POST localhost:3030/orders -H "Content-Type: application/json" -d "{\"customerName\": \"Daisy Duck\", \"total\": 30, \"status\": \"Shipped\", \"items\": []}"
	// curl -X POST localhost:3030/orderitems -H "Content-Type: application/json" -d "[{\"order_id\": 3, \"product_id\": 2, \"quantity\": 1}, {\"order_id\": 3, \"product_id\": 3, \"quantity\": 3}]"
	a.Router.HandleFunc("/orderitems", a.newOrderItems).Methods(http.MethodPost)

	//
	// This is how one end point react differently to different APIs method
	// http.Handle("/", a.Router)
}

func (a *App) Run() {
	fmt.Println("Server started and listening on port", a.Port)
	log.Fatal(http.ListenAndServe(a.Port, a.Router))
}
