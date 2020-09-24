package main

import "fmt"

func main() {
	nums := []int{4, 2, 1, 3}
	list := createListByNums(nums)

	insertionSortList(list)

	fmt.Println(list)
}

func createListByNums(nums []int) *ListNode {

	var head *ListNode = nil
	for i := len(nums) - 1; i >= 0; i-- {
		node := &ListNode{nums[i], head}
		head = node
	}
	return head
}

//数组中重复的元素，不用额外的空间，时间复杂度o(n)
func findDuplicates(nums []int) []int {
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//对链表进行插入排序
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	for pPre, p := head, head.Next; p != nil; {

		//走到比p大的位置
		pp, ppPre := head, head
		for ; pp != nil && pp.Val < p.Val && pp != p; pp = pp.Next {
			ppPre = pp
		}

		//开头就比它大，所以应该放在最前面
		if pp == head {
			p.Next, p, head = pp, p.Next, p
			pPre.Next = p
		} else if pp == p {
			p, pPre = p.Next, p
		} else {
			ppPre.Next, p.Next, p = p, pp, p.Next
			pPre.Next = p
		}

	}

	return head
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//不用额外空间的话怎么解决呢，不算递归话费的空间
func findMode(root *TreeNode) []int {
	//只有左子树可能与根相等。最多的？先递归一次找最大的数量，然后？
	return nil
}

//二叉搜索树中的众数
func findMode2(root *TreeNode) []int {
	//思路：右边的肯定比根大，只有左可能相等
	//用一个map记录各数出现的频率。
	m := make(map[int]int)
	readTree(root, m)
	maxCount := 0
	for _, count := range m {
		if count > maxCount {
			maxCount = count
		}
	}

	res := []int{}
	for k, v := range m {
		if v == maxCount {
			res = append(res, k)
		}
	}
	return res
}

func readTree(root *TreeNode, m map[int]int) {
	if root == nil {
		return
	}
	if _, ok := m[root.Val]; ok {
		m[root.Val]++
	} else {
		m[root.Val] = 1
	}
	readTree(root.Left, m)
	readTree(root.Right, m)
}
