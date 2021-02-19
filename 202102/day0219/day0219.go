package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(thirdMax([]int{1, 2, -2147483648}))
}

//414. 第三大的数
func thirdMax(nums []int) int {

	first, second, third := nums[0], math.MinInt32, math.MinInt32

	secondFlag, thirdFlag := false, false
	for i := 1; i < len(nums); i++ {
		if nums[i] == first || (nums[i] == second && secondFlag) || (nums[i] == third && thirdFlag) {
			continue
		}
		if nums[i] > first {
			first, second, third = nums[i], first, second
			if secondFlag {
				thirdFlag = true
			}
			secondFlag = true
		} else if nums[i] > second {
			second, third = nums[i], second
			if secondFlag {
				thirdFlag = true
			}
			secondFlag = true
		} else if nums[i] > third || (!thirdFlag && nums[i] == third) {
			third = nums[i]
			thirdFlag = true
		}

	}
	if thirdFlag {
		return third
	}
	return first
}
