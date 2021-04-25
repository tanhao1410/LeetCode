package main

import "fmt"

func main() {
	one, two, tree, four := &TreeNode{Val: 1}, &TreeNode{Val: 2}, &TreeNode{Val: 3}, &TreeNode{Val: 4}
	two.Left = one
	two.Right = four
	four.Left = tree
	res := increasingBST(two)
	fmt.Print(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//897. 递增顺序搜索树
func increasingBST(root *TreeNode) *TreeNode {

	var res *TreeNode = nil
	var temp *TreeNode = nil

	var middleRead func(root *TreeNode)
	middleRead = func(root *TreeNode) {
		if root == nil {
			return
		}
		middleRead(root.Left)
		if temp == nil {
			res = root
			temp = root
		} else {
			temp.Right = root
			root.Left = nil
			temp = root
		}

		middleRead(root.Right)
	}

	middleRead(root)

	return res
}
