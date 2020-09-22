package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//合并二叉树
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	//明显的递归的方式，左子树与右子树分别与t2的左右子树合并
	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	t1.Val += t2.Val
	//递归调用
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)

	return t1
}
