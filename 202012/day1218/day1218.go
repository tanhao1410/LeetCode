package main

import "math/rand"

func main() {

}

//398. 随机数索引
type Solution struct {
	m map[int][]int
}

func Constructor(nums []int) Solution {
	m := make(map[int][]int)
	for k, v := range nums {
		if _, ok := m[v]; ok {
			m[v] = append(m[v], k)
		} else {
			m[v] = []int{k}
		}
	}
	return Solution{
		m: m,
	}
}

func (this *Solution) Pick(target int) int {
	ints := this.m[target]
	index := rand.Intn(len(ints))
	return this.m[target][index]
}

//每日一题：389. 找不同
func findTheDifference(s string, t string) byte {
	m := make([]int, 26)
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']--
		m[t[i]-'a']++
	}
	m[t[len(s)]-'a']++
	for i := 0; ; i++ {
		if m[i] > 0 {
			return byte(i + 'a')
		}
	}
}
