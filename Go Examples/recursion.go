package main

import "fmt"

func fact(a int) int {
	if a <= 2 {
		return a
	}
	return a * fact(a-1)
}

func main() {

	fmt.Println(fact(6))
}
