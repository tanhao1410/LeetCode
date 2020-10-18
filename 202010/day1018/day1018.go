package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//每日一题：19.删除链表的倒数第N个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//双指针法：fast先走n步
	fast := head
	for ; n > 0; n-- {
		fast = fast.Next
	}
	var slow *ListNode
	for ; fast != nil; fast = fast.Next {
		if slow == nil {
			slow = head
		} else {
			slow = slow.Next
		}
	}

	if slow == nil {
		return head.Next
	}

	slow.Next = slow.Next.Next
	return head
}
