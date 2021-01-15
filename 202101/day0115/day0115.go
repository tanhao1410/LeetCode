package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(removeStones([][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}))
}

//每日一题：947. 移除最多的同行或同列石头--方法错误，未解决
func removeStones(stones [][]int) int {
	//用一个数代表一个点 int = 20000 * stones[0]  + stones[1]

	stones2 := make([]int, len(stones))
	for k, v := range stones {
		stones2[k] = v[0]*20000 + v[1]
	}

	//记录最小的哪一个
	min, minNum := math.MaxInt32, 0
	m := make(map[int]map[int]bool)
	for i := 0; i < len(stones2); i++ {
		for j := 0; j < len(stones2); j++ {
			if i != j && (stones2[i]/20000 == stones2[j]/20000 || stones2[i]%20000 == stones2[j]%20000) {
				if _, ok := m[stones2[i]]; ok {
					m[stones2[i]][stones2[j]] = true
				} else {
					m[stones2[i]] = map[int]bool{stones2[j]: true}
				}
			}
		}

		//即该点不与任何点在一行，一列
		if len(m[stones2[i]]) == 0 {
			delete(m, stones2[i])
		} else if len(m[stones2[i]]) < min {
			min = len(m[stones2[i]])
			minNum = stones2[i]
		}
	}

	res := 0
	//删去最小的哪一个
	for min < math.MaxInt32 {

		//删除
		delNum := minNum
		delete(m, delNum)
		res++

		min = math.MaxInt32
		minNum = -1
		//删除里面包含它的所有点
		for k, v := range m {
			if v[delNum] {
				delete(v, delNum)
			}
			if len(v) == 0 {
				delete(m, k)
			} else if len(v) < min {
				min = len(v)
				minNum = k
			}
		}

	}

	return res
}
