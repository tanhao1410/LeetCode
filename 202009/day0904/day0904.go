package main

import (
	"fmt"
	"math"
)

func main() {

	nums := []int{2, 6, 2, 2, 7}
	fmt.Println(maxSumDivThree2(nums))
}

func maxSumDivThree(nums []int) int {
	dp := []int{0, 0, 0}
	for _, v := range nums {
		mod := v % 3
		switch mod {
		case 2:
			flag := false
			if (dp[1]+v)%3 == 0 && dp[0] < dp[1]+v {
				dp[0] = dp[1] + v
				flag = true
			}
			if (dp[2]+v)%3 == 1 && dp[1] < dp[2]+v {
				dp[1] = dp[2] + v
			}
			if (dp[0]+v)%3 == 2 && dp[2] < dp[0]+v {
				dp[2] = dp[0] + v
				if flag {
				}
			}
		case 1:
			if (dp[2]+v)%3 == 0 && dp[0] < dp[2]+v {
				dp[0] = dp[2] + v
			}
			if (dp[0]+v)%3 == 1 && dp[1] < dp[0]+v {
				dp[1] = dp[0] + v
			}
			if (dp[1]+v)%3 == 2 && dp[2] < dp[1]+v {
				dp[2] = dp[1] + v
			}
		case 0:
			dp[0] = dp[0] + v
			if (dp[1]+v)%3 == 1 {
				dp[1] = dp[1] + v
			}
			if (dp[2]+v)%3 == 2 {
				dp[2] = dp[2] + v
			}
		}
	}
	return dp[0]
}

//5,1,8 ->1 8 ->18
//被三整除的元素最大和 和被三整除
func maxSumDivThree2(nums []int) int {
	//思路：如果一个数本来就能被三整除，那么，肯定加上
	//那么剩余不加的数：。

	//从mod1和mod2中找出两个最小的数
	minMod1, min2Mod1 := math.MaxInt32, math.MaxInt32
	minMod2, min2Mod2 := math.MaxInt32, math.MaxInt32

	res := 0
	mod1Count, mod2Count := 0, 0
	for i := 0; i < len(nums); i++ {
		switch nums[i] % 3 {
		case 0:
			res += nums[i]
		case 1:
			if nums[i] <= minMod1 {
				minMod1, min2Mod1 = nums[i], minMod1
			} else if nums[i] < min2Mod1 {
				min2Mod1 = nums[i]
			}
			mod1Count += nums[i]
		case 2:
			if nums[i] <= minMod2 {
				minMod2, min2Mod2 = nums[i], minMod2
			} else if nums[i] < min2Mod2 {
				min2Mod2 = nums[i]
			}
			mod2Count += nums[i]
		}
	}

	//四种情况，从mod1中去除一个，两个
	one := res + mod1Count + mod2Count - minMod1
	two := res + mod1Count + mod2Count - minMod1 - min2Mod1
	three := res + mod1Count + mod2Count - minMod2
	four := res + mod1Count + mod2Count - minMod2 - min2Mod2

	if one%3 == 0 {
		res = one
	}
	if two%3 == 0 && two > res {
		res = two
	}
	if three%3 == 0 && three > res {
		res = three
	}
	if four%3 == 0 && four > res {
		res = four
	}

	//还有个如果所有加一块%3==0
	if (res+mod1Count+mod2Count)%3 == 0 {
		res = res + mod1Count + mod2Count
	}

	return res
}

func Fibs(n int) {
	if n == 1 {

	}
}
