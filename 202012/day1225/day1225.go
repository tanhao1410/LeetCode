package main

import "sort"

func main() {

}

//每日一题：455. 分发饼干
func findContentChildren(g []int, s []int) int {
	if len(s) == 0 {
		return 0
	}
	res := 0
	//先排序孩子
	sort.Ints(g)
	sort.Ints(s)
	i, j := len(g)-1, len(s)-1
	for i >= 0 && j >= 0 {
		//先从需要最多的饼干的孩子算起，先从最大的饼干算起，若饼干能满足i --，j--，则分配。不能满足的话，i--
		if s[j] >= g[i] {
			i--
			j--
			res++
		} else {
			i--
		}
	}
	return res
}
