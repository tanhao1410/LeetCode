package main

import "fmt"

func main() {
	fmt.Println(detectCycle(createCycleListByNums([]int{1, 2})).Val)
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
