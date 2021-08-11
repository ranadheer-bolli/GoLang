package main

import "fmt"

func addtwo(a int, b int) int {
	return a + b
}

func addthree(a, b, c int) int {
	return a + b + c
}

func main() {
	res1 := addtwo(10, 20)
	res2 := addthree(10, 30, 60)
	fmt.Println(res1, res2)
}
