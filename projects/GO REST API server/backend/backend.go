package backend

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

const DBFILE = "../db/movies.db"

type App struct {
	DB   *sql.DB
	Port string
}

func (a *App) Initialize() {
	var err error
	a.DB, err = sql.Open("sqlite3", DBFILE)

	if err != nil {
		log.Fatal(err.Error())
	}
}

type Movie struct {
	id       int64
	title    string
	director string
	year     int
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func (a *App) Run() {
	http.HandleFunc("/", helloWorld)
	fmt.Println("Server started and listening on port", a.Port)
	log.Fatal(http.ListenAndServe(a.Port, nil))
}
