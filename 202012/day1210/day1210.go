package main

import "fmt"

func main() {
	fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 5, 5, 10, 20, 20, 20}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//23. 合并K个升序链表
func mergeKLists(lists []*ListNode) *ListNode {
	//遍历所有链表的头，找到最小的，放到返回结果中。如果某个链表结束了，就不用遍历了，可以用一个数组来记录，每个链表应该访问的位置
	var head *ListNode
	var p *ListNode
	for {
		minKey := -1
		for k, v := range lists {
			if v != nil && (minKey == -1 || v.Val < lists[minKey].Val) {
				minKey = k
			}
		}
		if minKey == -1 {
			return head
		}
		if head == nil {
			head, lists[minKey] = lists[minKey], lists[minKey].Next
			p = head
		} else {
			p.Next, lists[minKey] = lists[minKey], lists[minKey].Next
			p = p.Next
		}
	}
}

//每日一题：860. 柠檬水找零
func lemonadeChange(bills []int) bool {
	//用一个map记录5,10,20的个数。如果收到了5，直接m[5]+1，如果收到10,m[10=+1,5--,如果收到20，则10--。5++
	five, ten := 0, 0
	for _, v := range bills {
		switch v {
		case 5:
			five++
		case 10:
			ten++
			five--
			if five < 0 {
				return false
			}
		case 20:
			if ten > 0 {
				ten--
				five--
			} else {
				five -= 3
			}

			if five < 0 || ten < 0 {
				return false
			}
		}
	}

	return true
}
