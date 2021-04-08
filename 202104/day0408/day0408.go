package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

func main() {
	//
	for i := 0; i < 200; i++ {
		fmt.Print(rand.Intn(1000), ",")
	}
}

//494. 目标和
func findTargetSumWays(nums []int, S int) int {
	//暴力方法，2^20  递归形式
	res := 0

	var next func(index, preSum int)
	next = func(index, preSum int) {
		//最后一个
		if index == len(nums)-1 {
			if preSum+nums[index] == S {
				res += 1
			}
			//如果是0的话，+-都可以
			if preSum-nums[index] == S {
				res += 1
			}
			return
		}
		//+ -
		next(index+1, preSum+nums[index])
		next(index+1, preSum-nums[index])

	}
	next(0, 0)
	return res
}

//每日一题：153. 寻找旋转排序数组中的最小值
func findMin(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	//数组是递增的
	if nums[len(nums)-1] > nums[0] {
		return nums[0]
	}

	start, end := 0, len(nums)-1
	middle := (start + end) / 2
	//说明前面是递增
	if nums[middle] > nums[end] {
		return findMin(nums[middle+1:])
	} else {
		return findMin(nums[:middle+1])
	}
}

//322. 零钱兑换
func coinChange(coins []int, amount int) int {
	//背包问题： 最少硬币个数
	//dp[i][j] 凑成j金额所需要的最少数量 i 为coins
	// dp[i + 1][j] =
	//1.不用这个金额 dp[i + 1][j] = dp[i][j]
	//2.用这个金额 dp[i+1][j] = dp[i][j-coins[i+1] * n] + n 关键问题，n 应该怎么办？ 找到它的最小值
	//还需要考虑总数
	dp := make([][]int, len(coins))
	for i := 0; i < len(coins); i++ {
		dp[i] = make([]int, amount+1)
	}
	for i := 0; i < len(coins); i++ {
		for j := 0; j < amount+1; j++ {
			//就一种硬币时，直接看是否能整除即可
			if i == 0 {
				if j%coins[i] == 0 {
					dp[i][j] = j / coins[i]
				} else {
					dp[i][j] = -1
				}
			} else {
				//多了一种硬币后，可以选择用新硬币的数量 假设为n
				min := 100000
				for n := 0; n*coins[i] <= j; n++ {
					if dp[i-1][j-coins[i]*n] != -1 { //说明前面可以凑成
						//能凑成，但不一定是最小的。
						if dp[i-1][j-coins[i]*n]+n < min {
							min = dp[i-1][j-coins[i]*n] + n
						}
					}
				}

				if min == 100000 {
					dp[i][j] = -1
				} else {
					dp[i][j] = min
				}
			}
		}
	}
	return dp[len(coins)-1][amount]
}

//18. 四数之和
func fourSum(nums []int, target int) [][]int {
	// 先排序，从小到大
	sort.Ints(nums)
	res := [][]int{}
	//从中选出两个
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {

			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			preSum := nums[i] + nums[j]
			//在接下来的空间中找到 target - preSum

			m := make(map[int]int)
			//记录上一个
			pre := math.MaxInt32
			for k := j + 1; k < len(nums); k++ {
				if v, ok := m[target-preSum-nums[k]]; ok && k > v { //比前面的存在的下标要大才可以
					//不能和上一步用的数字相同
					if pre == math.MaxInt32 || pre != nums[k] {
						//找到一个
						item := []int{nums[i], nums[j], target - preSum - nums[k], nums[k]}
						res = append(res, item)
					}
					pre = nums[k]
				} else {
					m[nums[k]] = k
				}
			}
		}
	}
	return res
}
