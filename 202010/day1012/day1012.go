package main

import "math"

func main() {

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
	if root == nil{
		return math.MaxInt32
	}
	leftDif := math.MaxInt32
	rightDif := math.MaxInt32

	if root.Left != nil{
		p := root.Left
		for ; p.Right != nil;p = p.Right{
		}
		leftDif = root.Val - p.Val
	}
	if root.Right != nil{
		p := root.Right
		for ; p.Left != nil;p = p.Left{
		}
		rightDif = p.Val - root.Val
	}
	leftTreeDif := getMinimumDifference(root.Left)
	rightTreeDif := getMinimumDifference(root.Right)

	res := leftDif
	if rightDif < res{
		res = rightDif
	}
	if leftTreeDif < res{
		res = leftTreeDif
	}
	if rightTreeDif < res{
		res = rightTreeDif
	}
	return res
}
