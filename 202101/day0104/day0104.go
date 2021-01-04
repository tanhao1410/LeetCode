package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(longestWord([]string{"pg", "ptgt", "tgpppttg", "tptttgg", "pgttggtpt", "t", "ptg", "ppgp", "g", "ptgpptpgg"}))
}

//面试题 17.15. 最长单词
func longestWord(words []string) string {
	m := make(map[string]bool)
	for _, v := range words {
		m[v] = true
	}
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) == len(words[j]) {
			for i := 0; i < len(words[i]); i++ {
				if words[i] != words[j] {
					//字母序小的在前
					return words[i] < words[j]
				}
			}
		}
		//大的在前
		return len(words[i]) > len(words[j])
	})

	var canCompose func(s string, canSingle bool) bool
	canCompose = func(s string, canSingle bool) bool {
		if canSingle && m[s] {
			return true
		}
		for i := 1; i < len(s); i++ {
			pre, tail := s[:i], s[i:]
			//前缀是单词
			if m[pre] && canCompose(tail, true) {
				return true
			}
		}

		return false
	}

	for i := 0; i < len(words); i++ {
		word := words[i]
		//截取单词，看是否存在于字典中
		//可以由多个单词组合！
		if canCompose(word, false) {
			return word
		}
	}
	return ""
}

//面试题 01.04. 回文排列
func canPermutePalindrome(s string) bool {
	//用一个hash来判断，如果，所有的字母都是偶数个或只有一个一个字母是奇数个，那么返回true
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			m[s[i]]++
		} else {
			m[s[i]] = 1
		}
	}
	n := 0
	for _, count := range m {
		if count%2 != 0 {
			n++
		}
	}
	return n <= 1
}

//面试题 01.03. URL化
func replaceSpaces(S string, length int) string {
	//两种方式，第一种，重新生成一个字符串，比较简单。第二种，如何在原来的字符串上生成呢。
	//将空格替换为%20，后面的空格怎么办？截取？s
	//直接替换不行。不知道有几个空格，截取的数量就不对。
	s := strings.ReplaceAll(S, " ", "%20")
	i, j := 0, 0
	for ; i < length; i++ {
		if s[j] == '%' {
			j += 3
		} else {
			j++
		}
	}
	return s[:j]
}

//372. 超级次方
func superPow(a int, b []int) int {
	dp := make([]int, 2000)
	dp[0] = a % 1337

	//求个位次方的
	dp2 := make([]int, 10)
	dp2[0] = a % 1337
	for i := 1; i < 10; i++ {
		dp2[i] = (dp2[0] * dp2[i-1]) % 1337
	}

	//a ^ 10,100,1000,....10000
	for i := 1; i < 2000; i++ {
		//10个相乘然后取模
		dp[i] = (((((dp[i-1] * dp[i-1]) % 1337) * ((dp[i-1] * dp[i-1]) % 1337) % 1337) * (((dp[i-1] * dp[i-1]) % 1337) * ((dp[i-1] * dp[i-1]) % 1337) % 1337)) % 1337 * (dp[i-1] * dp[i-1]) % 1337) % 1337
	}

	res := 1

	//先把b代表的数纬度降下来
	//dp[1]代表的是10次方
	for i := 0; i < len(b)-1; i++ {
		if b[i] == 0 {
			continue
		}
		one := (dp[len(b)-i-1] * res) % 1337
		for j := 1; j < b[i]; j++ {
			one = (one * one) % 1337
		}
		res = one
	}

	//最后一位处理
	lastNum := b[len(b)-1]
	if lastNum != 0 {
		return (dp2[lastNum-1] * res) % 1337
	}

	return res
}

//每日一题：509. 斐波那契数
func fib(n int) int {
	dp := make([]int, n+1)
	if n == 0 {
		return 0
	}
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
