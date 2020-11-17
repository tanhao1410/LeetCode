package main

import "fmt"

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
}

//剑指 Offer 31. 栈的压入、弹出序列
func validateStackSequences(pushed []int, popped []int) bool {
	//准备好一个栈
	//先看pop序列的下一个，如果是在pushed的顶部，那么弹出即可，poped走一步，
	//如果不是，那么pushed里面的数字开始依次压入栈中，直到碰到poped的下一个数字。poped走一步
	//由于栈中有数字了，在看pushed顶部的时候，加入看栈顶，栈顶是的也可以
	stack := []int{}
	i, j := 0, 0
	for i < len(popped) && (j < len(popped) || len(stack) > 0) {
		if len(stack) == 0 {
			stack = append(stack, pushed[j])
			j++
		}
		for stack[len(stack)-1] != popped[i] && j < len(popped) {
			stack = append(stack, pushed[j])
			j++
		}
		if stack[len(stack)-1] == popped[i] {
			//pop出去一个
			stack = stack[:len(stack)-1]
			i++
		} else {
			return false
		}
	}
	return i == len(popped)
}
