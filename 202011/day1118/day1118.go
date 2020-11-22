package main

import (
	"sort"
	"strconv"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//剑指 Offer 55 - I. 二叉树的深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

//剑指 Offer 27. 二叉树的镜像
func mirrorTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = root.Right, root.Left
		mirrorTree(root.Left)
		mirrorTree(root.Right)
	}
	return root
}

//剑指 Offer 45. 把数组排成最小的数
func minNumber(nums []int) string {
	res := ""
	sort.Slice(nums, func(i int, j int) bool {
		num1 := strconv.Itoa(nums[i])
		num2 := strconv.Itoa(nums[j])
		//组合起来比较即可
		s1 := num1 + num2
		s2 := num2 + num1
		//先从第一位开始比较，都没有比出大小的话
		for i := 0; i < len(s1); i++ {
			if s2[i] > s1[i] {
				return true
			} else if s2[i] < s1[i] {
				return false
			}
		}
		return false
	})
	for i := 0; i < len(nums); i++ {
		res += strconv.Itoa(nums[i])
	}
	return res
}

//每日一题：134. 加油站
func canCompleteCircuit(gas []int, cost []int) int {
	//思路：可以从那一站出发，并走一圈，如果不行，下一个，都不行，返回-1
	for i := 0; i < len(gas); i++ {
		all := gas[i] //起始油量
		cur := i
		for j := (i + 1) % len(gas); cost[cur] <= all; j = (j + 1) % len(gas) {
			if j == i {
				return i
			}
			all += gas[j]    //补充油
			all -= cost[cur] //消耗油
			cur = j
		}
	}

	return -1
}
