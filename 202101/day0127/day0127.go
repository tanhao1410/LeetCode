package main

import (
	"sort"
	"strings"
)

func main() {

}

//633. 平方数之和
func judgeSquareSum(c int) bool {
	//思路：记录所有的平方数，判断是否存在两个之和
	m := map[int]bool{}
	for i := 0; i < 1<<16 && i*i <= c; i++ {
		m[i*i] = true
	}

	for k, _ := range m {
		if m[c-k] {
			return true
		}
	}
	return false
}

//692. 前K个高频单词
func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	for _, s := range words {
		if _, ok := m[s]; ok {
			m[s] += 1
		} else {
			m[s] = 1
		}
	}

	type StrNum struct {
		s string
		c int
	}
	l := []*StrNum{}

	for k, v := range m {
		l = append(l, &StrNum{
			s: k,
			c: v,
		})
	}

	sort.Slice(l, func(i, j int) bool {
		//数量相等的话
		if l[i].c == l[j].c {
			return strings.Compare(l[i].s, l[j].s) < 0
		}
		return l[i].c > l[j].c
	})

	res := make([]string, k)
	for i := 0; i < k; i++ {
		res[i] = l[i].s
	}

	return res
}

//676. 实现一个魔法字典
type MagicDictionary struct {
	m map[int][]string
}

func Constructor() MagicDictionary {
	return MagicDictionary{m: map[int][]string{}}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, s := range dictionary {
		sLen := len(s)
		if _, ok := this.m[sLen]; ok {
			this.m[sLen] = append(this.m[sLen], s)
		} else {
			this.m[sLen] = []string{s}
		}
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	//从同长度的里面找
	sLen := len(searchWord)
	if words, ok := this.m[sLen]; ok {
		for _, word := range words {
			diff := 0
			//判断是否仅一个字母不相同
			for i := 0; i < sLen; i++ {
				if word[i] != searchWord[i] {
					diff++
				}
			}
			if diff == 1 {
				return true
			}
		}
	}
	return false
}
