package main

import "sort"

func main() {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	queue.Peek()  // 返回 1
	queue.Pop()   // 返回 1
	queue.Empty() // 返回 false
}

//面试题 16.06. 最小差
func smallestDifference(a []int, b []int) int {
	//先排序,两个指针，谁小，谁往后走，并更新最小值
	sort.Ints(a)
	sort.Ints(b)
	min := 2147483647
	for i, j := 0, 0; i < len(a) && j < len(b); {
		//遇到相等的了，直接返回即可
		if a[i] < b[j] {
			if b[j]-a[i] < min {
				min = b[j] - a[i]
			}
			i++
		} else if a[i] > b[j] {
			if a[i]-b[j] < min {
				min = a[i] - b[j]
			}
			j++
		} else {
			return 0
		}
	}
	return min
}

//面试题 03.04. 化栈为队
type MyStack struct {
	nums []int
}

func (this *MyStack) push(x int) {
	this.nums = append(this.nums, x)
}
func (this *MyStack) pop() int {
	if this.isEmpty() {
		return -1
	}
	res := this.nums[len(this.nums)-1]
	this.nums = this.nums[:len(this.nums)-1]
	return res
}
func (this *MyStack) peek() int {
	if this.isEmpty() {
		return -1
	}
	res := this.nums[len(this.nums)-1]
	return res
}
func (this *MyStack) size() int {
	return len(this.nums)
}
func (this *MyStack) isEmpty() bool {
	return this.size() == 0
}

type MyQueue struct {
	stack1 *MyStack
	stack2 *MyStack
}

func Constructor() MyQueue {
	return MyQueue{
		stack1: &MyStack{},
		stack2: &MyStack{},
	}
}

func (this *MyQueue) Push(x int) {
	this.stack1.push(x)
}

func (this *MyQueue) Pop() int {
	if this.stack2.isEmpty() {
		for !this.stack1.isEmpty() {
			ele := this.stack1.pop()
			this.stack2.push(ele)
		}
	}
	return this.stack2.pop()
}

func (this *MyQueue) Peek() int {
	if this.stack2.isEmpty() {
		for !this.stack1.isEmpty() {
			ele := this.stack1.pop()
			this.stack2.push(ele)
		}
	}
	return this.stack2.peek()
}

func (this *MyQueue) Empty() bool {
	return this.stack2.isEmpty() && this.stack1.isEmpty()
}

//面试题 08.05. 递归乘法
func multiply(A int, B int) int {
	if A == 1 {
		return B
	}
	if B == 1 {
		return A
	}
	if A > B {
		return multiply(A, B-1) + A
	}
	return multiply(A-1, B) + B
}

//每日一题：959. 由斜杠划分区域
func regionsBySlashes(grid []string) int {

	return 0
}
