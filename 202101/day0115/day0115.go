package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(removeStones([][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}))
	fmt.Println(countArrangement(15))
}

//526. 优美的排列
func countArrangement(n int) int {
	res := 0
	nums := make([][]int, n)
	for i := 0; i < n; i++ {
		numsItem := []int{}
		//先加上比它小的
		for j := 1; j <= i+1; j++ {
			if (i+1)%j == 0 {
				numsItem = append(numsItem, j)
			}
		}
		for j := i + 2; j <= n; j++ {
			if j%(i+1) == 0 {
				numsItem = append(numsItem, j)
			}
		}
		nums[i] = numsItem
	}

	var assignNum func(index int, used []int)
	assignNum = func(index int, used []int) {
		if index == n {
			res++
			return
		}
		//遍历可以选择的数
		for _, v := range nums[index] {
			//没被使用
			if used[v-1] == 0 {
				used[v-1] = 1
				assignNum(index+1, used)
				//不使用了
				used[v-1] = 0
			}
		}
	}
	assignNum(0, make([]int, n))
	return res
}

//面试题 16.20. T9键盘
func getValidT9Words(num string, words []string) []string {
	res := []string{}
	//思路：每一个单词都有一个对应的数字
	m := map[byte]string{'a': "2", 'b': "2", 'c': "2", 'd': "3", 'e': "3", 'f': "3", 'g': "4", 'h': "4", 'i': "4", 'j': "5", 'k': "5", 'l': "5", 'm': "6", 'n': "6", 'o': "6", 'p': "7", 'q': "7", 'r': "7", 's': "7", 't': "8", 'v': "8", 'u': "8", 'w': "9", 'x': "9", 'y': "9", 'z': "9"}

	word2num := func(word string) string {
		res := ""
		for j := 0; j < len(word); j++ {
			//单词对应的数字
			res += m[word[j]]
		}
		return res
	}
	for i := 0; i < len(words); i++ {
		if len(words[i]) == len(num) && num == word2num(words[i]) {
			res = append(res, words[i])
		}
	}
	return res
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
