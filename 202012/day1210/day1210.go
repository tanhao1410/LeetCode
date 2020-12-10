package main

import "fmt"

func main() {
	fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 5, 5, 10, 20, 20, 20}))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//99. 恢复二叉搜索树
func recoverTree(root *TreeNode) {
	//中序遍历应该是有序的
	//思路：先找到不合适的地方。
	//关键是要和谁交换。如果能找到两个不合适的节点，他们俩交换。如果只找到一个的话，
	//方法1：中序遍历后，找到两个位置交换变有序。即为结果。
	nums := []int{}

	var readTree func(node *TreeNode)
	readTree = func(root *TreeNode) {
		if root != nil {
			readTree(root.Left)
			nums = append(nums, root.Val)
			readTree(root.Right)
		}
	}
	readTree(root)

	//冒泡法排序
	for flag := true; flag; {
		flag = false
		for i := 0; i < len(nums)-1; i++ {
			//后一个比前一个要小
			if nums[i+1] < nums[i] {
				//交换
				flag = true
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
	//写树
	var writeTree func(node *TreeNode)
	writeTree = func(root *TreeNode) {
		if root != nil {
			writeTree(root.Left)
			root.Val = nums[0]
			nums = nums[1:]
			writeTree(root.Right)
		}
	}
	writeTree(root)
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
