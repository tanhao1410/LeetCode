package main

import "fmt"

func main() {
	fmt.Println(findNthDigit(3012))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//面试题 02.06. 回文链表
func isPalindrome(head *ListNode) bool {
	//简单思路：用一个栈通过进栈出栈即可做到时间复杂度o(n)
	//用 O(n) 时间复杂度和 O(1) 空间复杂度，逆置后半截链表
	size := 0
	for p := head; p != nil; p = p.Next {
		size++
	}

	if size <= 1 {
		return true
	}
	//从中间开始截取
	size = (size + 1) / 2
	p := head
	for ; size >= 0; p = p.Next {
		size--
	}

	//逆置链表
	pp, q := p, p.Next
	for q != nil {
		q.Next, pp, q = pp, q, q.Next
	}
	p.Next = nil

	for p != nil && head != nil {
		if p.Val != head.Val {
			return false
		}
		p, head = p.Next, head.Next
	}

	return true
}

//400. 第N个数字
func findNthDigit(n int) int {

	if n < 9 {
		return n
	}

	//求10^n
	tenN := func(n int) int {
		res := 1
		for i := 0; i < n; i++ {
			res *= 10
		}
		return res
	}

	//获取某个数的第n位，从后面数
	getNum := func(n int, index int) int {
		res := 1
		for i := 0; i <= index; i++ {
			res = n % 10
			n = n / 10
		}
		return res
	}

	for i, sum := 0, 0; ; i++ {

		nextSum := sum + 9*tenN(i)*(i+1)
		if n <= nextSum {
			//n - sum = 还剩多少个数没数，从1000.。。开始为1数起
			//n - sum / i + 1 = 应该数几个数。
			//目标数 = (n-sum)/(i+1) + tenN(i)

			//找到了该数
			targetNum := (n-sum)/(i+1) + tenN(i)
			index := (n - sum) % (i + 1)
			if index == 0 {
				return getNum(targetNum-1, 0)
			}

			return getNum(targetNum, i+1-index)
		}
		sum = nextSum
	}
}

//每日一题：399. 除法求值
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	res := []float64{}

	return res
}
