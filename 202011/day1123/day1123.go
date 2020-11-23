package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMinArrowShots([][]int{{9, 12}, {1, 10}, {4, 11}, {8, 12}, {3, 9}, {6, 9}, {6, 7}}))
}

//每日一题：452. 用最少数量的箭引爆气球
func findMinArrowShots(points [][]int) int {
	//思路：先排序，按照start排序，如果起点相同，那么短的放在前面
	sort.Slice(points, func(i, j int) bool {
		if points[j][0] > points[i][0] {
			return true
		} else if points[j][0] < points[i][0] {
			return false
		} else {
			return points[j][1] > points[i][1]
		}
	})

	res := 0
	//从第一个的末端射出
	for i := 0; i < len(points); {

		//先求end的值，end的值并不一定是第一个线的末尾。
		//而应该是开始小于第一个end的所有的线段的最小的那个。
		end := points[i][1]
		for j := i; j < len(points) && points[j][0] <= points[i][1]; j++ {
			if points[j][1] < end {
				end = points[j][1]
			}
		}

		res += 1
		//只要下一个线段的开始比射出的地方小，那么就算刺中了
		//问题，如果下面的结束比end小就不行了。
		for ; i < len(points) && points[i][0] <= end; i++ {
		}

	}

	return res
}
