package main

import "fmt"

func main() {
	//nums := []int{4, 2, 1, 3}
	//list := createListByNums(nums)
	//
	//insertionSortList(list)
	//
	//fmt.Println(list)
	fmt.Println(generateParenthesis(10))
}

//括号生成
func generateParenthesis(n int) []string {
	res := &[]string{}
	m := make(map[string]bool)
	create("", n, n, res,true,m)
	return *res
}

func create(s string, l int, r int, res *[]string, canL bool,m map[string]bool) {

	if l == 0 && r == 0 {
		if !m[s]{
			*res = append(*res, s)
			m[s] = true
		}
	}

	//可以补（，
	for i := 1; i <= l && canL; i++ {
		ls := ""
		for j := 0; j < i; j++ {
			ls += "("
		}
		create(s+ls, l-i, r, res,false,m)
	}

	// 可以补 ）
	for i := 1; i <= r-l; i++ {
		ls := ""
		for j := 0; j < i; j++ {
			ls += ")"
		}
		create(s+ls, l, r-i, res,true,m)
	}

}

//跳跃游戏
func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	//思路：从后往前看，该位置能否到达末尾，
	nums[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] == 0 {
			continue
		}

		if i+nums[i] >= len(nums)-1 {
			nums[i] = 1
			continue
		}

		//如果说它后面能到达的都是0,那么它也是0
		flag := false
		for j := 1; j <= nums[i]; j++ {
			if nums[i+j] != 0 {
				flag = true
				break
			}
		}
		if flag {
			nums[i] = 1
		} else {
			nums[i] = 0
		}
	}
	return nums[0] != 0
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
