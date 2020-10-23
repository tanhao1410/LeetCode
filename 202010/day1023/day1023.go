package main

func main() {

}

//114.二叉树展开为链表
func flatten(root *TreeNode)  {
	//递归方式，将左边展开，然后放入到右边
	if root == nil{
		return
	}

	if root.Left == nil{
		flatten(root.Right)
		return
	}

	flatten(root.Left)
	flatten(root.Right)
	temp := root.Right
	root.Right,root.Left = root.Left,nil
	for root = root.Right;root.Right != nil;root= root.Right{
	}
	root.Right = temp

}

//108.将有序数组转换为二叉搜索树
func sortedArrayToBST(nums []int) *TreeNode {
	//将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树
	var root *TreeNode
	if len(nums) == 0{
		return root
	}
	root = &TreeNode{nums[len(nums)/2],nil,nil}
	root.Right = sortedArrayToBST(nums[len(nums)/2+1:])
	root.Left = sortedArrayToBST(nums[:len(nums)/2])
	return root
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
