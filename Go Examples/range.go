package main

import "fmt"

func main() {

	nums := []int{1, 2, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum is: ", sum)

	for i, v := range nums {
		if v == 2 {
			fmt.Println("index is: ", i)
		}
	}

	m := map[string]int{"a": 1, "b": 2, "c": 3}

	for key, value := range m {
		fmt.Println("key: ", key, " value: ", value)
	}

	for key := range m {
		fmt.Println("key: ", key)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
