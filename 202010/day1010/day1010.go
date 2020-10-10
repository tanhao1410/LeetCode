package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//环形链表2
func detectCycle(head *ListNode) *ListNode {

	if head == nil{
		return head
	}

	for fast,slow := head,head;fast != nil;{
		slow,fast = slow.Next,fast.Next
		if fast == nil{
			return nil
		}
		fast = fast.Next

		//说明有环了，在此相遇
		if fast == slow{
			for point := head;point != slow;{
				point,slow = point.Next,slow.Next
			}
			return slow
		}
	}

	return nil
}
