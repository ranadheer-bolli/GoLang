package main

import "fmt"

func swap(a, b int) {
	temp := a
	a = b
	b = temp
}

func swapByPointer(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	a := 10
	b := 20
	fmt.Println(a, b)
	swap(a, b)
	fmt.Println(a, b)
	swapByPointer(&a, &b)
	fmt.Println(a, b)
	fmt.Print(&a, &b)

}
