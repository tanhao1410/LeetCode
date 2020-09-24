package main

import "fmt"

func main() {
	//nums := []int{4, 2, 1, 3}
	//list := createListByNums(nums)
	//
	//insertionSortList(list)
	//
	//fmt.Println(list)
	//fmt.Println(generateParenthesis(10))
	array := []string{"A", "1"}
	fmt.Println(findLongestSubarray(array))
}

func findLongestSubarray(array []string) []string {
	//先把array转化成 array[i] 的value 代表前面（包括自己）有多少字符
	dp := make([]int, len(array))
	for i := 0; i < len(array); i++ {
		if !(array[i][0] >= '0' && array[i][0] <= '9') {
			if i-1 < 0 {
				dp[i] = 1
			} else {
				dp[i] = dp[i-1] + 1
			}
		} else {
			if i-1 < 0 {
				dp[i] = 0
			} else {
				dp[i] = dp[i-1]
			}
		}
	}

	isNumCountEquChar := func(start, end int) bool {

		charCount := 0
		if start == 0 {
			charCount = dp[end]
		} else {
			charCount = dp[end] - dp[start-1]
		}

		numCount := end - start + 1 - charCount

		return charCount == numCount
	}

	maxSize, start := 0, 0
	for i := 0; i < len(array)-1; i++ {
		for j := i + maxSize + 1; j < len(array); j++ {
			if isNumCountEquChar(i, j) {
				if j-i > maxSize {
					maxSize = j - i + 1
					start = i
				}
			}
		}
		if maxSize > len(array)-i {
			break
		}
	}

	res := []string{}
	for i := start; i < start+maxSize; i++ {
		res = append(res, array[i])
	}

	return res
}

//括号生成
func generateParenthesis(n int) []string {
	res := &[]string{}
	m := make(map[string]bool)
	create("", n, n, res, true, m)
	return *res
}

func create(s string, l int, r int, res *[]string, canL bool, m map[string]bool) {

	if l == 0 && r == 0 {
		if !m[s] {
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
		create(s+ls, l-i, r, res, false, m)
	}

	// 可以补 ）
	for i := 1; i <= r-l; i++ {
		ls := ""
		for j := 0; j < i; j++ {
			ls += ")"
		}
		create(s+ls, l, r-i, res, true, m)
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
