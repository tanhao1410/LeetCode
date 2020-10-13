package main

func main() {

}

//74.搜索二维矩阵
func searchMatrix(matrix [][]int, target int) bool {
	//直接二分
	m := len(matrix)
	if m == 0{
		return false
	}
	n := len(matrix[0])
	if n == 0{
		return false
	}

	start,end := 0,m*n-1
	middle := (start+end)/2
	for start <= end{
		middleValue := matrix[middle/n][middle%n]
		if middleValue == target{
			return true
		}else if middleValue > target{
			end = middle -1
		}else{
			start = middle +1
		}
		middle = (start+end)/2
	}

	return false
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
