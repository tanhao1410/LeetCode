package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		fmt.Print(rand.Intn(1000), ",")
	}
}

//213. 打家劫舍 II
func rob(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}

	//动态规划算法：+ 用一个数组额外记录，开始的地方
	startIndex := make([]int, len(nums))
	dp := make([]int, len(nums))
	dp1 := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {

		if i == 1 {

			if nums[1] > nums[0] {
				dp[1] = nums[1]
				startIndex[1] = 1
			} else {
				dp[1] = nums[0]
			}

			dp1[1] = nums[1]

		} else if i == len(nums)-1 {

			//最后一个特殊处理
			if nums[i]+dp[i-2] > dp[i-1] {
				if startIndex[i-2] == 0 {

					if dp1[i-2]+nums[i] > dp[i-1] {
						//要最后一个数，不要第一个数了
						dp[i] = dp1[i-2] + nums[i]
					} else {
						dp[i] = dp[i-1]
					}

				} else {

					dp[i] = dp[i-2] + nums[i]
				}
			} else {
				dp[i] = dp[i-1]
			}

		} else if i > 1 {

			if nums[i]+dp[i-2] > dp[i-1] {
				dp[i] = nums[i] + dp[i-2]
				startIndex[i] = startIndex[i-2]
			} else {
				dp[i] = dp[i-1]
				startIndex[i] = startIndex[i-1]
			}

			//dp1是不能选第一个数的
			if i == 2 {
				if nums[2] > nums[1] {
					dp1[2] = nums[2]
				} else {
					dp1[2] = dp1[1]
				}
			} else {
				if nums[i]+dp1[i-2] > dp1[i-1] {
					dp1[i] = nums[i] + dp1[i-2]
				} else {
					dp1[i] = dp1[i-1]
				}
			}

		}
	}

	return dp[len(nums)-1]
}
