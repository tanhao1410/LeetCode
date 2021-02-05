package main

import (
	"fmt"
)

func main() {
	fmt.Println(equalSubstring("krpgjbjjznpzdfy", "nxargkbydxmsgby", 14))
	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
}

//581. 最短无序连续子数组
func findUnsortedSubarray(nums []int) int {
	//它前面最大的数，包括自身
	max := make([]int, len(nums))
	//它后面最小的数
	min := make([]int, len(nums))

	//子数组的开头应为：它前面的最大的数 要比后面最小的数 要 小
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			max[i] = nums[i]
		} else {
			if max[i-1] < nums[i] {
				max[i] = nums[i]
			} else {
				max[i] = max[i-1]
			}
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			min[i] = nums[i]
		} else {
			if min[i+1] < nums[i] {
				min[i] = min[i+1]
			} else {
				min[i] = nums[i]
			}
		}
	}

	//找连续子数组的开头，即 不满足后面最小的要比前面最大还要大
	start := 0
	for ; start < len(nums) && max[start] <= min[start]; start++ {
	}

	//找连续子数组的结尾 即 不满足 前面最大的比后面最小的还要大
	end := len(nums) - 1
	for ; end >= 0 && max[end] <= min[end]; end-- {
	}

	if start > end {
		return 0
	}

	return end - start + 1
}

//1539. 第 k 个缺失的正整数
func findKthPositive(arr []int, k int) int {
	seq := 0
	for i, j := 1, 0; ; {
		if j < len(arr) {
			if arr[j] == i {
				i++
				j++
			} else {
				seq++
				if seq == k {
					return i
				}
				i++
			}
		} else {
			seq++
			if seq == k {
				return i
			}
			i++
		}
	}
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
