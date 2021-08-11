package main

import (
	"fmt"
)

func main() {

	a := make([]int, 3)
	fmt.Println("empty", a)

	a[0] = 0
	a[1] = 1
	a[2] = 2

	fmt.Println("set", a)

	fmt.Println("get", a[1])

	a = append(a, 3)

	fmt.Println("appended", a[3])
	fmt.Println("a", a)

	b := make([]int, len(a))
	fmt.Println("b:", b)

	copy(b, a)
	fmt.Println("b:copy of a", b)

	l := a[1:3]
	fmt.Println("sl1:", l)

	l = a[2:]
	fmt.Println("sl2:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerlen := i + 1
		twoD[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
