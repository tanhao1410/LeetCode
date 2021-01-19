package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(findNumberOfLIS([]int{2, 2, 2, 2, 2}))
	fmt.Println(triangleNumber([]int{0, 1, 0, 1}))
}

//611. 有效三角形的个数
func triangleNumber(nums []int) int {
	//先排序
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			//短边的两边之和
			sum := nums[i] + nums[j]
			//用二分法的方式查找第三个的量
			start, end := i+2, len(nums)-1
			middle := (start + end) / 2
			for ; start <= end; middle = (start + end) / 2 {
				if nums[middle] >= sum {
					end = middle - 1
				} else {
					start = middle + 1
				}
			}
			if end > j {
				res += (end - j)
			}
		}
	}
	return res
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
