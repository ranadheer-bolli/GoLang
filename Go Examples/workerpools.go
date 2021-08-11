package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {

	for ele := range jobs {
		fmt.Println(id, ele)
		fmt.Println("worker ", id, " started job ", ele)
		time.Sleep(time.Second)
		fmt.Println("worker ", id, " finished job ", ele)
		results <- ele * 2
	}

}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	close(jobs)

	for j := 1; j <= 5; j++ {
		<-results
	}

}
