package main

import (
	"math/rand"
)

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//面试题 02.01. 移除重复节点
func removeDuplicateNodes(head *ListNode) *ListNode {
	//思路：采用map记录重复的数
	m := make(map[int]bool)
	var res *ListNode = &ListNode{
		0, nil,
	}
	for resP, p := res, head; p != nil; {
		if !m[p.Val] {
			resP.Next = p
			resP = p
			p = p.Next
			resP.Next = nil
			m[p.Val] = true
		} else {
			p = p.Next
		}
	}
	return res.Next
}

//519. 随机翻转矩阵
type Solution struct {
	zero []int
	one  []int
}

func Constructor(n_rows int, n_cols int) Solution {
	zero := make([]int, n_rows*n_cols)
	for i := 0; i < n_rows; i++ {
		for j := 0; j < n_cols; j++ {
			zero[i*n_cols+j] = 10000*i + j
		}
	}
	one := make([]int, 0)
	return Solution{
		zero: zero,
		one:  one,
	}
}

func (this *Solution) Flip() []int {
	//均匀随机的将矩阵中的 0 变为 1，并返回该值的位置下标，如何在剩下的里面进行随机呢。
	if len(this.zero) == 0 {
		return nil
	}
	index := rand.Intn(len(this.zero))
	res := []int{this.zero[index] / 10000, this.zero[index] % 10000}
	//从zero中删去该数，加入到one中
	this.one = append(this.one, index)
	this.zero[index] = this.zero[len(this.zero)-1]
	this.zero = this.zero[:len(this.zero)-1]
	return res
}

func (this *Solution) Reset() {
	this.zero = append(this.zero, this.one...)
	this.one = make([]int, 0)
}
