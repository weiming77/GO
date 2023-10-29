package porting

import (
	"bufio"
	"log"
	"os"
	"strings"

	//"sync"
	"sync/atomic"
	"time"
)

const (
	BATCHSIZE  = 1000
	BATCHBREAK = "GO"
)

// It could be scripts, csv or json file..
type IDataFileReader interface {
	GetDataFile() string
	ReadDataFile()
}

type Lines []string

type ChunkData struct {
	SeqNo uint32 `json:"seqNo"`
	Data  *Lines `json:"lines"`
}

type DataFileReader struct {
	DataFile  string
	GoByBatch bool
	Content   []*ChunkData
	SeqGen    atomic.Uint32
	//mu        sync.Mutex
	lg          *log.Logger
	iTimer      *int32
	iBatchCount uint
}

func NewDataFileReader(fileName string, byBatch bool, l *log.Logger) *DataFileReader {
	return &DataFileReader{
		DataFile:  fileName,
		GoByBatch: byBatch,
		Content:   []*ChunkData{},
		lg:        l,
	}
}

func (r *DataFileReader) NextSeq() uint32 {
	return r.SeqGen.Add(1)
}

func (r *DataFileReader) Countdown(i int32) {
	atomic.AddInt32(r.iTimer, i)
}

func (r *DataFileReader) GetDataFile() string {
	return r.DataFile
}

func (r *DataFileReader) processData(d Lines, ch chan<- uint32, done chan<- int) {
	id := r.NextSeq()
	ch <- id
	timestamp := time.Now()

	//r.mu.Lock()
	r.Content = append(r.Content, &ChunkData{SeqNo: id, Data: &d})
	//r.mu.Unlock()

	s := time.Since(timestamp)
	r.lg.Printf("Chunk ID %d with %d line(s) processed in %v\n", id, len(d), s)

	done <- int(s)
}

func (r *DataFileReader) ReadDataFile() error {
	file, err := os.Open(r.DataFile)
	if err != nil {
		//r.lg.Fatal(err)
		return err
	}
	defer file.Close()

	ch := make(chan uint32, 5)
	done := make(chan int)
	start := time.Now()

	var (
		iLn          uint   = 0
		iAccumulated uint64 = 0
		dat          Lines  = []string{}
	)

	reader := bufio.NewScanner(file)
	for reader.Scan() {
		if (!r.GoByBatch && strings.EqualFold(reader.Text(), BATCHBREAK)) ||
			(r.GoByBatch && iLn == BATCHSIZE) {
			go r.processData(dat, ch, done)

			dat = []string{}
			iLn = 0
		}

		if !strings.EqualFold(reader.Text(), BATCHBREAK) {
			iLn += 1
			dat = append(dat, reader.Text())
		}
	}
	//if err := reader.Err(); err != nil {
	//	log.Fatal(err)
	//}

	// just in case we have remaining
	if len(dat) > 0 {
		go r.processData(dat, ch, done)
	}

	var bNoMore bool
	for !bNoMore {
		select {
		case iTask, ok := <-ch:
			if !ok {
				bNoMore = true
				break
			}
			r.Countdown(1)
			r.iBatchCount += 1
			r.lg.Printf("Task %d is completed.\n", iTask)

		case iSec := <-done:
			iAccumulated += uint64(iSec)
			r.Countdown(-1)

		default:
			if r.iBatchCount > 0 && *r.iTimer == 0 {
				r.lg.Printf("Total %d batch(es) with total processing %d secs completed in %v.\n", r.iBatchCount, iAccumulated, time.Since(start))
				close(ch)
			}
		}
	}

	return nil
}
