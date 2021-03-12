package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(isValidSerialization2("9,3,4,#,#,1,#,#,2,#,6,#,#"))
}

//面试题 04.04. 检查平衡性
func isBalanced(root *TreeNode) bool {
	//求左子树高度，求右子树高度，两者如果差距大于1，返回false
	//怎么样效率比较高呢，先求子树
	_, res := heightAndBalanced(root)
	return res
}

func heightAndBalanced(root *TreeNode) (int, bool) {

	if root == nil {
		return 0, true
	}

	leftHeight, leftBalanced := heightAndBalanced(root.Left)
	rightHeight, rightBalaced := heightAndBalanced(root.Right)

	resHeight := 0
	resBalanced := true
	//高度由最高的决定，
	if leftHeight > rightHeight {
		resHeight = leftHeight + 1
		resBalanced = (leftHeight-resHeight) < 2 && leftBalanced && rightBalaced
	} else {
		resHeight = rightHeight + 1
		resBalanced = (rightHeight-leftHeight) < 2 && leftBalanced && rightBalaced
	}

	return resHeight, resBalanced
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//面试题 04.10. 检查子树
func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t2 == nil {
		return true
	}

	middleT1 := serializedTree(t1)
	middleT2 := serializedTree(t2)
	preT1 := serializedTree2(t1)
	pret2 := serializedTree2(t2)

	return strings.Contains(middleT1, middleT2) && strings.Contains(preT1, pret2)

}

//中序遍历
func serializedTree(root *TreeNode) string {
	res := ""
	if root != nil {
		if root.Left != nil {
			res += serializedTree(root.Left)
		}
		res += ","
		res += strconv.Itoa(root.Val)
		res += ","
		if root.Right != nil {
			res += serializedTree(root.Right)
		}
	}
	return strings.Trim(res, ",")
}

func serializedTree2(root *TreeNode) string {
	res := ""
	if root != nil {
		res += strconv.Itoa(root.Val)
		res += ","
		if root.Left != nil {
			res += serializedTree2(root.Left)
		}
		res += ","
		if root.Right != nil {
			res += serializedTree2(root.Right)
		}
	}
	return strings.Trim(res, ",")
}

//331. 验证二叉树的前序序列化
func isValidSerialization(preorder string) bool {
	if preorder == "#" {
		return true
	}
	//思路：一个数字后面紧接两个#,那么他们三个可以看做一个#，如果最后整个字符串变为一个#，则为true
	nodes := strings.Split(preorder, ",")
	for i := 0; i < len(nodes); i++ {
		//是数字的话，看下它紧接着的后面两个是否是#
		if nodes[i] != "#" && i+2 < len(nodes) && nodes[i+1] == "#" && nodes[i+2] == "#" {
			newPreorder := strings.Join(nodes[:i], ",")
			newPreorder += ","
			newPreorder += strings.Join(nodes[i+2:], ",")
			return isValidSerialization(strings.Trim(newPreorder, ","))
		}
	}

	return false
}

//331. 验证二叉树的前序序列化
func isValidSerialization2(preorder string) bool {
	if preorder == "#" {
		return true
	}
	//思路：一个数字后面紧接两个#,那么他们三个可以看做一个#，如果最后整个字符串变为一个#，则为true
	nodes := strings.Split(preorder, ",")
	newPreorder := ""
	for i := 0; i < len(nodes); i++ {
		//是数字的话，看下它紧接着的后面两个是否是#
		if nodes[i] != "#" && i+2 < len(nodes) && nodes[i+1] == "#" && nodes[i+2] == "#" {
			newPreorder += "#"
			//跳过两个
			i += 2
		} else {
			newPreorder += nodes[i]
		}
		if i != len(nodes)-1 {
			newPreorder += ","
		}
	}
	fmt.Println(newPreorder)
	//没有一个可以消除的，直接返回false
	if newPreorder == preorder {
		return false
	}

	return isValidSerialization2(newPreorder)
}
