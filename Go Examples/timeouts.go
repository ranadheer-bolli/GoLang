package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "Good"
	}()

	select {
	case msg := <-c1:
		fmt.Println(msg)
	case <-time.After(time.Second):
		fmt.Println("timeout 2")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "Morning"
	}()

	select {
	case msg := <-c2:
		fmt.Println(msg)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
