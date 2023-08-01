package main

import (
	"fmt"
	"log"
	"os"

	"github.com/weiming77/GO/CDC/porting"
)

func main() {
	if porting.PortConfiguration.Tables == nil {
		log.Println("No data for porting!")
		os.Exit(1)
	}

	for i := 0; i < len(porting.PortConfiguration.Tables); i++ {
		fmt.Println(porting.PortConfiguration.Tables[i].Table, porting.PortConfiguration.Tables[i].File, porting.PortConfiguration.Tables[i].Schema)
	}
}
