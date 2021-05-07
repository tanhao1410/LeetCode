package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//1721. 交换链表中的节点
func swapNodes(head *ListNode, k int) *ListNode {

	p := head
	for n := 1; n < k; n++ {
		p = p.Next
	}
	start := p

	end := head
	for ; p.Next != nil; p, end = p.Next, end.Next {
	}

	start.Val, end.Val = end.Val, start.Val

	// for p,n:=head,1;p != nil;p ,n = p.Next,n + 1{
	//     if n == k{
	//         p.Val,end.Val = end.Val,p.Val
	//     }
	// }

	return head
}
