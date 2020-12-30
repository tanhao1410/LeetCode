package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(printBin(0.1))
	fmt.Println(lastStoneWeight([]int{2, 2}))
}

//面试题 05.02. 二进制数转字符串
func printBin(num float64) string {

	//得到2^n 分之一
	get2N := func(n int) float64 {
		res := 1.0
		for i := 0; i < n; i++ {
			res /= 2.0
		}
		return res
	}

	oneIndex := make(map[int]bool)
	for i := 1; i < 18; i++ {
		base := get2N(i)
		if num >= base {
			oneIndex[i] = true
			num -= base
		}
	}

	if num != 0 {
		return "ERROR"
	}

	res := "0."

	for j := 1; j <= 17; j++ {
		if oneIndex[j] {
			res += "1"
		} else {
			res += "0"
		}
	}
	//再根据该数反过来计算，如果相等的话，说明是的，如果不相等，说明数不行。
	return strings.TrimRight(res, "0")
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
