package main

import "fmt"

func main() {
	fmt.Println(lastStoneWeight([]int{2, 2}))
}

//面试题 08.07. 无重复字符串的排列组合
func permutation(S string) []string {

	res := []string{}
	if len(S) == 1 {
		res = append(res, S)
		return res
	}

	trimS := func(s string, b byte) string {
		res := make([]byte, len(s)-1)
		for i, j := 0, 0; i < len(s); i++ {
			if s[i] != b {
				res[j] = s[i]
				j++
			}
		}
		return string(res)
	}

	//采用递归的方式
	for i := 0; i < len(S); i++ {
		for _, v := range permutation(trimS(S, S[i])) {
			res = append(res, string(S[i])+v)
		}
	}
	return res
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
