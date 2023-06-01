package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// record is weather record
type Record struct {
	Time    time.Time
	Station string
	Temp    float64 `json:"temperature"` // celsius
	Rain    float64
}

func readRecord(r io.Reader) (Record, error) {
	var rec Record
	dec := json.NewDecoder(r)
	if err := dec.Decode(&rec); err != nil {
		return Record{}, nil
	}

	return rec, nil
}

func MainFunc() {
	file, err := os.Open("Record.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rec, err := readRecord(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", rec)

}
