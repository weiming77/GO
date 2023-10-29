package main

import (
	"log"
	"os"

	"github.com/weiming77/GO/CDC/porting/config"
	porting "github.com/weiming77/GO/CDC/porting/reader"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func main() {
	ch := make(chan uint)
	reader := make(chan *porting.DataFileReader)

	//l := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Logging is not avaiable!")
	}
	defer f.Close()

	// basically output your logs to any destination that implements the io.Writer interface,
	// so you have a lot of flexibility when deciding where to log messages in your application.
	log.SetOutput(f)

	InfoLogger = log.New(f, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(f, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(f, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)

	if config.PortConfiguration.Tables == nil {
		ErrorLogger.Println("No data for porting!")
	}

	for i := 0; i < len(config.PortConfiguration.Tables); i++ {
		if config.PortConfiguration.Tables[i].Selected {
			InfoLogger.Println(config.PortConfiguration.Tables[i].Table, config.PortConfiguration.Tables[i].Data)

			go func(f, t string, c chan<- *porting.DataFileReader, d chan<- uint) {
				d <- 1 // start job count
				InfoLogger.Println("About to spawn file reader")
				r := porting.NewDataFileReader(f, true, InfoLogger)
				if err := r.ReadDataFile(); err != nil {
					ErrorLogger.Printf("process %s fail due to %v\n", t, err)
					c <- &porting.DataFileReader{DataFile: f} // just simply return empty
					return                                    // job done with fail status
				}
				InfoLogger.Println("File reader has done its'' work")
				c <- r // job done with success status

			}(config.PortConfiguration.Tables[i].Data, config.PortConfiguration.Tables[i].Table, reader, ch)
		}
	}

	var (
		bNoMore         bool
		iNoofJob, iDone uint
	)
	for !bNoMore {
		select {
		case p, ok := <-reader:
			if !ok {
				bNoMore = true
				break
			}
			iDone += 1
			InfoLogger.Printf("%s has %d batch(es).\n", p.DataFile, len(p.Content))

		case i := <-ch:
			iNoofJob += i

		default:
			if iNoofJob == iDone {
				close(reader)
			}
		}

	}
}
