package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(construct([][]int{{1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}}))
}

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

//427. 建立四叉树
func construct(grid [][]int) *Node {

	//看是否都相同
	flag := false
	pre := grid[0][0]
N:
	for _, row := range grid {
		for i := 1; i < len(row); i++ {
			if row[i] != pre {
				flag = true
				break N
			}
		}
	}

	//有不相同的
	if flag {
		//需要切分，先生成一个非叶子节点

		//初始化四个部分
		topLeft := make([][]int, len(grid)/2)
		topRight := make([][]int, len(grid)/2)
		bottomLeft := make([][]int, len(grid)/2)
		bottomRight := make([][]int, len(grid)/2)
		for i := 0; i < len(grid)/2; i++ {
			topLeft[i] = make([]int, len(grid)/2)
			topRight[i] = make([]int, len(grid)/2)
			bottomLeft[i] = make([]int, len(grid)/2)
			bottomRight[i] = make([]int, len(grid)/2)
		}

		for i := 0; i < len(grid)/2; i++ {
			for j := 0; j < len(grid)/2; j++ {
				topLeft[i][j] = grid[i][j]
			}

			for j := len(grid) / 2; j < len(grid); j++ {
				topRight[i][j-len(grid)/2] = grid[i][j]
			}
		}

		for i := len(grid) / 2; i < len(grid); i++ {
			for j := 0; j < len(grid)/2; j++ {
				bottomLeft[i-len(grid)/2][j] = grid[i][j]
			}

			for j := len(grid) / 2; j < len(grid); j++ {
				bottomRight[i-len(grid)/2][j-len(grid)/2] = grid[i][j]
			}
		}
		return &Node{
			Val:         true,
			IsLeaf:      false,
			TopLeft:     construct(topLeft),
			TopRight:    construct(topRight),
			BottomLeft:  construct(bottomLeft),
			BottomRight: construct(bottomRight),
		}
	} else {
		//都相同的话，直接返回一个即可
		val := false
		if grid[0][0] == 1 {
			val = true
		}
		return &Node{
			Val:         val,
			IsLeaf:      true,
			TopLeft:     nil,
			TopRight:    nil,
			BottomLeft:  nil,
			BottomRight: nil,
		}
	}
}

//495. 提莫攻击
func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}
	//先排序时间
	sort.Ints(timeSeries)
	for res, i := 0, 0; ; i++ {
		//最后一次攻击
		if i == len(timeSeries)-1 {
			return res + duration
		}
		//如果它的下一个比当前时间大2及以上
		if timeSeries[i+1] >= timeSeries[i]+duration {
			res += duration
		} else {
			res += timeSeries[i+1] - timeSeries[i]
		}
	}
}

//746. 使用最小花费爬楼梯
func minCostClimbingStairs(cost []int) int {
	//思路：采用动态规划的方式dp[i]为从i处走到顶点所花费的，
	dp := make([]int, len(cost))
	dp[len(cost)-1] = cost[len(cost)-1] // 倒数第二个的最小花费是确定的
	dp[len(cost)-2] = cost[len(cost)-2] // 倒数第二个的最小花费是确定的
	for i := len(cost) - 3; i >= 0; i-- {
		if dp[i+1] > dp[i+2] {
			dp[i] = cost[i] + dp[i+2]
		} else {
			dp[i] = cost[i] + dp[i+1]
		}
	}
	if dp[0] > dp[1] {
		return dp[1]
	}
	return dp[0]
}
