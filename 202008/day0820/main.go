package day0820

import "sort"

func intersect(nums1 []int, nums2 []int) []int {

	n1 := sort.IntSlice(nums1)
	sort.Sort(n1) //对n1进行了排序
	//res := make([]int,10)
	res := []int{}
	//两层循环
	for _, v1 := range nums1 {
		//通过判断v1是否在nums2中不可以，在nums2中用过一次的数就不能再用了。
		for k2, v2 := range nums2 {
			if v1 == v2 {
				res = append(res, v1)
				if k2 < len(nums2)-1 {
					left := nums2[0:k2]
					right := nums2[k2+1:]
					nums2 = append(left, right...)
				} else {
					nums2 = nums2[0:k2]
				}
				break
			}
		}
	}

	return res
}

//求二叉树的直径
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if head == nil || root == nil {
		return false
	}
	//思路，遍历树
	res := false

	ReadTree(root, head, &res)
	return res
}

func ReadTree(root *TreeNode, head *ListNode, res *bool) {
	if root != nil {
		//说明匹配上了。
		if root.Val == head.Val {
			if head.Next == nil {
				//说明是最后一个了
				*res = true
				return
			} else {
				ReadTree(root.Left, head.Next, res)
				ReadTree(root.Right, head.Next, res)
			}
		} else {
			ReadTree(root.Right, head, res)
			ReadTree(root.Left, head, res)
		}

	}
}

//回文链表， O(n) 时间复杂度和 O(1) 空间复杂度
func isPalindrome(head *ListNode) bool {
	//遍历一遍，找到长度，从中间开始
	return false
}
