package main

import "fmt"

func main() {
	//fmt.Println(reverseWords("  hello   world i s ha   "))
	fmt.Println(maxProduct([]int{0, 0, 0, 1, 1, 1, 2, 3, -2, 3, 3, -1, -9, -2, -2, 0, 0, 0, -2, -2, -2, -2}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}
//160.相交链表
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	if headB == nil || headA == nil{
		return nil
	}
	zeroFlag := 0
	pa,pb := headA,headB
	for pa != pb{
		if pa.Next == nil{
			pa = headB
			zeroFlag++
			if zeroFlag > 1{
				return nil
			}
		}else{
			pa = pa.Next
		}
		if pb.Next == nil{
			pb = headA
		}else{
			pb = pb.Next
		}
	}
	return pa
}

//152.乘积最大子数组
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	i := 0
	for ; i < len(nums) && nums[i] == 0; i++ {
	}
	//以0分割，直接*，遇到变成负值了，记录，当前最大的，接着×，

	x := nums[i]
	y := nums[i]
	i++
	for ; i < len(nums); i++ {
		if nums[i] == 0 {
			if res < 0 {
				res = 0
			}
			for ; i < len(nums) && nums[i] == 0; i++ {
			}
			if i == len(nums) {
				break
			}
			x = nums[i]
			y = nums[i]
			i++
		} else if nums[i] < 0 {
			if y < 0 {
				x = y * nums[i]
				y = x
			} else {
				x = 1
				y *= nums[i]
			}
		} else {
			x *= nums[i]
			y *= nums[i]
		}

		if x > res {
			res = x
		}

	}
	return res
}

//151.翻转字符串里的单词
func reverseWords(s string) string {
	res := ""
	for i, j := len(s)-1, len(s)-1; j >= 0; {
		if s[i] == ' ' {
			i--
			j = i
			continue
		}
		if s[j] == ' ' {
			//说明可以生成一个单词了
			word := s[j+1 : i+1]
			res = res + word + " "
			i = j - 1
			j = i
			continue
		}
		if j == 0 {
			word := s[:i+1]
			return res + word
		}
		j--
	}
	if len(res) > 0 {
		return res[:len(res)-1]
	}
	return res

}

func minDepth2(root *TreeNode) int {
	//思路：
	if root == nil {
		return 0
	}
	left := minDepth2(root.Left)
	right := minDepth2(root.Right)

	if left == 0 || (left > right && right != 0) {
		return right + 1
	}
	return left + 1
}

//111.二叉树的最小深度
func minDepth(root *TreeNode) int {
	//思路：求左边，求右边，然后返回最小的。效率有点低。一旦超过了，就应该不继续往下求了。
	m := make(map[*TreeNode]int, 0)
	var getMinDepth func(root *TreeNode) int
	getMinDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		if root.Right == nil && root.Left == nil {
			return 1
		}
		var left int
		var right int
		if v, ok := m[root.Left]; ok {
			left = v
		} else {
			left = getMinDepth(root.Left)
		}
		if v, ok := m[root.Right]; ok {
			right = v
		} else {
			right = getMinDepth(root.Right)
		}

		if left == 0 || (left > right && right != 0) {
			m[root] = right + 1
			return right + 1
		} else {
			m[root] = left + 1
			return left + 1
		}
	}
	return getMinDepth(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//每日一题：144.二叉树的前序遍历
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}
	stackLen := len(stack)
	for stackLen > 0 {
		//出栈
		node := stack[stackLen-1]
		stack = stack[:stackLen-1]
		res = append(res, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		stackLen = len(stack)
	}
	return res
}
