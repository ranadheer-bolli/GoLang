package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 21, 5}
	sort.Ints(arr)
	fmt.Println(arr)

	s := []string{"c", "bb", "ba", "a"}
	sort.Strings(s)
	fmt.Println(s)

	check := sort.IntsAreSorted(arr)
	fmt.Println("sorted: ", check)

}
