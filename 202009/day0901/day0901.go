package main

import "fmt"

func main() {

	//nums := []int{111,11}
	//fmt.Println(longestSubarray2(nums,10))

	//fmt.Println(isPalindrome(121))
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))

}

//最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	//说明至少有两个字符串了，找两个字符串之间的公共前缀
	prefix := GetTwoStrCommonPrefix(strs[0], strs[1])

	for j := 2; j < len(strs); j++ {
		prefix = GetTwoStrCommonPrefix(prefix, strs[j])
	}

	return prefix
}

func GetTwoStrCommonPrefix(s1, s2 string) string {
	i := 0
	for ; i < len(s1) && i < len(s2) && s1[i] == s2[i]; i++ {
	}
	return s1[:i]
}

//判断是否是回文数
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	res := 0
	for temp := x; temp > 0; temp = temp / 10 {
		res = res*10 + temp%10
	}
	return res == x
}

func longestSubarray2(nums []int, limit int) int {

	//dp[i] 值代表的是以nums[i]为结尾的最大的符合条件的子串长度
	dp := make([]int, len(nums))
	res := 1
	for i := 1; i < len(nums); i++ {

		if nums[i] == nums[i-1] {
			dp[i] = dp[i-1] + 1 //i可以直接加入原子序列，比原来多一个
			continue
		}
		j, flag := 1, false

		for ; j < dp[i-1]+1; j++ {

			if !IsInLimit(nums[i], nums[i-j], limit) {
				flag = true //有不符合的了，需要记住这个不符合的位置
				break
			}
		}

		if flag {
			dp[i] = j //从前面j开始不符合的
		} else {
			dp[i] = dp[i-1] + 1 //i可以直接加入原子序列，比原来多一个
		}

		if dp[i] > res {
			res = dp[i] //返回最大的那个
		}
	}
	return res
}

//绝对值差 不超过限制的最长连续 子数组
func longestSubarray(nums []int, limit int) int {
	//新思路：动态规划的思想呢，dp[i][j] 从i->j截取 。若
	dp := [][]int{}
	//初始化dp
	for i := 0; i < len(nums); i++ {
		row := make([]int, len(nums))
		dp = append(dp, row)
	}
	//就是找到i-j的最大值
	res := 1
	dp[0][0] = 1
	for i := 0; i < len(nums); i++ {
		dp[i][i] = 1
		for j := i + 1; j < len(nums); j++ {
			if dp[i][j-1] == 0 {
				break
			} else {
				//看nums[j]加入其中符不符合要求
				flag := false
				for k := i; k < j; k++ {
					if !IsInLimit(nums[k], nums[j], limit) {
						flag = true
					}
				}
				if !flag {
					dp[i][j] = dp[i][j-1]
					if j-i+1 > res {
						res = j - i + 1
					}
				}
			}
		}
	}
	return res
}

func IsInLimit(n1, n2, limit int) bool {
	if n1 > n2 {
		return n1-n2 <= limit
	}
	return n2-n1 <= limit
}
