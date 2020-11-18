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

	//思路：
	res := ""
	//把数组排序
	//排序的方式数字中最大的数，以及最大数所在的位置，位置越往后，说明它越小
	//如果最大数一样，位置也一样，
	sort.Slice(nums, func(i int, j int) bool {
		//num1,num2 := nums[i],nums[j]
		//直接比较即可
		num1 := strconv.Itoa(nums[i])
		num2 := strconv.Itoa(nums[j])
		//先从第一位开始比较，都没有比出大小的话
		for i := 0; i < len(num1) && i < len(num2); i++ {
			if num1[i] > num2[i] {
				return true
			} else if num1[i] < num2[i] {
				return false
			}
		}
		//
		return len(num1) < len(num2)
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
