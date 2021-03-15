package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//面试题 04.12. 求和路径
func pathSum(root *TreeNode, sum int) int {
	//递归，每次返回到一个节点时，会传过来一个已有的值。如果加上val == sum，那么结果+1，传递给下一层的时候，加上自身
	res := 0
	pathSumContainParents(root, sum, 0, &res)

	//求它的左右节点的
	if root.Left != nil {
		res += pathSum(root.Left, sum)
	}
	if root.Right != nil {
		res += pathSum(root.Right, sum)
	}

	return res
}

func pathSumContainParents(root *TreeNode, sum int, already int, count *int) {
	if root == nil {
		return
	}

	newAlready := already + root.Val
	if newAlready == sum {
		*count += 1
	}
	pathSumContainParents(root.Left, sum, newAlready, count)
	pathSumContainParents(root.Right, sum, newAlready, count)

}
