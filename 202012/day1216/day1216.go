package main

import (
	"strings"
)

func main() {

}

//每日一题：290. 单词规律
func wordPattern(pattern string, s string) bool {
	m := make(map[byte]string)
	ss := strings.Split(s, " ")
	if len(s) == 0 || len(pattern) != len(ss) {
		return false
	}
	for i := 0; i < len(ss); i++ {
		if pre, ok := m[pattern[i]]; ok {
			if pre != ss[i] {
				return false
			}
		} else {
			m[pattern[i]] = ss[i]
		}
	}
	m2 := make(map[string]bool)
	for _, v := range m {
		m2[v] = true
	}

	return len(m2) == len(m)
}
