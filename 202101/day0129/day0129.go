package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumEffortPath([][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}}))
}

//每日一题：1631. 最小体力消耗路径
func minimumEffortPath(heights [][]int) int {

	//采用动态规划算法，每更新一个位置的体力消耗后，都要看能否更新它周围的，直到它周围的不能改变了为止

	//初始化动态规划表
	dp := make([][]int, len(heights))
	for i := 0; i < len(heights); i++ {
		dp[i] = make([]int, len(heights[0]))
		for j := 0; j < len(heights[0]); j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 0

	getAbstractNum := func(n1, n2 int) int {
		if n1 > n2 {
			return n1 - n2
		}
		return n2 - n1
	}

	max := func(n1, n2 int) int {
		if n1 > n2 {
			return n1
		}
		return n2
	}

	var updateDp func(x, y int)
	updateDp = func(i, j int) {

		//它的上方是已经可以到达的
		curNum := heights[i][j]
		if i > 0 && dp[i-1][j] != math.MaxInt32 {
			if max(getAbstractNum(curNum, heights[i-1][j]), dp[i][j]) < dp[i-1][j] {
				dp[i-1][j] = max(getAbstractNum(curNum, heights[i-1][j]), dp[i][j])
				//递归更新它周围的
				updateDp(i-1, j)
			}
		}
		//从下方
		if i < len(dp)-1 && dp[i+1][j] != math.MaxInt32 {
			if max(getAbstractNum(curNum, heights[i+1][j]), dp[i][j]) < dp[i+1][j] {
				dp[i+1][j] = max(getAbstractNum(curNum, heights[i+1][j]), dp[i][j])
				//递归更新它周围的
				updateDp(i+1, j)
			}
		}

		//从左方
		if j > 0 && dp[i][j-1] != math.MaxInt32 {
			if max(getAbstractNum(curNum, heights[i][j-1]), dp[i][j]) < dp[i][j-1] {
				dp[i][j-1] = max(getAbstractNum(curNum, heights[i][j-1]), dp[i][j])
				//递归更新它周围的
				updateDp(i, j-1)
			}
		}

		//从右方
		if j < len(dp[0])-1 && dp[i][j+1] != math.MaxInt32 {
			if max(getAbstractNum(curNum, heights[i][j+1]), dp[i][j]) < dp[i][j+1] {
				dp[i][j+1] = max(getAbstractNum(curNum, heights[i][j+1]), dp[i][j])
				//递归更新它周围的
				updateDp(i, j+1)
			}
		}
	}

	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[0]); j++ {
			//与矩阵前后左右比较，如果位置已经到达，

			curRes := dp[i][j]
			curNum := heights[i][j]

			//它的上方是已经可以到达的
			if i > 0 && dp[i-1][j] != math.MaxInt32 {
				if max(getAbstractNum(curNum, heights[i-1][j]), dp[i-1][j]) < curRes {
					curRes = max(getAbstractNum(curNum, heights[i-1][j]), dp[i-1][j])
				}
			}
			//从下方
			if i < len(dp)-1 && dp[i+1][j] != math.MaxInt32 {
				if max(getAbstractNum(curNum, heights[i+1][j]), dp[i+1][j]) < curRes {
					curRes = max(getAbstractNum(curNum, heights[i+1][j]), dp[i+1][j])
				}
			}

			//从左方
			if j > 0 && dp[i][j-1] != math.MaxInt32 {
				if max(getAbstractNum(curNum, heights[i][j-1]), dp[i][j-1]) < curRes {
					curRes = max(getAbstractNum(curNum, heights[i][j-1]), dp[i][j-1])
				}
			}

			//从右方
			if j < len(dp[0])-1 && dp[i][j+1] != math.MaxInt32 {
				if max(getAbstractNum(curNum, heights[i][j+1]), dp[i][j+1]) < curRes {
					curRes = max(getAbstractNum(curNum, heights[i][j+1]), dp[i][j+1])
				}
			}

			dp[i][j] = curRes

			//更新curRes后，还需要更新它周围的dp
			updateDp(i, j)
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]

}
