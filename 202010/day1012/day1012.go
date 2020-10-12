package main

import (
	"fmt"
	"math"
)

func main() {
	//nums := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	//fmt.Println(uniquePathsWithObstacles(nums))
	fmt.Println(simplifyPath("/..aa/...hidden/"))
}

//有大于两个点的话算路径名
//71.简化路径 "/a//b////c/d//././/../"-->/a/b/c
func simplifyPath(path string) string {
	res := []byte{'/'}
	pre := path[0]
	for i := 1; i < len(path); i++ {
		//消除多余的“/”
		if path[i] == '/' && path[i] == pre {
			continue
		}

		//处理./
		if path[i] == '/' && pre == '.' {
			pre = path[i]
			continue
		}

		//处理../
		if path[i] == '.' && pre == '.' {
			i++ // 看..后面是什么
			if (i < len(path) && path[i] == '/') || i == len(path) {
				//要从结果中取出上一级
				if len(res) != 1 {
					j := len(res) - 2
					for ; res[j] != '/' && i >= 0; j-- {
					}
					res = res[0 : j+1]
				}
			} else {
				res = append(res, '.', '.')
				for ; i < len(path) && path[i] != '/'; i++ {
					res = append(res, path[i])
				}
				res = append(res, '/')
			}
			pre = '/'
			continue
		}

		//以.作为文件名
		if pre == '.' && path[i] != '/' {
			res = append(res, '.')
		}

		if path[i] != '.' {
			res = append(res, path[i])
		}

		pre = path[i]
	}

	if len(res) != 1 && res[len(res)-1] == '/' {
		res = res[0 : len(res)-1]
	}
	return string(res)
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
	if obstacleGrid[n-1][m-1] == 1 {
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
