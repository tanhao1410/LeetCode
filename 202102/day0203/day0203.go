package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{1, 0, 3}))
}

//173. 二叉搜索树迭代器
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	nums []int
}

func Constructor(root *TreeNode) BSTIterator {
	//遍历树，然后将树中的数字放入到数组中，中序递归。
	nums := []int{}
	var readTree func(root *TreeNode)
	readTree = func(root *TreeNode) {
		if root != nil {
			readTree(root.Right)
			nums = append(nums, root.Val)
			readTree(root.Left)
		}
	}
	readTree(root)
	return BSTIterator{
		nums: nums,
	}

}

func (this *BSTIterator) Next() int {
	res := this.nums[len(this.nums)-1]
	this.nums = this.nums[:len(this.nums)-1]
	return res
}

func (this *BSTIterator) HasNext() bool {
	return len(this.nums) > 0
}

//268. 丢失的数字
func missingNumber(nums []int) int {
	//最简单思路：空间复杂度o(n) ，或者排序完成。
	temp := make([]bool, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		temp[nums[i]] = true
	}

	for i := 0; ; i++ {
		if !temp[i] {
			return i
		}
	}

}

//292. Nim 游戏
func canWinNim(n int) bool {
	return n%4 == 0
}
