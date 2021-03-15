package main

func main() {

}

//面试题 05.01. 插入
func insertBits(N int, M int, i int, j int) int {
	//i - j 位置改为0
	num := 0
	for k := 0; k < 32; k++ {
		if k > j || k < i {
			num += 1 << k
		}
	}
	num &= N
	//求插入的数M，i-0代表M应该右移的位置数
	m := M << i
	res := int32(num | m)
	return int(res)
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
