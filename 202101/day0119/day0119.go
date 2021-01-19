package main

import "math"

func main() {

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
