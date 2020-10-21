package main

func main() {

}

//189.旋转数组
func rotate(nums []int, k int)  {
	if len(nums) < 2{
		return
	}
	k = k % len(nums)
	for ;k > 0;k--{
		temp := nums[len(nums) -1]
		for i:=len(nums) -1;i > 0;i--{
			nums[i] = nums[i-1]
		}
		nums[0] = temp
	}
}

//91.解码方法
func numDecodings(s string) int {
	//用动态规划算法了
	//以0开头的肯定是返回结果为0
	//从短字符串逐渐上升到长字符串即可解决时间超时问题
	dp := make([]int, len(s))
	//dp[i]代表的是以i为开头的所有的解码方法
	if s[len(s)-1] == '0' {
		dp[len(s)-1] = 0
	} else {
		dp[len(s)-1] = 1
	}

	for i := len(s) - 2; i >= 0; i-- {
		if s[i] > '2' || (s[i] == '2' && s[i+1] > '6' ){
			dp[i] = dp[i+1]
		} else if s[i] == '0' {
			dp[i] = 0
		} else if s[i+1] == '0' {
			if i+2 < len(s) {
				dp[i] = dp[i+2]
			} else {
				dp[i] = 1
			}
		} else {
			if i+2 < len(s) {
				dp[i] = dp[i+1] + dp[i+2]
			}else{
				dp[i] = 2
			}
		}
	}
	return dp[0]
}

//每日一题：925.长按键入
func isLongPressedName(name string, typed string) bool {
	//双指针法解决
	i, j := 0, 0
	for ; i < len(name) && j < len(typed); {

		if name[i] != typed[j] {
			return false
		}
		//与后一个字母不相等
		if i+1 < len(name) {
			if name[i] != name[i+1] {
				for ; j < len(typed) && typed[j] == name[i]; j++ {
				}
			} else {
				j++
			}
			i++
		} else {
			//是最后一个子目了
			for ; j < len(typed) && typed[j] == name[i]; j++ {
			}
			return j == len(typed)
		}
	}

	return i == len(typed) && j == len(typed)
}
