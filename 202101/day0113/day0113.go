package main

import "fmt"

func main() {
	fmt.Println(findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}))
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
