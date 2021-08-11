// Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.

package main

import (
	"fmt"
)

func sum(nums ...int) {
	fmt.Print(nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	fmt.Println(" sum : ", sum)

}

func main() {

	sum(1, 2, 3)
	sum(100, 300, 500)

	nums := []int{1, 2, 3, 4}
	sum(nums...)

}
