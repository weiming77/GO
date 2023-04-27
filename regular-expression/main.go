package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

/*
12 shares of MSFT for $234.57
10 shares of TSLA for $692.40
*/
var transRe = regexp.MustCompile(`(\d+) shares of ([A-Z]+) for \$(\d+(\.\d+)?)`)

// for your information: https://pkg.go.dev/regexp/syntax

// translation is a b
type Transaction struct {
	Symbol string
	Volumn int
	Price  float64
}

func parseLine(line string) (Transaction, error) {
	matches := transRe.FindStringSubmatch(line)
	if matches == nil {
		return Transaction{}, fmt.Errorf("Bad line: %q", line)
	}
	var t Transaction
	t.Symbol = matches[2]
	t.Volumn, _ = strconv.Atoi(matches[1])
	t.Price, _ = strconv.ParseFloat(matches[3], 64)
	return t, nil
}
func main() {
	lines := []string{
		"12 shares of MSFT for $234.57",
		"10 shares of TSLA for $692.40",
	}

	var t Transaction
	var err error
	for _, line := range lines {
		t, err = parseLine(line)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", t) // {symbol:MSFT Volumn:12 Price:234.57}
	}
}
