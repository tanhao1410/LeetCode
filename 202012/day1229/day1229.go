package main

import "fmt"

func main() {
	fmt.Println(pathWithObstacles([][]int{{0, 1, 0}, {0, 1, 1}, {0, 0, 0}}))
	fmt.Println(minPatches([]int{1, 1, 2, 3, 5, 10}, 200))
	fmt.Println(smallestK([]int{4, 5}, 4))
}

//面试题 17.14. 最小K个数
func smallestK(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return nil
	}
	//思路：通过快速排序的思想进行
	//1.先根据第一个数进行划分，将数组划分成两段，前面都是小于它的，后面都是大于等于它的
	i, j := 1, len(arr)-1
	for j >= i {

		for i < len(arr) && arr[i] < arr[0] {
			i++
		}

		for j > 0 && arr[j] >= arr[0] {
			j--
		}

		//交换
		if j > i {
			arr[i], arr[j] = arr[j], arr[i]
		}

	}
	if j >= 0 {
		arr[0], arr[j] = arr[j], arr[0]
	}

	if j == k-1 {
		return arr[:k]
	} else if j > k-1 {
		return smallestK(arr[:j], k)
	} else {
		return append(arr[:j+1], smallestK(arr[j+1:], k-j-1)...)
	}
}

//每日一题：330. 按要求补齐数组
func minPatches(nums []int, n int) int {
	res := 0
	//设开始时的覆盖范围是x = 1,index = 0
	//这里的x代表的是小于x的都可以，所以for循环中是x<=n
	x, index := 1, 0
	for x <= n {
		//可以直接扩大覆盖范围
		if index < len(nums) && nums[index] <= x {
			x += nums[index]
			index++
		} else {
			//不够了，需要补充数字了，补充x
			res++
			x += x
		}
	}
	return res
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
