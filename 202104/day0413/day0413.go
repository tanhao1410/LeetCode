package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//每日一题：783. 二叉搜索树节点最小距离
func minDiffInBST(root *TreeNode) int {
	res := 100000
	pre := -100000
	var middleBST func(root *TreeNode)
	middleBST = func(root *TreeNode) {
		if root == nil {
			return
		}
		middleBST(root.Left)
		if root.Val-pre < res {
			res = root.Val - pre
		}
		pre = root.Val
		middleBST(root.Right)
	}
	middleBST(root)

	return res
}
