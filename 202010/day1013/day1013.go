package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//24.两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	p2,p1,res := head,head.Next,head.Next
	var pre *ListNode = nil
	for ;p1 != nil ;{
		//交换p1与p2
		p1.Next,p2.Next = p2,p1.Next
		if pre != nil{
			pre.Next= p1
		}
		pre = p2
		p2 = p2.Next
		if p2 == nil{
			break
		}
		p1 = p2.Next
	}
	return res
}
