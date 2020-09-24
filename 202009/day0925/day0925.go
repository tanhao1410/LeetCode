package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//分隔链表
func partition(head *ListNode, x int) *ListNode {
	//思路:一个指针往前走，停在比x大的第一个位置的前面。
	//第二个指针往后走，遇到比x小的就插入到前面指针的后面即可
	if head == nil {
		return head
	}

	var first *ListNode // 指向第一个不小于x的前一个位置
	p := head           //指向第一个不小于x的节点
	pPre := head        // 指向p的前一个节点
	for ; p.Val < x && p != nil; p = p.Next {
		pPre = p
		first = p
	}

	for p != nil {
		if p.Val >= x {
			pPre = p
			p = p.Next
		} else {
			if first == nil {
				//first == nil ,
				p, first, pPre.Next = p.Next, p, p.Next
				first.Next = head
				head = first
			} else {
				p, pPre.Next, first.Next, p.Next = p.Next, p.Next, p, first.Next
				first = first.Next
			}

		}
	}

	return head

}

func removeZeroSumSublists(head *ListNode) *ListNode {

	return nil
}
