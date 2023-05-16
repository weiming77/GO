package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

const DBFILE = "../db/movies.db"

type Movie struct {
	id       int64
	title    string
	director string
	year     int
}

func init() {
	log.Println("Initialization")
	err := os.Setenv("CGO_ENABLED", "1")
	if err != nil {
		log.Fatal("Environment variable: CGO_ENABLED is not accessible.")
	}
	ienv, _ := strconv.Atoi(os.Getenv("CGO_ENABLED"))
	log.Printf("CGO_ENABLED=%d", ienv)
}

func main() {
	db, err := sql.Open("sqlite3", DBFILE)
	if err != nil {
		log.Fatal(err.Error())
	}

	rows, err := db.Query("Select * from Movies")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var m Movie
		rows.Scan(&m.id, &m.title, &m.director, &m.year)
		fmt.Println("Movie:", m.id, "Tille:", m.title, "director:", m.director, "At year:", m.year)
	}
}
