package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//每日一题：222. 完全二叉树的节点个数
func countNodes(root *TreeNode) int {
	//最简单的思路就是遍历即可。但效率？
	//直接求第n层个数即可。
	res := 0
	if root != nil {
		res = countNodes(root.Left) + countNodes(root.Right) + 1
	}
	return res
}
