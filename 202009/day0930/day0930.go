package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//二叉搜索树中的插入操作
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	//思路：递归法，如果比root大，往右边走，右边为空，插入到右边，否则递归
	if root == nil{
		return &TreeNode{val,nil,nil}
	}
	if root.Val > val{

		//比根小，说明应该插入到左边，或往下继续传递
		if root.Left == nil{
			newNode := &TreeNode{val,nil,nil}
			root.Left = newNode
		}else{
			insertIntoBST(root.Left,val)
		}
	}else{
		if root.Right == nil{
			newNode := &TreeNode{val,nil,nil}
			root.Right = newNode
		}else{
			insertIntoBST(root.Right,val)
		}
	}

	return root
}
