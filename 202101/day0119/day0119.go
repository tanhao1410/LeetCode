package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(findNumberOfLIS([]int{2, 2, 2, 2, 2}))
	fmt.Println(triangleNumber([]int{0, 1, 0, 1}))
	fmt.Println(findShortestSubArray([]int{1, 1, 2, 2, 3, 3, 4, 4}))

	fmt.Println(minCostConnectPoints2([][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}))

}

//每日一题：1584. 连接所有点的最小费用
func minCostConnectPoints3(points [][]int) int {
	res := 0
	//思路：从一个点开始，每一次都找一个最近的，加入其中。直到没有加入进来的为空。
	//joined := map[int]bool{0: true}

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

	//记录点到树的最终距离，先加入第一个节点。
	dis := make([]int, len(points))
	for i := 1; i < len(points); i++ {
		dis[i] = math.MaxInt32
	}

	//下一个循环加入离树最近的点
	checkPoint := 0
	for i := 1; i < len(points); i++ {
		//到树的最近距离是多少，是检查点

		//确定谁离树最近？然后把它加入树中。
		min := math.MaxInt32
		minIndex := -1
		for j := 0; j < len(points); j++ {

			if dis[j] > 0 {
				//更新点到树的距离
				p2TreeDis := getDis(points[checkPoint], points[j])
				if p2TreeDis < dis[j] {
					dis[j] = p2TreeDis
				}

				if dis[j] < min {
					min = dis[j]
					minIndex = j
				}
			}

		}
		//更新检查点
		checkPoint = minIndex
		dis[checkPoint] = 0
		res += min
	}
	return res
}

//每日一题：1584. 连接所有点的最小费用
func minCostConnectPoints2(points [][]int) int {
	res := 0
	//此为最小生成树，采用克鲁斯卡尔加边法。每一次加入最小的边。

	//边的集合，边用四个数来表示。
	edges := make([][]int, (len(points)*(len(points)-1))/2)
	k := 0
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			edges[k] = []int{i, j}
			k++
		}
	}

	getDis := func(i, j int) int {
		res := 0
		if points[i][0]-points[j][0] > 0 {
			res += points[i][0] - points[j][0]
		} else {
			res += points[j][0] - points[i][0]
		}
		if points[i][1]-points[j][1] > 0 {
			res += points[i][1] - points[j][1]
		} else {
			res += points[j][1] - points[i][1]
		}
		return res
	}

	//排序边
	sort.Slice(edges, func(i, j int) bool {
		return getDis(edges[i][0], edges[i][1]) < getDis(edges[j][0], edges[j][1]) //小的在前
	})

	c := make([]int, len(points))
	//每一个点都对应一个并查集，开始时，并查集是自己的序号
	for i := 0; i < len(points); i++ {
		c[i] = i
	}

	//获取点的并查集序号
	getCIndex := func(index int) int {
		for c[index] != index {
			c[index] = c[c[index]]
			index = c[index]
		}
		return index
	}

	//合并两个并查集
	unionC := func(i, j int) {
		iCIndex := getCIndex(i)
		jCIndex := getCIndex(j)
		c[jCIndex] = iCIndex
	}

	already := make(map[int]bool)

	for _, edge := range edges {

		if already[edge[0]] && already[edge[1]] {
			continue
		}
		//判断这个边的两个点是否都在同一个并查集中
		if getCIndex(edge[0]) != getCIndex(edge[1]) {
			//不同的并查集，可以加入进来
			//合并两个并查集
			unionC(edge[0], edge[1])
			res += getDis(edge[0], edge[1])
			already[edge[0]] = true
			already[edge[1]] = true
		}
	}

	return res
}

//697. 数组的度
func findShortestSubArray(nums []int) int {
	//先求度
	m := make(map[int]int)
	du := 0
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			m[nums[i]]++
		} else {
			m[nums[i]] = 1
		}
		if m[nums[i]] > du {
			du = m[nums[i]]
		}
	}
	duNum := make(map[int][]int)
	for k, v := range m {
		if v == du {
			duNum[k] = []int{math.MaxInt32, math.MinInt32}
		}
	}

	//记录每一个数第一此和最后一次出现的位置
	for i := 0; i < len(nums); i++ {
		if indexs, ok := duNum[nums[i]]; ok {
			if i < indexs[0] {
				indexs[0] = i
			}
			if i > indexs[1] {
				indexs[1] = i
			}
		}
	}
	res := math.MaxInt32
	for _, indexs := range duNum {
		if indexs[1]-indexs[0]+1 < res {
			res = indexs[1] - indexs[0] + 1
		}
	}
	return res

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
