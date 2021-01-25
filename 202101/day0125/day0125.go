package main

func main() {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	queue.Peek()  // 返回 1
	queue.Pop()   // 返回 1
	queue.Empty() // 返回 false
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
