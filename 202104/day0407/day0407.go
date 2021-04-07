package main

func main() {

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
