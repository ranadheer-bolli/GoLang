package main

import "fmt"

func main() {

	queue := make(chan string, 3)

	queue <- "1"
	queue <- "2"
	queue <- "3"
	close(queue)

	for ele := range queue {
		fmt.Println(ele)
	}

}
