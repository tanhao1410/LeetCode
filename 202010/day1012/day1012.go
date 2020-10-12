package main

import (
	"fmt"
	"math"
)

func main() {
	nums := [][]int{{0,0,0}, {0,1,0},{0,0,0}}
	fmt.Println(uniquePathsWithObstacles(nums))
}

//63. 不同路径 II
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}
	n := len(obstacleGrid)

	if len(obstacleGrid[0]) == 0 {
		return 1
	}
	m := len(obstacleGrid[0])

	//永远到不了
	if obstacleGrid[n-1][m-1] == 1{
		return 0
	}

	obstacleGrid[n-1][m-1] = 1
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if obstacleGrid[i][j] == 1 {
				if i == n-1 && j == m-1 {
					obstacleGrid[i][j] = 1
				} else {
					obstacleGrid[i][j] = 0
				}
			} else {
				if i+1 < n {
					obstacleGrid[i][j] += obstacleGrid[i+1][j]
				}
				if j+1 < m {
					obstacleGrid[i][j] += obstacleGrid[i][j+1]
				}
			}
		}
	}

	return obstacleGrid[0][0]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//530.二叉搜索树的最小绝对差
func getMinimumDifference(root *TreeNode) int {
	//思路：最小的肯定是它的左右子树最小的，或它和左右孩子之间的最小
	//任意两节点的值的绝对值差，左子树的最大值，右子树的最小值。
	if root == nil {
		return math.MaxInt32
	}
	leftDif := math.MaxInt32
	rightDif := math.MaxInt32

	if root.Left != nil {
		p := root.Left
		for ; p.Right != nil; p = p.Right {
		}
		leftDif = root.Val - p.Val
	}
	if root.Right != nil {
		p := root.Right
		for ; p.Left != nil; p = p.Left {
		}
		rightDif = p.Val - root.Val
	}
	leftTreeDif := getMinimumDifference(root.Left)
	rightTreeDif := getMinimumDifference(root.Right)

	res := leftDif
	if rightDif < res {
		res = rightDif
	}
	if leftTreeDif < res {
		res = leftTreeDif
	}
	if rightTreeDif < res {
		res = rightTreeDif
	}
	return res
}
