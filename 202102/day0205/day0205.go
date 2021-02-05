package main

import (
	"fmt"
)

func main() {
	fmt.Println(equalSubstring("krpgjbjjznpzdfy", "nxargkbydxmsgby", 14))

}

//每日一题：1208. 尽可能使字符串相等
func equalSubstring(s string, t string, maxCost int) int {
	//思路：判断两字符串个差值是多少。优先选择小的。最大不超过maxCost

	nums := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] > t[i] {
			nums = append(nums, int(s[i]-t[i]))
		} else {
			nums = append(nums, int(t[i]-s[i]))
		}
	}

	//必须是连续的
	//采用滑动窗口的形式
	window, windowValue := 0, 0
	start := 0
	for i := 0; i < len(nums); i++ {
		if windowValue+nums[i] <= maxCost {
			window++
			windowValue += nums[i]
		} else {
			//窗口移动直到windowValue<= maxConst
			for windowValue, start = windowValue+nums[i]-nums[start], start+1; windowValue > maxCost; {
				i++
				if i >= len(nums) || start >= len(nums) {
					return window
				}
				windowValue += nums[i]
				windowValue -= nums[start]
				start++
			}
		}
	}
	return window
}
