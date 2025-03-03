package main 

import "fmt"

func main() {
	tasks := []Task{
		&EmailTask{Email: "email@codeheim.io", Subject: "test", MessageBody: "test"},
		&ImageProcessingTask{ImageUrl: "/image/sample1.jpg"},
		&EmailTask{Email: "emai2@codeheim.io", Subject: "test", MessageBody: "test"},
		&ImageProcessingTask{ImageUrl: "/image/sample2.jpg"},
		&EmailTask{Email: "emai3@codeheim.io", Subject: "test", MessageBody: "test"},
		&ImageProcessingTask{ImageUrl: "/image/sample3.jpg"},
		&EmailTask{Email: "emai4@codeheim.io", Subject: "test", MessageBody: "test"},
		&ImageProcessingTask{ImageUrl: "/image/sample4.jpg"},
	}

	wp := WorkerPool {
		Tasks: tasks,
		concurrency: 5,
	}

	wp.Run()
	fmt.Println("All tasks have completed!")
}
