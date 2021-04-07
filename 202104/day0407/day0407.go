package main

import "fmt"

func main() {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 4, 3))
}

//474. 一和零
func findMaxForm(strs []string, m int, n int) int {
	//dp[i][j][k] 0..i个字符串 需要 j个0，k 个 1

	//当前字符串可以选择，可以不选择
	//dp[i][j][k] = dp[i - 1][j][k] //不选择当前字符串
	//dp[i][j][k] = dp[i-1][j-zeroCount][k-oneCount]
	//选择其中比较大的。

	//dp[0][0][0] = 0;
	dp := make([][][]int, len(strs))
	for i := 0; i < len(strs); i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j < m+1; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}

	countNum := func(num int, str string) int {
		zero, one := 0, 0
		for _, v := range str {
			if v == '1' {
				one++
			} else {
				zero++
			}
		}
		if num == 1 {
			return one
		}
		return zero
	}

	for i := 0; i < len(strs); i++ {
		for j := 0; j < m+1; j++ {
			for k := 0; k < n+1; k++ {

				if i == 0 {
					//第一个字符串能满足最多m个0，n个1的情况下，最大返回也就是1，不满足就是0
					if countNum(0, strs[0]) <= j && countNum(1, strs[0]) <= k {
						dp[0][j][k] = 1
					}
				} else if i > 0 {
					//可以选择当前字符串,则它的前面需要的0，1要满足条件
					if j-countNum(0, strs[i]) >= 0 && k-countNum(1, strs[i]) >= 0 {
						if 1+dp[i-1][j-countNum(0, strs[i])][k-countNum(1, strs[i])] > dp[i-1][j][k] {
							dp[i][j][k] = 1 + dp[i-1][j-countNum(0, strs[i])][k-countNum(1, strs[i])]
						} else {
							dp[i][j][k] = dp[i-1][j][k]
						}
					} else {
						//不选择当前字符串
						dp[i][j][k] = dp[i-1][j][k]
					}
				}
			}
		}
	}

	return dp[len(strs)-1][m][n]
}

//258. 各位相加
func addDigits(num int) int {
	if num < 10 {
		return num
	}

	count := 0
	for ; num != 0; num /= 10 {
		count += (num % 10)
	}

	return addDigits(count)
}

//1143. 最长公共子序列
func longestCommonSubsequence(text1 string, text2 string) int {
	//先写个递归的方式，
	dp := make([][]int, len(text1))
	for i := 0; i < len(text1); i++ {
		dp[i] = make([]int, len(text2))
	}
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {

			if text1[i] == text2[j] {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}
			} else {
				if i > 0 && j > 0 {
					//谁大用谁
					if dp[i-1][j] > dp[i][j-1] {
						dp[i][j] = dp[i-1][j]
					} else {
						dp[i][j] = dp[i][j-1]
					}
				} else if i > 0 {
					//即j == 0
					dp[i][0] = dp[i-1][0]
				} else if j > 0 {
					//即 i == 0
					dp[0][j] = dp[0][j-1]
				}
			}

		}
	}
	return dp[len(text1)-1][len(text2)-1]
}

func longestCommonSubsequence2(text1 string, text2 string) int {
	//先写个递归的方式，
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}
	res := 0

	for j := 0; j < len(text1); j++ {
		//寻找第一个字母的情况
		first := text1[j]

		for i := 0; i < len(text2); i++ {
			if text2[i] == first {
				//找到匹配的了
				res1 := 1 + longestCommonSubsequence(text1[j+1:], text2[i+1:])
				if res1 > res {
					res = res1
				}
				//只要找到了就跳出
				break
			}
		}

		//跳过此字母进行匹配的情况
		res2 := longestCommonSubsequence(text1[j+1:], text2)
		if res < longestCommonSubsequence(text1[j+1:], text2) {
			res = res2
		}

	}
	return res
}
