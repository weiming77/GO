package main

import (
	"fmt"
	"sync"
	"time"
)

// Task definition
type Task interface {
	Process()
}

// Email task defination
type EmailTask struct {
	Email       string
	Subject     string
	MessageBody string
}

// A way to process the Email Task
func (t *EmailTask) Process() {
	fmt.Printf("Sending email to %s\n", t.Email)
	// Simulate a time consuming process
	time.Sleep(2 * time.Second)
}

// Image task defination
type ImageProcessingTask struct {
	ImageUrl string
}

// A way to process Image Task
func (t *ImageProcessingTask) Process() {
	fmt.Printf("Processing image %s\n", t.ImageUrl)
	// Simulate a time consuming process
	time.Sleep(3 * time.Second)
}

// Define the worker pool
type WorkerPool struct {
	Tasks       []Task
	concurrency int
	taskChan    chan Task      // define a channel to send tasks to workers
	wg          sync.WaitGroup // a WaitGroup to synchronizing the completion of tasks, this will be used to wait for the tasks to be completed.
}

// functions to execute tasks in the worker pool
// - We need a method that receives tasks from the task Channel and processes them.
// - Let's define a worker method on the worker pool struct
func (wp *WorkerPool) worker() {
	for task := range wp.taskChan {
		// call the process method on the task to process it
		task.Process()
		// Now we will signal the completion of a task using the WaitGroup
		wp.wg.Done()
	}
}

// Now next come the run method initialize the channel
// set the concurrency creates the goroutines and
// sends tasks over the channel.
// we will firth initialize the tasks with the capacity equal to the number of the tasks.
func (wp *WorkerPool) Run() {
	// initialize the tasks channel
	wp.taskChan = make(chan Task, len(wp.Tasks))

	// let's start the worker goroutine number of workers are set to the concurrency
	for i := 0; i < wp.concurrency; i++ {
		go wp.worker()
	}

	// add the number of tasks to the waitgroup
	wp.wg.Add(len(wp.Tasks))
	// Send tasks to the tasks channel
	for _, task := range wp.Tasks {
		wp.taskChan <- task
	}

	// close the task channel after sending all tasks to signal,
	// no more tasks will be send
	close(wp.taskChan)

	// wait for all tasks to be finished
	wp.wg.Wait()
}
