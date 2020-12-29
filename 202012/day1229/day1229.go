package main

import "fmt"

func main() {
	fmt.Println(pathWithObstacles([][]int{{0, 1, 0}, {0, 1, 1}, {0, 0, 0}}))
}

//面试题 08.02. 迷路的机器人
func pathWithObstacles(obstacleGrid [][]int) [][]int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return nil
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	res := [][]int{}
	//返回一条可行的即可，dp代表某个结点是否能走到终点。
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	copy(dp, obstacleGrid)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dp[i][j] = obstacleGrid[i][j]
			} else if obstacleGrid[i][j] == 0 && ((i < m-1 && dp[i+1][j] == 0) || (j < n-1 && dp[i][j+1] == 0)) {
				dp[i][j] = 0
			} else {
				dp[i][j] = 1
			}
		}
	}
	if dp[0][0] == 1 {
		return nil
	}
	for i, j := 0, 0; i < m-1 && j < n-1; {
		res = append(res, []int{i, j})
		//向下走
		if i+1 < m && dp[i+1][j] == 0 {
			i = i + 1
		} else if j+1 < n && dp[i][j+1] == 0 {
			//向右走
			j = j + 1
		}
	}
	res = append(res, []int{m - 1, n - 1})
	return res
}
