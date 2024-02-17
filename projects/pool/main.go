package main

import (
	"fmt"
)

func main() {
	// create new tasks
	tasks := []Task{
		&EmailTask{Email: "WeiMing.lee@yahoo.com", Subject: "Just test", MessageBody: "Thanks Wei Ming"},
		&ImageProcessingTask{ImageUrl: "www.yahoo.com"},
		&EmailTask{Email: "WeiKeat.lee@hotmail.com", Subject: "Just test", MessageBody: "Thank you"},
		&ImageProcessingTask{ImageUrl: "www.hotmail.com"},
		&EmailTask{Email: "WeiLit.lee@msn.com", Subject: "Just test", MessageBody: "Sorry Wei Lit"},
		&ImageProcessingTask{ImageUrl: "www.msn.com"},
		&EmailTask{Email: "Angie.Ng@yahoo.com", Subject: "Just test", MessageBody: "Love you!"},
		&ImageProcessingTask{ImageUrl: "www.yahoo.com"},
		&EmailTask{Email: "Naomi.Goh@facebook.com", Subject: "Just test", MessageBody: "How are you?"},
		&ImageProcessingTask{ImageUrl: "www.facebook.com"},
		&EmailTask{Email: "Jun.lee@yahoo.com", Subject: "Just test", MessageBody: "You are so handsome"},
		&ImageProcessingTask{ImageUrl: "www.yamaha.com"},
		&EmailTask{Email: "Na.lee@yes.com.my", Subject: "Just test", MessageBody: "Naugthy You"},
		&ImageProcessingTask{ImageUrl: "www.hollywood.com"},
		&EmailTask{Email: "Edward.Terng@bi.com", Subject: "Just test", MessageBody: "Just run"},
		&ImageProcessingTask{ImageUrl: "www.bi.com.my"},
		&EmailTask{Email: "Sally.Wong@butterinc.com", Subject: "Just test", MessageBody: "Good Bye"},
		&ImageProcessingTask{ImageUrl: "www.monster.inc.com"},
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
