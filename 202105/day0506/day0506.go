package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Print(nthUglyNumber(12))
}

//剑指 Offer 49. 丑数
func nthUglyNumber(n int) int {
	//把所有的丑数全都求出来，
	nums := []int{}

	getN := func(base, n int) int {
		res := 1
		for i := 0; i < n; i++ {
			res *= base
		}
		return res
	}

	for i := 0; i < 32; i++ {
		for j := 0; j < 32; j++ {
			for k := 0; k < 16; k++ {
				num := getN(2, i) * getN(3, j) * getN(5, k)
				if num > 0 {
					nums = append(nums, num)
				} else {
					break
				}
			}
		}
	}

	sort.Ints(nums)

	return nums[n-1]
}

//1720. 解码异或后的数组
func decode(encoded []int, first int) []int {
	res := make([]int, len(encoded)+1)
	res[0] = first
	for i := 1; i <= len(encoded); i++ {
		res[i] = res[i-1] ^ encoded[i-1]
	}
	return res
}

//剑指 Offer 59 - II. 队列的最大值
type Node struct {
	Val  int
	Next *Node
}

type MaxQueue struct {
	MaxNode *Node
	Head    *Node
	Tail    *Node
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if this.MaxNode == nil {
		return -1
	}
	return this.MaxNode.Val
}

func (this *MaxQueue) Push_back(value int) {
	newNode := &Node{
		Val: value,
	}
	//没有节点的时候
	if this.Tail == nil {
		this.Tail = newNode
		this.Head = this.Tail
		this.MaxNode = this.Tail
		return
	}
	//有节点，新节点插入到尾后面，更新尾
	this.Tail.Next = newNode
	this.Tail = newNode

	//如果比最大值大，更新最大值节点
	if value >= this.MaxNode.Val {
		this.MaxNode = this.Tail
	}

	return
}

func (this *MaxQueue) Pop_front() int {

	if this.Head == nil {
		return -1
	}
	res := this.Head.Val

	//如果就一个，直接返回即可
	if this.Head == this.Tail {
		this.Head = nil
		this.Tail = nil
		this.MaxNode = nil
		return res
	} else if this.Head == this.MaxNode {
		//寻找新的最大值
		maxNode := this.Head.Next
		for p := maxNode.Next; p != nil; p = p.Next {
			if p.Val > maxNode.Val {
				maxNode = p
			}
		}
		this.MaxNode = maxNode
	}

	this.Head = this.Head.Next
	return res
}
