package main

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//示例：
//
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807

func main() {

	//node1 := &ListNode{2,nil}
	//node2 := &ListNode{3,node1}
	//node3 := &ListNode{6,node2}
	//
	//
	//node4 := &ListNode{2,nil}
	//node5 := &ListNode{2,node4}
	//node6 := &ListNode{2,node2}
	//node7 := &ListNode{2,node2}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//想法，先求数，再相加，o(n1+n2)，一次遍历，

	//想法2：复用某一个链表，省得重新创建。复用l1的吧
	var res *ListNode = l1
	var tempL1 *ListNode = l1
	flag := 0 //表示是否有进位
	for l1 != nil && l2 != nil {
		l1.Val = l1.Val + l2.Val + flag
		if l1.Val > 9 {
			l1.Val -= 10
			flag = 1
		} else {
			flag = 0
		}
		tempL1 = l1
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 != nil {
		l1.Val += flag
	}
	//需要知道l1的最后一个地址。
	if l2 != nil {

		tempL1.Next = l2
		//l2比较长的话，还需要接着往前进位
		for flag == 1 || l2 == nil {
			l2.Val += flag
			if l2.Val > 9 {
				flag = 1
				l2.Val -= 10
			} else {
				flag = 0
			}

			l2 = l2.Next
		}

		//一直进位到最后，但是还缺少一个呢？？？

	}

	if l1 == nil && l2 == nil && flag == 1 {
		newNode := &ListNode{
			Val:  flag,
			Next: nil,
		}
		tempL1.Next = newNode
	}

	return res
}
