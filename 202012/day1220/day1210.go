package main

import "fmt"

func main() {
	fmt.Print(removeDuplicateLetters("bcabc"))
}

//205. 同构字符串
func isIsomorphic(s string, t string) bool {
	m1 := make(map[uint8]byte)
	m2 := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if v, ok := m1[s[i]]; ok {
			if v != t[i] {
				return false
			}
			if v2, ok := m2[t[i]]; ok {
				if v2 != s[i] {
					return false
				}
			} else {
				return false
			}
		} else {
			m1[s[i]] = t[i]
			if _, ok := m2[t[i]]; ok {
				return false
			} else {
				m2[t[i]] = s[i]
			}
		}
	}

	return true
}

//每日一题：316. 去除重复字母
func removeDuplicateLetters(s string) string {
	res := ""
	//思路：先记录有哪些字母出现了，最后返回的字符串中包含这些字母
	//从小往大试，先试最小的，如果它的后面除了自己，所有的字母都出现过，那么，就选择该字母
	//如果有没出现过的，试下一个最小的。直到找到一个。
	//然后再重新开始，已经选择过的，就不用再判断了
	//用一个数组记录字母第一次出现的位置，如果没有出现，设为-1
	m := make([]int, 26)

	for _, v := range s {
		if m[v-'a'] == 0 {
			m[v-'a'] = 1
		}
	}
	alreadyIndex := -1
	//判断比index 大的位置，除了自己是否其它的字母在之后都出现过
	behindExit := func(i int) int {
		//统计还剩多少个字母
		count1 := 0
		for j := 0; j < 26; j++ {
			if m[j] > 0 && j != i {
				count1++
			}
		}
		m2 := make([]int, 26)
		//找到后面字符中i第一次出现的位置
		al := 0
		for j := alreadyIndex + 1; j < len(s); j++ {
			if s[j]-'a' == uint8(i) {
				al = j
				break
			}
		}
		//统计在i后面，还有多少个字母，不统计已经加入结果集中的
		for j := al + 1; j < len(s); j++ {
			if m2[s[j]-'a'] == 0 && m[s[j]-'a'] > 0 {
				m2[s[j]-'a'] = 1
			}
		}
		for j := 0; j < 26; j++ {
			if m2[j] > 0 && j != i {
				count1--
			}
		}
		if count1 <= 0 {
			alreadyIndex = al
			return 1
		} else {

			return -1
		}
	}

	stop := func() bool {
		for i := 0; i < 26; i++ {
			if m[i] > 0 {
				return false
			}
		}
		return true
	}

	for !stop() {
		//从最小的开始找
		for i := 0; i < 26; i++ {
			if m[i] > 0 {
				//说明该字母出现了，看它后面是否所有的字母都有。m[i]为该字母首次出现的位置
				if behindExit(i) > 0 {
					res += string(i + 'a')
					//即该字母已经选了，它前面的都不能再选了
					m[i] = -1
					break
				}
			}
		}
	}

	return res
}
