package main

import "fmt"

func main() {
	fmt.Println(myPow(2, -1))
}
func myPow(x float64, n int) float64 {
	res := x
	if n == 0 || x == 1 {
		return 1
	}
	if n == -1 && n%2 == 0 {
		return 1
	} else if n == -1 && n%2 == 1 {
		return -1
	}
	if n > 0 {
		for i := 1; i < n; i++ {
			res *= x
			if res == 0 {
				return 0
			}
		}
	} else {

		res2 := 1 / res
		for i := 1; i < -n; i++ {
			res *= x
			res2 = 1 / res
			if res2 == 0 {
				return 0
			}
		}
		res = res2

	}
	return res
}
