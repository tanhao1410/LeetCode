package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(tail(6, 0, 0, 0))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//面试题 04.02. 最小高度树
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	//中间的作为新的根节点
	middle := len(nums) / 2
	root := TreeNode{
		Val:   nums[middle],
		Left:  sortedArrayToBST(nums[:middle]),
		Right: sortedArrayToBST(nums[middle+1:]),
	}
	return &root
}

func tail(n, temp, pre1, pre2 int) int {
	if n+1 == temp {
		//结束了
		fmt.Println(pre1 + pre2)
		return pre1 + pre2
	}

	if temp == 1 {
		return tail(n, temp+1, 0, 0)
	} else if temp == 2 {
		return tail(n, temp+1, 1, 0)
	} else {
		return tail(n, temp+1, pre1+pre2, pre1)
	}

}

//面试题 16.19. 水域大小
func pondSizes(land [][]int) []int {

	if len(land) == 0 || len(land[0]) == 0 {
		return nil
	}

	//计算所有池塘的大小，然后排序返回
	res := []int{}

	//求池塘的面积
	var getSum func(i, j int) int
	getSum = func(i, j int) int {
		if i < 0 || j < 0 || i >= len(land) || j >= len(land[0]) || land[i][j] != 0 {
			return 0
		}
		//表示计算过了
		land[i][j] = -1
		res := 1
		res += getSum(i+1, j)
		res += getSum(i-1, j)
		res += getSum(i, j+1)
		res += getSum(i, j-1)
		res += getSum(i+1, j+1)
		res += getSum(i+1, j-1)
		res += getSum(i-1, j+1)
		res += getSum(i-1, j-1)
		return res
	}

	for i := 0; i < len(land); i++ {

		for j := 0; j < len(land[0]); j++ {
			if land[i][j] == 0 {
				//说明碰到了池塘
				res = append(res, getSum(i, j))
			}
		}
	}

	sort.Ints(res)
	return res
}

//面试题 08.01. 三步问题
func waysToStep(n int) int {
	if n <= 2 {
		return n
	}
	// return waysToStep(n - 3)  + waysToStep(n - 2) + waysToStep(n - 1)

	//采用数组
	dp := make([]int, n+1)
	dp[1], dp[2], dp[3] = 1, 2, 4
	for i := 4; i < n+1; i++ {
		dp[i] = (dp[i-1] + dp[i-2] + dp[i-3]) % 1000000007
	}
	return dp[n]
}
