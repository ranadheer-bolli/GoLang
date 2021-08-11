package main

import (
	"errors"
	"fmt"
)

type argerror struct {
	arg  int
	prob string
}

func f1(a int) (int, error) {
	if a == 42 {
		return -1, errors.New("can't work with 42")
	}
	return a + 3, nil

}

func (e *argerror) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(a int) (int, error) {
	if a == 42 {
		return -1, &argerror{a, "can't work with it"}
	}
	return a + 3, nil

}

func main() {

	a := []int{7, 42}

	for _, v := range a {
		if res, e := f1(v); e != nil {
			fmt.Println("f1 failed: ", e)
		} else {
			fmt.Println("f1 worked: ", res)
		}
	}

	for _, v := range a {
		if res, e := f2(v); e != nil {
			fmt.Println("f1 failed: ", e)
		} else {
			fmt.Println("f1 worked: ", res)
		}
	}
}
