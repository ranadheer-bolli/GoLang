package main

import (
	"fmt"
)

func main() {

	messages := make(chan string)
	signals := make(chan string)

	select {
	case msg := <-messages:
		fmt.Println("received messages", msg)
	default:
		fmt.Println("messages not received")
	}

	msg := "hello"
	select {
	case messages <- msg:
		fmt.Println("message sent", msg)
	default:
		fmt.Println("not sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received messages", msg)
	case signal := <-signals:
		fmt.Println("recieved signals", signal)
	default:
		fmt.Println("no activity")
	}

}
