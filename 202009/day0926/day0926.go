package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//二叉树的前序遍历 ：中左右
func preorderTraversal(root *TreeNode) []int {
	res := &[]int{}
	preorderTraversal2(root,res)
	return *res
}

func preorderTraversal2(root *TreeNode,res *[]int){
	if root == nil{
		return
	}
	*res = append(*res, root.Val)
	preorderTraversal2(root.Left,res)
	preorderTraversal2(root.Right,res)
}
//二叉树的后序遍历 ：左右中
func postorderTraversal(root *TreeNode) []int {
	res := &[]int{}
	postorderTraversal2(root,res)
	return *res
}

func postorderTraversal2(root *TreeNode,res *[]int){
	if root == nil{
		return
	}
	postorderTraversal2(root.Left,res)
	postorderTraversal2(root.Right,res)
	*res = append(*res, root.Val)
}

//路径总和2
func pathSum(root *TreeNode, sum int) [][]int {
	//从根到叶子节点总和为目标数
	res := &[][]int{}
	if root == nil{
		return *res
	}
	item := []int{}
	pathSum2(root,sum,item,res)
	return *res
}

func pathSum2(root *TreeNode,sum int, item []int,res *[][]int){

	//root为叶子节点
	if root.Right == nil && root.Left == nil{
		if root.Val == sum{
			item = append(item, root.Val)
			*res = append(*res, item)
		}
		return
	}

	if root.Left != nil{
		itemLeft :=make([]int,len(item))
		copy(itemLeft,item)
		itemLeft = append(itemLeft, root.Val)
		pathSum2(root.Left,sum - root.Val,itemLeft,res)
	}

	if root.Right != nil{
		itemRight := make([]int,len(item))
		copy(itemRight,item)
		itemRight = append(itemRight, root.Val)
		pathSum2(root.Right,sum - root.Val,itemRight,res)
	}

}
