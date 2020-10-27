package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//每日一题：144.二叉树的前序遍历
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil{
		return res
	}
	stack := []*TreeNode{root}
	stackLen := len(stack)
	for stackLen > 0{
		//出栈
		node := stack[stackLen-1]
		stack = stack[:stackLen -1]
		res = append(res, node.Val)
		if node.Right != nil{
			stack = append(stack, node.Right)
		}
		if node.Left != nil{
			stack = append(stack,node.Left)
		}
		stackLen = len(stack)
	}
	return res
}
