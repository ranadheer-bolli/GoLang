package main

import "fmt"

func calculate(a, b int) (int, int, int) {
	res1 := a + b
	res2 := a - b
	res3 := a * b
	return res1, res2, res3

}

func main() {
	ans1, ans2, ans3 := calculate(30, 10)
	fmt.Println(ans1, ans2, ans3)

	_, b, _ := calculate(50, 20)
	fmt.Println(b)
}
