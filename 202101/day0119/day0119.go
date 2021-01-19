package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findNumberOfLIS([]int{2, 2, 2, 2, 2}))
}

//673. 最长递增子序列的个数
func findNumberOfLIS(nums []int) int {

	//思路：以i为结尾的最长子序列长度，
	max := 0
	dp := make([]int, len(nums))
	road := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		//往前比较，只要遇到比当前数大的就跳过去
		preMax := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && preMax < dp[j] {
				preMax = dp[j]
			}
		}
		roadNum := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && preMax == dp[j] {
				roadNum += road[j]
			}
		}
		if roadNum == 0 {
			roadNum = 1
		}
		road[i] = roadNum
		dp[i] = preMax + 1
		if dp[i] > max {
			max = dp[i]
		}
	}

	res := 0

	for i := 0; i < len(dp); i++ {
		if dp[i] == max {
			res += road[i]
		}
	}
	return res
}

//每日一题：1584. 连接所有点的最小费用
func minCostConnectPoints(points [][]int) int {
	res := 0
	//思路：从一个点开始，每一次都找一个最近的，加入其中。直到没有加入进来的为空。
	joined := map[int]bool{0: true}

	getDis := func(point1, point2 []int) int {
		res := 0
		if point1[0]-point2[0] > 0 {
			res += point1[0] - point2[0]
		} else {
			res += point2[0] - point1[0]
		}
		if point1[1]-point2[1] > 0 {
			res += point1[1] - point2[1]
		} else {
			res += point2[1] - point1[1]
		}
		return res
	}

	for len(joined) != len(points) {
		minDis := math.MaxInt32
		minDisIndex := -1
		for joinedPoint, _ := range joined {
			for i := 0; i < len(points); i++ {
				if !joined[i] && getDis(points[i], points[joinedPoint]) < minDis {
					minDisIndex = i
					minDis = getDis(points[i], points[joinedPoint])
				}
			}
		}
		res += minDis
		joined[minDisIndex] = true
	}

	return res
}
