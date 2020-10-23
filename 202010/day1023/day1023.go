package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//105.从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	var root *TreeNode = nil
	//前序遍历的第一个节点为根节点
	if len(preorder) == 0{
		return root
	}
	root = &TreeNode{preorder[0],nil,nil}

	i:=0
	for ;i < len(inorder) && inorder[i] != preorder[0];i++{
	}
	leftInorder := inorder[:i]
	leftPreorder := preorder[1:len(leftInorder)+1]
	root.Left = buildTree(leftPreorder,leftInorder)
	rightInorder := inorder[i+1:]
	rightPreorder := preorder[len(preorder) - len(rightInorder):]
	root.Right = buildTree(rightPreorder,rightInorder)
	return root
}
