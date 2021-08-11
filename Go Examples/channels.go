package main

import "fmt"

func main() {

	messages := make(chan string)

	go func() {
		messages <- "heyaaa"
	}()

	text := <-messages
	fmt.Println(text)

}
