package main

import "fmt"

func main() {
	fmt.Println(lastStoneWeight([]int{2, 2}))
}

//每日一题：1046. 最后一块石头的重量
func lastStoneWeight(stones []int) int {

	for m1, m2 := -1, -1; ; stones[m1], stones[m2] = stones[m1]-stones[m2], 0 {
		m1, m2 = -1, -1
		for i := 0; i < len(stones); i++ {
			if m1 == -1 && stones[i] > 0 {
				m1 = i
			} else if m2 == -1 && stones[i] > 0 {
				if stones[i] > stones[m1] {
					m2 = m1
					m1 = i
				} else {
					m2 = i
				}
			} else if m1 != -1 && m2 != -1 && stones[i] > stones[m1] {
				m1, m2 = i, m1
			} else if m1 != -1 && m2 != -1 && stones[i] > stones[m2] {
				m2 = i
			}
		}

		if m2 < 0 || stones[m2] == 0 {
			if m1 == -1 {
				return 0
			} else {
				return stones[m1]
			}
		}
	}
}
