package main

import "fmt"

func main() {
	fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 5, 5, 10, 20, 20, 20}))
}

//每日一题：860. 柠檬水找零
func lemonadeChange(bills []int) bool {
	//用一个map记录5,10,20的个数。如果收到了5，直接m[5]+1，如果收到10,m[10=+1,5--,如果收到20，则10--。5++
	five, ten := 0, 0
	for _, v := range bills {
		switch v {
		case 5:
			five++
		case 10:
			ten++
			five--
			if five < 0 {
				return false
			}
		case 20:
			if ten > 0 {
				ten--
				five--
			} else {
				five -= 3
			}

			if five < 0 || ten < 0 {
				return false
			}
		}
	}

	return true
}
