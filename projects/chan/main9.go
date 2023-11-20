package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync/atomic"
	"time"
)

type TTask struct {
	iID   int
	iTime int
}

func planTasks(CPU, noofTasks int) <-chan TTask {
	res := make(chan TTask, CPU)

	go func(tasks int) {
		for i := 0; i < tasks; i++ {
			// pretend as an busy go-routine
			iSec := rand.Intn(5) + 1
			fmt.Printf("Task %d will take %d secs for processing...\n", i+1, iSec)
			res <- TTask{iID: i + 1, iTime: iSec}
		}
		close(res)
	}(noofTasks)

	return res
}

func doTasks(CPU int, tasks <-chan TTask) <-chan TTask {
	var counter atomic.Int32
	done := make(chan bool, CPU)

	res := make(chan TTask, CPU)
	// sender
	for t := range tasks {
		counter.Add(1)

		go func(t TTask, d chan<- bool) {
			time.Sleep(time.Duration(t.iTime) * time.Nanosecond)
			fmt.Printf("Task %d is taking %v secs...\n", t.iID, t.iTime)
			res <- t
			// Everything is get done
			d <- true
		}(t, done)
	}

	// receiver
	go func() {
		for {
			select {
			case <-done:
				counter.Add(-1)
			default:
				if counter.Load() == 0 {
					fmt.Println("All tasks are completed!")
					close(res)
					return
				}
			}
		}
	}()

	return res
}

func showTasks(r <-chan TTask, iAccumulated *int64) {
	for {
		select {
		case t, ok := <-r:
			if !ok {
				fmt.Println("Summary is ready.\n")
				return
			}

			*iAccumulated += int64(t.iTime)
			fmt.Printf("Task %d is completed in %v secs.\n", t.iID, t.iTime)
		default:
			//fmt.Println("Program is falling asleep...")
		}
	}
}

func main() {
	const NOOF_TASKS = 100000
	var iAccumulated int64
	iCPU := runtime.NumCPU()
	now := time.Now()

	chTasks := planTasks(iCPU, NOOF_TASKS)
	chResult := doTasks(iCPU, chTasks)
	showTasks(chResult, &iAccumulated)

	fmt.Printf("Total %d tasks with total processing %d secs completed in %v.\n", NOOF_TASKS, iAccumulated, time.Since(now))
}
