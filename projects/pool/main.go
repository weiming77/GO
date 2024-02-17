package main

import (
	"fmt"
)

func main() {
	// create new tasks
	tasks := make([]Task, 27)
	for i := 0; i < len(tasks); i++ {
		tasks[i] = Task{ID: i + 1}
	}

	// create a worker pool
	wp := WorkerPool{
		Tasks:       tasks,
		concurrency: 5, // number of workers that can run at a time
	}

	// run the pool
	wp.Run()
	fmt.Println("All tasks have been processed!")
}
