package main

import (
	"fmt"
	"sort"
)

func main() {
	queue := Constructor()
	queue.PushFront(1)
	queue.PushBack(2)
	queue.PushMiddle(3)
	queue.PushMiddle(4)
	fmt.Println(queue.PopFront())
	fmt.Println(queue.PopMiddle())
	fmt.Println(queue.PopMiddle())
	fmt.Println(queue.PopBack())
	fmt.Println(queue.PopFront())
}

//1670. 设计前中后队列
type Node struct {
	Val  int
	Pre  *Node
	Next *Node
}
type FrontMiddleBackQueue struct {
	Size   int
	Head   *Node
	Tail   *Node
	Middle *Node
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	node := &Node{
		Val: val,
	}
	//队列如果为空，
	if this.Size == 0 {
		this.Head = node
		this.Tail = node
		this.Middle = node
	} else {

		//更新头节点
		node.Next = this.Head
		this.Head.Pre = node

		this.Head = node

		//更新中间节点,size 是奇数，前面插入一个数后，middle需要往前挪，否则不用动
		if this.Size%2 == 1 {
			this.Middle = this.Middle.Pre
		}

	}
	this.Size += 1
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	node := &Node{
		Val: val,
	}
	//队列如果为空，
	if this.Size == 0 {
		this.Head = node
		this.Tail = node
		this.Middle = node
	} else if this.Size == 1 {
		//直接插入到前面即可
		this.PushFront(val)
		this.Size -= 1
	} else {

		//如果size是奇数，插入到middle前面，否则插入到middle后面
		if this.Size%2 == 0 {

			node.Next = this.Middle.Next
			node.Pre = this.Middle
			this.Middle.Next.Pre = node
			this.Middle.Next = node
		} else {

			node.Next = this.Middle
			node.Pre = this.Middle.Pre
			this.Middle.Pre.Next = node
			this.Middle.Pre = node
		}

		//更新middle
		this.Middle = node

	}
	this.Size += 1
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	node := &Node{
		Val: val,
	}
	//队列如果为空，
	if this.Size == 0 {
		this.Head = node
		this.Tail = node
		this.Middle = node
	} else {

		//更新尾节点
		node.Pre = this.Tail
		this.Tail.Next = node

		this.Tail = node

		//更新中间节点,size 是奇数，前面插入一个数后，middle需要往前挪，否则不用动
		if this.Size%2 == 0 {
			this.Middle = this.Middle.Next
		}

	}
	this.Size += 1
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if this.Size == 0 {
		return -1
	}
	res := this.Head.Val

	if this.Size == 1 {
		this.Head = nil
		this.Middle = nil
		this.Tail = nil
	} else {
		this.Head = this.Head.Next
		this.Head.Pre = nil

		//更新middle
		if this.Size%2 == 0 {
			this.Middle = this.Middle.Next
		}
	}
	this.Size -= 1

	return res
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if this.Size == 0 {
		return -1
	}
	res := this.Middle.Val

	if this.Size == 1 {
		this.Head = nil
		this.Middle = nil
		this.Tail = nil
	} else if this.Size == 2 {
		return this.PopFront()
	} else {

		if this.Size%2 == 0 {

			this.Middle.Pre.Next = this.Middle.Next
			this.Middle.Next.Pre = this.Middle.Pre

			this.Middle = this.Middle.Next
		} else {

			this.Middle.Pre.Next = this.Middle.Next
			this.Middle.Next.Pre = this.Middle.Pre

			this.Middle = this.Middle.Pre

		}

	}
	this.Size -= 1

	return res
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if this.Size == 0 {
		return -1
	}
	res := this.Tail.Val

	if this.Size == 1 {
		this.Head = nil
		this.Middle = nil
		this.Tail = nil
	} else {
		this.Tail = this.Tail.Pre
		this.Tail.Next = nil

		//更新middle
		if this.Size%2 == 1 {
			this.Middle = this.Middle.Pre
		}
	}
	this.Size -= 1

	return res
}

//287. 寻找重复数
func findDuplicate(nums []int) int {
	//思路：如何判断存在一个重复的，排序后，最后一个数一定是小于1+  n，而不是n+1
	sort.Ints(nums)
	start, end := 0, len(nums)-1
	for middle := (start + end) / 2; ; middle = (start + end) / 2 {
		if end > start && nums[end] == nums[start] {
			return nums[end]
		}
		//重复值在后面
		if middle-start == nums[middle]-nums[start] || end-middle > nums[end]-nums[middle] {
			start = middle
		} else if end-middle == nums[end]-nums[middle] || middle-start > nums[middle]-nums[start] {
			end = middle
		}

	}
}
