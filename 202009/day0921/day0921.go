package main

import "fmt"

func main() {
	myList := Constructor()
	linkedList := &myList
	linkedList.AddAtTail(9)
	fmt.Println(linkedList.Get(0))
}

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//链表组件：

func numComponents(head *ListNode, G []int) int {
	m := make(map[int]bool, len(G))
	for _, v := range G {
		m[v] = true
	}
	//思路：从头开始走
	res := 0
	flag := false //组件的开头是否已经有了
	for ; head != nil; head = head.Next {
		if m[head.Val] {
			//如果存在
			flag = true
		} else {
			if flag {
				res++
			}
			flag = false
		}
	}
	if flag {
		res++
	}
	return res
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{10000, nil}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if this == nil || this.Val != 10000 {
		return -1
	}
	this = this.Next
	for ; index > 0 && this != nil; index-- {
		this = this.Next
	}
	if this == nil || index < 0 {
		return -1
	}
	return this.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	if this == nil || this.Val != 10000 {
		*this = Constructor()
		this.Next = &MyLinkedList{val, nil}
	} else {
		newNode := MyLinkedList{val, this.Next}
		this.Next = &newNode
	}
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this == nil || this.Val != 10000 {
		*this = Constructor()
	}

	newNode := MyLinkedList{val, nil}
	prev := this
	this = this.Next
	for ; this != nil; this = this.Next {
		prev = this
	}
	prev.Next = &newNode
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if this == nil || this.Val != 10000 {
		*this = Constructor()
	}
	if index <= 0 {
		this.AddAtHead(val)
		return
	}
	prev := this
	this = this.Next
	for ; index > 0 && this != nil; index-- {
		prev = this
		this = this.Next
	}
	if index > 0 {
		return
	}
	if index == 0 {
		newNode := MyLinkedList{val, this}
		prev.Next = &newNode
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this == nil || this.Val != 10000 {
		*this = Constructor()
	}
	prev := this
	this = this.Next

	for ; index > 0 && this != nil; index-- {
		prev = this
		this = this.Next
	}
	if this != nil {
		prev.Next = this.Next
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//给定一个二叉搜索树（Binary Search Tree），把它转换成为累加树（Greater Tree)，
//使得每个节点的值是原来的节点值加上所有大于它的节点值之和
//var temp int = 0
//问题：LeetCode测试的时候，测到下一个用例的时候，没有将该值清0

func convertBST(root *TreeNode) *TreeNode {
	//后序遍历
	var temp int = 0
	convertBST2(root, &temp)
	return root
}

func convertBST2(root *TreeNode, temp *int) *TreeNode {
	if root == nil {
		return root
	}
	convertBST2(root.Right, temp)
	root.Val += *temp
	*temp = root.Val
	convertBST2(root.Left, temp)

	return root
}
