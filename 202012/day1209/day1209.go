package main

import (
	"fmt"
)

func main() {
	fmt.Println(uniquePaths(1, 1))
}

//207. 课程表
func canFinish(numCourses int, prerequisites [][]int) bool {
	//思路：先解决没有前驱的课程。如果最后没有课程了，说明都完成了，否则，完成不了
	course := make([][]int, numCourses)
	//图的表示，谁指向自己，就加入自己的集合中
	for _, v := range prerequisites {
		course[v[1]] = append(course[v[1]], v[0])
	}
	haveDel := make(map[int]bool)
	for {
		//去除没有前提的课程
		i := 0
		for ; i < numCourses; i++ {
			if len(course[i]) == 0 && !haveDel[i] {
				break
				//说明该课程应被删除
				//所有需要先学习该课程的前提课程中，都删除该课程
			}
		}

		if i == numCourses {
			//如果一轮结束下来，没有一门课程被删除，说明循环结束
			break
		}
		//开始删除
		for j := 0; j < len(course); j++ {
			for k := 0; k < len(course[j]); k++ {
				if course[j][k] == i {
					course[j] = append(course[j][:k], course[j][k+1:]...)
					break
				}
			}
		}
		//需要记录哪些是已经删除了的
		haveDel[i] = true
	}
	for i := 0; i < numCourses; i++ {
		if len(course[i]) > 0 {
			return false
		}
	}

	return true
}

//343. 整数拆分
func integerBreak(n int) int {
	// n 不小于 2 且不大于 58。
	//特殊处理n = 2,n = 3 情况
	if n < 4 {
		return n - 1
	}
	dp := []int{1, 1, 2, 3}
	for i := 4; i < n+1; i++ {
		max := 0
		for j := 2; j < i; j++ {
			if dp[j]*dp[i-j] > max {
				max = dp[j] * dp[i-j]
			}
		}
		dp = append(dp, max)
	}
	return dp[n]
}

//每日一题：62. 不同路径
func uniquePaths(m int, n int) int {
	//思路：动态规划做法
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[m-1][n-1] = 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			//如果它下面可以走的话
			if i < m-1 {
				dp[i][j] += dp[i+1][j]
			}
			if j < n-1 {
				dp[i][j] += dp[i][j+1]
			}
		}
	}
	return dp[0][0]
}
