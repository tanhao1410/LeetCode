package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//每日一题：143.重排链表
func reorderList(head *ListNode) {
	//求总数
	count := 0
	for p := head; p != nil; p = p.Next {
		count++
	}
	if count <=2{
		return
	}
	//逆置后面的
	count = count / 2
	middle,pre := head,head
	for ; count >= 0; count-- {
		pre = middle
		middle = middle.Next
	}
	pre.Next = nil

	//逆置middle与end之间的链表
	if middle.Next != nil{

		m ,n:= middle,middle.Next
		for;n != nil;{
			n.Next,m,n = m,n,n.Next
		}
		middle.Next = nil
		middle = m
		//middle为新头
		for pp := head;pp != nil && middle != nil;{
			pp.Next,middle.Next ,middle,pp = middle,pp.Next,middle.Next,pp.Next
		}
	}else{
		//即middle 就一个数
		head.Next,middle.Next = middle,head.Next
	}
}
