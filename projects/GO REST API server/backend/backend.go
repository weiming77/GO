package backend

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

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
	// This is how one end point react differently to different APIs method
	http.Handle("/", a.Router)
}

func (a *App) Run() {
	fmt.Println("Server started and listening on port", a.Port)
	log.Fatal(http.ListenAndServe(a.Port, a.Router))
}
