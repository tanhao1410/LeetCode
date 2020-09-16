package main

import (
	"math"
)

func main2() {
	//nums := []int{1,2,3,4}
	//fmt.Println(maximumProduct(nums))
}

//三个数的最大乘积
func maximumProduct(nums []int) int {
	//思路：刚好三个数，直接返回
	if len(nums) == 3 {
		return nums[0] * nums[1] * nums[2]
	}

	//先找最大的三个数
	max1, max2, max3 := math.MinInt32, math.MinInt32, math.MinInt32

	//最小的两个数（负数的情况）
	min1, min2 := math.MaxInt32, math.MaxInt32

	fuCount, zhCount := 0, 0
	for i := 0; i < len(nums); i++ {

		if nums[i] < 0 {
			fuCount++
		} else {
			zhCount++
		}

		//最大的三个数
		if nums[i] >= max1 {
			max1, max2, max3 = nums[i], max1, max2
		} else if nums[i] >= max2 {
			max2, max3 = nums[i], max2
		} else if nums[i] >= max3 {
			max3 = nums[i]
		}

		//最小的两个数
		if nums[i] <= min1 {
			min1, min2 = nums[i], min1
		} else if nums[i] < min2 {
			min2 = nums[i]
		}
	}

	if fuCount < 2 || zhCount == 0 {
		//负数的个数小于2个或没有正数时，返回最大的三个数之积
		return max1 * max2 * max3
	}

	if zhCount < 3 {
		//正数小于3的时候
		return min1 * min2 * max1
	}

	//最后只剩下一种情况，至少三个正数，至少2个负数了
	if min1*min2 > max3*max2 {
		return min1 * min2 * max1
	}

	return max3 * max2 * max1
}
