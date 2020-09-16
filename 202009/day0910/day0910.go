package main

import "fmt"

func main() {
	//fmt.Println(isUnique("letcod天天"))
	a := [2]int{5, 6}
	b := [2]int{5, 6}
	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}

	//if a[:] == b[:] {
	//	fmt.Println("equal")
	//} else {
	//	fmt.Println("not equal")
	//}
}

//判断所有字符都不相同--不使用额外的数据结构,0 <= len(s) <= 100
func isUnique(astr string) bool {
	//第一个思路，用map
	//第二，长度不是很长，
	nums := make([]byte, 128)
	for i := 0; i < len(astr); i++ {
		if nums[int(astr[i])] == 1 {
			return false
		}
		nums[int(astr[i])] = 1
	}
	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//要考虑0的情况
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//双复制法，反转后的
	ll1 := &ListNode{l1.Val, nil}
	ll2 := &ListNode{l2.Val, nil}

	l1len, l2len := 1, 1
	for l1 = l1.Next; l1 != nil; l1 = l1.Next {
		//前插法
		node := &ListNode{l1.Val, ll1}
		ll1 = node
		l1len++
	}

	for l2 = l2.Next; l2 != nil; l2 = l2.Next {
		//前插法
		node := &ListNode{l2.Val, ll2}
		ll2 = node
		l2len++
	}

	//l1为最长的，以l1为基准返回
	if l1len < l2len {
		ll1, ll2 = ll2, ll1
	}
	head := ll1

	//开始相加
	flag := 0
	for ; ll2 != nil; ll2 = ll2.Next {
		ll1.Val = ll1.Val + ll2.Val + flag
		if ll1.Val > 9 {
			ll1.Val %= 10
			flag = 1
		} else {
			flag = 0
		}
		ll1 = ll1.Next
	}

	for ; ll1 != nil; ll1 = ll1.Next {
		ll1.Val = ll1.Val + flag
		if ll1.Val > 9 {
			ll1.Val %= 10
			flag = 1
		} else {
			break
		}
	}

	//逆置结果返回即可
	var pre *ListNode = nil
	for head != nil {
		pre, head.Next, head = head, pre, head.Next
	}

	if flag == 1 {
		//最后一位有进位
		node := &ListNode{1, pre}
		pre = node
	}
	return pre
}

//输入链表不能修改 都不会以0开头，非空非负 (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4) --->7 -> 8 -> 0 -> 7
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode = nil
	//转成整数，相加，再生成链表=======>这种方式会越界的。不行
	num1 := 0
	for ; l1 != nil; l1 = l1.Next {
		num1 = num1*10 + l1.Val
	}
	num2 := 0
	for ; l2 != nil; l2 = l2.Next {
		num2 = num2*10 + l2.Val
	}
	sum := num1 + num2
	for ; sum != 0; sum = sum / 10 {
		val := sum % 10
		node := &ListNode{
			Val:  val,
			Next: res,
		}
		res = node
	}

	return res
}

//
