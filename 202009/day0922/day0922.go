package main

import (
	"fmt"
	"github.com/axgle/mahonia"
	"math"
)

func main() {
	//fmt.Println(detectCycle(createCycleListByNums([]int{1, 2})).Val)
	//comm := exec.Command("cmd.exe", "/C", "dir")
	//output, err := comm.Output()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(ConvertToString(output, "uft-8"))
	//comm.Run()
	node2 := &TreeNode{0, nil, nil}
	node1 := &TreeNode{0, node2, nil}
	root := &TreeNode{0, node1, nil}

	fmt.Println(minCameraCover(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minCameraCover(root *TreeNode) int {

	if root == nil {
		return 0
	}

	//问题？如果它的父节点已经有了的，这里返回1就不合适了
	if root.Left == nil && root.Right == nil {
		return 1
	}

	notInRootCount := notInRoot(root)
	//放在root上
	inRootCount := inRoot(root)

	if notInRootCount > inRootCount {
		return inRootCount
	} else {
		return notInRootCount
	}
}

func parentHave(root *TreeNode) int {
	//它的父节点有了，代表它可以没有

	if root == nil {
		return 0
	}
	//没有子节点了，直接返回0
	if root.Right == nil && root.Left == nil {
		return 0
	}

	count1 := minCameraCover(root)
	count2 := minCameraCover(root.Left) + minCameraCover(root.Right)
	if count1 > count2 {
		return count2
	}
	return count1
}

func inRoot(root *TreeNode) int {

	return 1 + parentHave(root.Right) + parentHave(root.Left)
}

//左孩子或右孩子必须放一个，或两个
func notInRoot(root *TreeNode) int {

	c1 := math.MaxInt32
	c2 := math.MaxInt32

	//左右节点可能有些为空，必须要考虑的
	if root.Left == nil {
		c2 = inRoot(root.Right) + minCameraCover(root.Left)
	}
	if root.Right == nil {
		c1 = inRoot(root.Left) + minCameraCover(root.Right)
	}

	//都不等于nil的情况没有考虑
	if root.Left != nil && root.Right != nil {
		c2 = inRoot(root.Right) + minCameraCover(root.Left)
		c1 = inRoot(root.Left) + minCameraCover(root.Right)
	}

	//然后才是返回小的
	if c1 > c2 {
		return c2
	}

	return c1
}

func ConvertToString(src []byte, tagCode string) string {
	enc := mahonia.NewDecoder("gbk")
	str := enc.ConvertString(string(src))
	return str
}

func createCycleListByNums(nums []int) *ListNode {

	var head *ListNode = nil
	temp := head
	for i := len(nums) - 1; i >= 0; i-- {
		node := &ListNode{nums[i], head}
		if i == len(nums)-1 {
			temp = node
		}
		head = node
	}
	temp.Next = head

	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//环行链表返回头
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	//1.快慢指针确定环
	for fast, slow := head, head; ; {

		fast, slow = fast.Next, slow.Next
		if fast != nil {
			fast = fast.Next
		} else {
			return nil
		}

		//赶上了，说明有环了
		if fast == slow {
			for fast = head; ; {
				if fast == slow {
					return slow
				}
				fast = fast.Next
				slow = slow.Next
			}
		}
	}
	//2.
	return nil
}
