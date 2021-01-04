package main

import "strings"

func main() {

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
