package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}))
	ex := []int{1, 2, 2, 2, 10}
	solution := Constructor(ex)
	res := make([]int, len(ex))
	for i := 0; i < 1000; i++ {
		res[solution.PickIndex()]++
	}
	fmt.Println(res)
}

//528. 按权重随机选择
type Solution struct {
	sum int
	w   []int
}

func Constructor(w []int) Solution {
	sum := 0
	for i := 0; i < len(w); i++ {
		sum += w[i]
		w[i] = sum
	}

	return Solution{
		sum: sum,
		w:   w,
	}
}

func (this *Solution) PickIndex() int {
	res := 0
	randNum := rand.Intn(this.sum) + 1
	//需要知道这个数在哪一个区间
	//通过二分法查找该数在哪个区间
	start, end := 0, len(this.w)-1
	middle := (start + end) / 2
	for ; start <= end; middle = (start + end) / 2 {
		if this.w[middle] == randNum {
			return middle
		}

		if this.w[middle] < randNum {
			if this.w[middle+1] >= randNum {
				return middle + 1
			}
			start = middle + 1
		} else {

			if middle == 0 {
				return 0
			}

			if this.w[middle-1] > randNum {
				end = middle - 1
			} else if this.w[middle-1] == randNum {
				return middle - 1
			} else {
				return middle
			}
		}

	}
	return res
}

//每日一题：684. 冗余连接
func findRedundantConnection(edges [][]int) []int {
	//先形成一个易于遍历的图 map[int] map[int]bool
	//然后进行广度优先遍历，如果能够走到原点，说明在此路有环
	m := make(map[int]map[int]bool)
	for i := 1; i <= len(edges); i++ {
		m[i] = make(map[int]bool)
	}

	for _, edge := range edges {
		start := edge[0]
		end := edge[1]
		m[start][end] = true
		m[end][start] = true
	}

	deletes := make(map[int]bool)
	getLen := func(start int) int {
		res := 0
		for k, _ := range m[start] {
			if !deletes[k] {
				res++
			}
		}
		return res
	}

	//去掉没有后续的，因为它肯定不会导致有环。
	for flag := true; flag; {
		flag = false
		for k, _ := range m {
			if !deletes[k] && getLen(k) == 1 {
				deletes[k] = true
				flag = true
			}
		}
	}
	//最后剩下的就是在换里面的了。

	for i := len(edges) - 1; ; i-- {
		if !deletes[edges[i][0]] && !deletes[edges[i][1]] {
			return edges[i]
		}
	}

}
