package main

import "fmt"

func main() {
	nums := []int{1, 0, 1, 2, 0, 0}
	sortColors(nums)
	fmt.Println(nums)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//每日一题：222. 完全二叉树的节点个数
func countNodes(root *TreeNode) int {
	//最简单的思路就是遍历即可。但效率？
	//直接求第n层个数即可。
	res := 0
	if root != nil {
		res = countNodes(root.Left) + countNodes(root.Right) + 1
	}
	return res
}

//75. 颜色分类
func sortColors(nums []int) {
	//思路：计数法
	zero, one := 0, 0
	for _, v := range nums {
		if v == 0 {
			zero++
		} else if v == 1 {
			one++
		}
	}

	for i := 0; i < len(nums); i++ {
		if i < zero {
			nums[i] = 0
		} else if i < zero+one {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}
}
