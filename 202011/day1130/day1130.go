package main

import "fmt"

func main() {
	fmt.Println(reorganizeString("ababa"))
}

//每日一题：767. 重构字符串
func reorganizeString(S string) string {
	res := ""
	//记录所有字符串的个，
	m := make([]int, 26)
	for _, v := range S {
		m[v-'a'] += 1
		//如果某一个大于一半的话，直接返回""
		if m[v-'a'] > (len(S)+1)/2 {
			return res
		}
	}
	//拼接结果返回
	nextLetter := func(cur int) int {
		res := -1
		max := 0
		for k, v := range m {
			if v > max && k != cur {
				max = v
				res = k
			}
		}
		if res != -1 {
			m[res] -= 1
		}
		return res
	}
	//从里面找最多的哪个作为下一个字母
	next := nextLetter(-1)

	for next != -1 {
		res += string(next + 'a')
		next = nextLetter(next)
	}

	return res
}
