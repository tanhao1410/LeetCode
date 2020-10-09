package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//判断是否是环行链表
func hasCycle(head *ListNode) bool {
	//双指针
	if head == nil{
		return false
	}

	for slow,fast := head,head.Next;fast != nil;{
		if slow == fast {
			return true
		}
		if fast.Next == nil  || fast.Next.Next == nil{
			return false
		}
		slow,fast = slow.Next,fast.Next.Next
	}

	return false
}
