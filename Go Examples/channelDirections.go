package main

import (
	"fmt"
	"time"
)

func ping(pings chan<- string, msg string) {
	pings <- msg

}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "hello")
	pong(pings, pongs)
	time.Sleep(time.Millisecond)
	fmt.Println(<-pongs)

}
