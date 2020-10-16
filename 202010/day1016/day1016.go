package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 3}
	fmt.Println(subsetsWithDup(nums))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//102.二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil{
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) != 0{

		item:= []int{}
		size := len(queue)
		for i:=0;i < size;i++{
			item = append(item, queue[i].Val)
			if queue[i].Left != nil{
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil{
				queue = append(queue,queue[i].Right)
			}
		}
		res = append(res, item)
		queue = queue[size:]
	}

	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//203.移除链表元素
func removeElements(head *ListNode, val int) *ListNode {

	var res *ListNode
	for ; head != nil && head.Val == val; head = head.Next {
	}
	res = head
	pre := res
	for p := res; p != nil; p = p.Next {
		if p.Val == val {
			pre.Next = p.Next
		} else {
			pre = p
		}
	}
	return res
}

//90.子集 II
func subsetsWithDup(nums []int) [][]int {
	//子集可能包含重复元素
	//先排序
	//只能选择比自己大的，所有的都按照有序来选择
	res := [][]int{[]int{}}
	sortNums(nums)
	nextNum([]int{}, nums, &res)
	return res
}

func nextNum(have, remain []int, res *[][]int) {
	if len(remain) == 0 {
		return
	}
	for i := 0; i < len(remain); i++ {
		if i != 0 && remain[i] == remain[i-1] {
			continue
		}
		if len(have) != 0 && remain[i] < have[len(have)-1] {
			continue
		}
		have2 := make([]int, len(have))
		copy(have2, have)
		have2 = append(have2, remain[i])
		*res = append(*res, have2)
		var remain2 []int
		if i == 0 {
			remain2 = remain[1:]
		} else if i == len(remain)-1 {
			remain2 = remain[:i]
		} else {
			remain2 = append(remain[:i], remain[i+1:]...)
		}
		nextNum(have2, remain2, res)
	}
}

func sortNums(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

//977.有序数组的平方
func sortedSquares(A []int) []int {
	res := make([]int, len(A))
	for start, end, i := 0, len(A)-1, len(A)-1; start <= end; i-- {
		start2 := A[start] * A[start]
		end2 := A[end] * A[end]
		if start2 > end2 {
			res[i] = start2
			start++
		} else {
			res[i] = end2
			end--
		}
	}
	return res
}
