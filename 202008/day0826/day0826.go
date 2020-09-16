package main

import (
	"fmt"
	"strings"
)

//给定一个按非递减顺序排序的整数数组 A，返回每个数字的平方组成的新数组，要求也按非递减顺序排序
func sortedSquares(a []int) []int {
	//主要是绝对值的比较了
	//思路：采用双指针，一个从前走，一个从后走，谁的绝对值大，谁放在新切片的后面，空间为o(n)
	start, end := 0, len(a)-1
	res := make([]int, end+1)
	for i := end; start <= end; i-- {
		if a[end] > 0 {
			if a[start]+a[end] < 0 {
				res[i] = a[start] * a[start]
				start++
			} else {
				res[i] = a[end] * a[end]
				end--
			}
		} else {
			res[i] = a[start] * a[start]
			start++
		}
	}

	//思路2：先找到第一个不为负的数，

	return res
}

func main() {
	repeatedSubstringPattern("abaababaab")
}

//是否由重复子串构成 "abcabcabcabc"
func repeatedSubstringPattern(s string) bool {
	//思路：应该从第一个字母开始
	//如果是重复子串构成，长度一定是它的约数。/2,/3,/4
	length := len(s)
	for n := 1; n <= length/2; n++ {
		if length%n == 0 {
			continue
		}
		//n代表了前n个子串能构成该字符串
		//最简单的方式，用已有的api，替换掉所有，看长度是否为0
		all := strings.ReplaceAll(s, s[0:n], "")
		fmt.Println(s[0:n])
		if len(all) == 0 {
			return true
		}
	}
	return false
}

func repeatedSubstringPattern2(s string) bool {
	//思路：应该从第一个字母开始
	//如果是重复子串构成，长度一定是它的约数。/2,/3,/4
	length := len(s)
	for n := 1; n <= length/2; n++ {
		if length%n == 0 {
			continue
		}
		//n代表了前n个子串能构成该字符串
		//s[0:n]
	}
	return false
}

func fib(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
