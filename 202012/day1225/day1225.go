package main

import "sort"

func main() {

}

//436. 寻找右区间
func findRightInterval(intervals [][]int) []int {
	res := make([]int, len(intervals))
	//记录以某个起点开头的区间下标
	m := make(map[int]int)
	//记录所有的起点
	a := make([]int, len(intervals))
	for k, v := range intervals {
		m[v[0]] = k
		a[k] = v[0]
	}
	//排序下所有的起点坐标
	sort.Ints(a)

	//看是否存在右区间
	search := func(start int) int {
		//就是判断m中是否有key大于等于start,有，返回最小的那个的value
		//从数组a中查找。二分法查找
		//已经大于最大的了，肯定不存在了，直接返回-1
		if a[len(a)-1] < start {
			return -1
		}
		b, e := 0, len(a)-1
		mi := (b + e) / 2
		for b <= e {
			if a[mi] == start {
				//以该数为起点的区间存在
				return m[a[mi]]
			} else if a[mi] > start {
				e = mi - 1
				mi = (b + e) / 2
			} else {
				b = mi + 1
				mi = (b + e) / 2
			}
		}
		return m[a[b]]
	}

	for k, v := range intervals {
		res[k] = search(v[1])
	}

	return res
}

//每日一题：455. 分发饼干
func findContentChildren(g []int, s []int) int {
	res := 0
	//先排序孩子
	sort.Ints(g)
	sort.Ints(s)
	for i, j := len(g)-1, len(s)-1; i >= 0 && j >= 0; i-- {
		//先从需要最多的饼干的孩子算起，先从最大的饼干算起，若饼干能满足i --，j--，则分配。不能满足的话，i--
		if s[j] >= g[i] {
			j, res = j-1, res+1
		}
	}
	return res
}
