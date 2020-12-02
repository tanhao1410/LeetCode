package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxNumber([]int{6, 7}, []int{6, 0, 4}, 5))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//236. 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//如果root的左子树和右子树有一个的话，说明，root就是最近公共祖先。有两个，接着往下走
	var count func(root *TreeNode, p, q int) int
	count = func(root *TreeNode, p, q int) int {
		if root != nil {
			if root.Val == p || root.Val == q {
				return 1 + count(root.Right, p, q) + count(root.Left, p, q)
			} else {
				return count(root.Right, p, q) + count(root.Left, p, q)
			}
		}
		return 0
	}
	left, right := count(root.Left, p.Val, q.Val), count(root.Right, p.Val, q.Val)
	for left != 1 && right != 1 {
		if left == 2 {
			root = root.Left
		} else {
			root = root.Right
		}
		left, right = count(root.Left, p.Val, q.Val), count(root.Right, p.Val, q.Val)
	}
	return root
}

//284. 顶端迭代器
type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return false
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	return 0
}

type PeekingIterator struct {
	Iterator *Iterator
	Peek     *int
}

func Constructor(iter *Iterator) *PeekingIterator {
	res := PeekingIterator{
		Iterator: iter,
		Peek:     nil,
	}
	return &res
}

func (this *PeekingIterator) hasNext() bool {
	if this.Peek != nil {
		return true
	}
	return this.Iterator.hasNext()
}

func (this *PeekingIterator) next() int {
	if this.Peek == nil {
		return this.Iterator.next()
	} else {
		res := *this.Peek
		this.Peek = nil
		return res
	}
}

func (this *PeekingIterator) peek() int {
	if this.Peek == nil {
		peekNum := this.Iterator.next()
		this.Peek = &peekNum
		return peekNum
	} else {
		return *this.Peek
	}
}

//每日一题：321. 拼接最大数
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	res := []int{}
	if k < 1 {
		return res
	}

	////所有数都要用到的时候
	//if k == len(nums1)+len(nums2) {
	//	//选择大的，当两个大小相同时，又要选择谁呢？递归
	//	if len(nums1) == 0 {
	//		return nums2
	//	} else if len(nums2) == 0 {
	//		return nums1
	//	}
	//
	//	if nums1[0] > nums2[0] {
	//		res = append(res, nums1[0])
	//		res = append(res, maxNumber(nums1[1:], nums2, k-1)...)
	//	} else if nums1[0] < nums2[0] {
	//		res = append(res, nums2[0])
	//		res = append(res, maxNumber(nums1, nums2[1:], k-1)...)
	//	} else {
	//		select1 := maxNumber(nums1[1:], nums2, k-1)
	//		select2 := maxNumber(nums1, nums2[1:], k-1)
	//		flag := true
	//		for i := 0; i < len(select1); i++ {
	//			if select2[i] > select1[i] {
	//				flag = false
	//				break
	//			} else if select1[i] > select2[i] {
	//				flag = true
	//				break
	//			}
	//		}
	//		res = append(res, nums1[0])
	//		if flag {
	//			res = append(res, select1...)
	//		} else {
	//			res = append(res, select2...)
	//		}
	//	}
	//	return res
	//}

	//k 小于 两数组大小之和
	index := len(nums1) + len(nums2) - k
	//从可以选的数中选中最大的
	max, maxIndex := math.MinInt32, 0
	for i := 0; i < len(nums1) && i <= index; i++ {
		if nums1[i] > max {
			max = nums1[i]
			maxIndex = i
		}
	}
	max2, max2Index := math.MinInt32, 0
	for i := 0; i < len(nums2) && i <= index; i++ {
		if nums2[i] > max2 {
			max2 = nums2[i]
			max2Index = i
		}
	}
	//比较max与max2的大小，选择大的，如果相等的话，
	if max > max2 {
		res = append(res, max)
		if maxIndex+1 <= len(nums1) {
			res = append(res, maxNumber(nums1[maxIndex+1:], nums2, k-1)...)
		} else {
			res = append(res, maxNumber([]int{}, nums2, k-1)...)
		}

	} else if max < max2 {
		res = append(res, max2)
		if max2Index+1 <= len(nums2) {
			res = append(res, maxNumber(nums1, nums2[max2Index+1:], k-1)...)
		} else {
			res = append(res, maxNumber(nums1, []int{}, k-1)...)
		}

	} else {
		//
		res = append(res, max)
		select1 := maxNumber(nums1[maxIndex+1:], nums2, k-1)
		select2 := maxNumber(nums1, nums2[max2Index+1:], k-1)
		flag := true
		for i := 0; i < len(select1); i++ {
			if select2[i] > select1[i] {
				flag = false
				break
			} else if select1[i] > select2[i] {
				flag = true
				break
			}
		}
		if flag {
			res = append(res, select1...)
		} else {
			res = append(res, select2...)
		}
	}

	return res
}
