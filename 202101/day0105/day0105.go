package main

import "fmt"

func main() {
	fmt.Println(largeGroupPositions("bbaaabbbc"))
}

//每日一题：830. 较大分组的位置
func largeGroupPositions(s string) [][]int {
	res := [][]int{}
	start, end := 0, 0
	for pre, i := s[0], 1; i < len(s); i++ {
		if s[i] == pre {
			end++
		} else {
			//遇到不相等的了
			if end-start >= 2 {
				res = append(res, []int{start, end})
			}
			start, end, pre = i, i, s[i]
		}
	}
	if end-start >= 2 {
		res = append(res, []int{start, end})
	}
	return res
}
