package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	//fmt.Println(permuteUnique(nums))
	head := createListByNums(nums)

	//reorderList(head)
	//findMiddleNode(head)
	tree := sortedListToBST(head)
	fmt.Println(tree)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//有序表转化成二叉搜索树
func sortedListToBST(head *ListNode) *TreeNode {
	//想法：找到中间的数，作为根节点。切割成两半。左边，右边同样的流程进行递归操作
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{head.Val, nil, nil}
	}
	middle := findMiddleNode(head)
	root := &TreeNode{middle.Val, sortedListToBST(head), sortedListToBST(middle.Next)}
	return root
}

//找list的中间节点，并切割
func findMiddleNode(head *ListNode) *ListNode {
	//快慢指针
	//一个节点的时候不用切割
	if head == nil || head.Next == nil {
		return nil
	}
	first, second, pre := head, head, head
	for first != nil {
		first = first.Next
		if first == nil {
			pre.Next = nil
			return second
		}
		pre = second
		first, second = first.Next, second.Next
	}
	pre.Next = nil
	return second
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func createListByNums(nums []int) *ListNode {

	var head *ListNode = nil
	for i := len(nums) - 1; i >= 0; i-- {
		node := &ListNode{nums[i], head}
		head = node
	}
	return head
}

//重排链表，不能只交换值。
//L0→Ln→L1→Ln-1→L2→Ln-2→…
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	//思路：第一个后面是最后一个，第二个后面是倒数第二个。。。。
	//先计算总数，把后面的几个拿出来，保存起来，然后拼接
	nodes := []*ListNode{}
	//求长度
	listSize := 0
	for p := head; p != nil; p = p.Next {
		listSize++
	}
	//从len/2 -1的位置开始截取，截取的第一个的前一个，应该要指向nil
	start := (listSize+1)/2 - 1
	prev := head
	for p := head; p != nil; p = p.Next {

		if start >= listSize {
			nodes = append(nodes, p)
		}
		if start == listSize-1 {
			prev = p
		}
		start++
	}
	prev.Next = nil

	//开始拼接
	for p, i := head, len(nodes)-1; i >= 0; p = p.Next.Next {
		p.Next, nodes[i].Next = nodes[i], p.Next
		i--
	}

	fmt.Println("...")
}

//全排列2
//给定一个可包含重复数字的序列，返回所有不重复的全排列
func permuteUnique(nums []int) [][]int {
	//思路：先排序，再递归
	//先选择一个数，再在剩下的里面选一个。依次循环
	res := &[][]int{}
	slice := sort.IntSlice(nums)
	slice.Sort()
	permute2(nums, []int{}, res)
	return *res
}

func permute2(nums []int, parts []int, res *[][]int) {
	//所有的数字都选完了，即递归结束
	if 0 == len(nums) {
		*res = append(*res, parts)
		return
	}
	//从剩下的数字中选择，剩下的数字为方法传递过来的数，不选择相同的数
	prev := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] == prev {
			continue
		}
		nums2 := []int{}
		for j := 0; j < len(nums); j++ {
			if j != i {
				nums2 = append(nums2, nums[j])
			}
		}
		part2 := []int{}
		for _, v := range parts {
			part2 = append(part2, v)
		}
		part2 = append(part2, nums[i])
		prev = nums[i]
		permute2(nums2, part2, res)
	}
}
