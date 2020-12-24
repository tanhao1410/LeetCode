package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 2, 4, 4, 3}
	fmt.Println(candy(nums))
}

//每日一题：135. 分发糖果
func candy(ratings []int) int {
	if len(ratings) < 2 {
		return len(ratings)
	}
	//思路：
	dp := make([]int, len(ratings))

	count := 0
	//如果一个数比两边都小，则它分配1个
	for i := 1; i < len(ratings)-1; i++ {
		if ratings[i] <= ratings[i-1] && ratings[i] <= ratings[i+1] {
			dp[i] = 1
			count++
		}
	}
	if ratings[1] >= ratings[0] {
		dp[0] = 1
		count++
	}
	if ratings[len(ratings)-1] <= ratings[len(ratings)-2] {
		dp[len(ratings)-1] = 1
		count++
	}

	for count < len(ratings) {
		//先给分数最小的分配一个，存在多个最小的key
		min := math.MaxInt32
		for k, v := range ratings {
			if dp[k] == 0 && v < min {
				min = v
			}
		}
		for k, v := range ratings {
			if v == min && dp[k] == 0 {
				//dp[k] = 1
				if (k+1 < len(ratings) && dp[k+1] != 0) || (k-1 > -1 && dp[k-1] != 0) {
					//它的两边已经有分配了
					//看是否有相等的，如果有相等的，先按相等的来
					if k+1 < len(ratings) && dp[k+1] != 0 && ratings[k+1] == ratings[k] {
						dp[k] = dp[k+1]
						if k-1 > -1 && ratings[k] > ratings[k-1] && (dp[k] <= dp[k-1] || dp[k] > dp[k-1]+1) {
							dp[k] = dp[k-1] + 1
						}
					} else if k-1 > -1 && dp[k-1] != 0 && ratings[k-1] == ratings[k] {
						dp[k] = dp[k-1]
						if k+1 < len(ratings) && ratings[k] > ratings[k+1] && (dp[k] <= dp[k+1] || dp[k] > dp[k+1]+1) {
							dp[k] = dp[k+1] + 1
						}
					} else {
						value := 0
						if k+1 < len(ratings) {
							value = dp[k+1]
						}
						if k-1 > -1 && dp[k-1] > value {
							value = dp[k-1]
						}
						dp[k] = value + 1
					}
					count++
				} else {
					dp[k] = 1
					count++
				}
			}
		}
	}

	res := 0
	for _, v := range dp {
		res += v
	}
	return res
}
