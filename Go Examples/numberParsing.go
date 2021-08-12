package main

import (
	"fmt"
	"strconv"
)

func main() {

	// this 64 tells how many bits of precision to parse.

	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// it will recognize hexa formatted numbers

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// Atoi is a convenience function for basic base-10 int parsing.
	// string to int conversion
	k, _ := strconv.Atoi("1357")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
