package main

import (
	"fmt"
	"time"
)

func display(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ": ", i)
	}
}

func main() {

	display("direct")

	go display("routine")

	//fmt.Println("yessss")
	go func(from string) {
		for i := 0; i < 3; i++ {
			fmt.Println(from, ": ", i)
		}

	}("second one")

	time.Sleep(time.Second)
	fmt.Println("done")

}
