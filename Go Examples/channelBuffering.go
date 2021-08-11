package main

import "fmt"

func main() {

	channel1 := make(chan int, 3)

	channel1 <- 1
	channel1 <- 2

	num1 := <-channel1
	num2 := <-channel1
	fmt.Println(num1, num2)
}
