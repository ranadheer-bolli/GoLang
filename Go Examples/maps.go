package main

import (
	"fmt"
)

func main() {

	m := make(map[string]int)

	m["1"] = 100
	m["2"] = 200
	m["3"] = 300

	fmt.Println("map is", m)
	fmt.Println("length of map is", len(m))

	firstvalue := m["1"]

	fmt.Println("Firstvalue: ", firstvalue)

	delete(m, "2")
	fmt.Println("map after delete", m)
	fmt.Println("length of map is", len(m))

	value, check := m["2"]

	fmt.Println("m[2]: ", value, check)

	n := map[string]int{"1": 10, "2": 20}
	fmt.Println(n)

}
