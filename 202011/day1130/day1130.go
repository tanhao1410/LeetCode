package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(reorganizeString("ababa"))
	fmt.Println(rangeBitwiseAnd(0, 1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//148. 排序链表
func sortList(head *ListNode) *ListNode {
	nums := []int{}
	//思路2：转换成数组，然后再排序
	for p := head; p != nil; p = p.Next {
		nums = append(nums, p.Val)
	}
	sort.Ints(nums)
	for p, i := head, 0; p != nil; p, i = p.Next, i+1 {
		p.Val = nums[i]
	}
	return head
}

//时间超时
func sortList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	for cur := head; cur != nil; {
		flag := true
		for p, pNext := cur, cur.Next; pNext != nil; p, pNext = pNext, pNext.Next {
			if pNext.Val < p.Val {
				//交换两个结点的值
				p.Next.Val, p.Val = p.Val, p.Next.Val
				flag = false
			}
		}
		if flag {
			break
		}
	}
	return head
}

//201. 数字范围按位与
func rangeBitwiseAnd(m int, n int) int {
	res := 0
	//高位
	//m,n,从高位往低比较，直到不相等的，前面相等的作为结果的高位
	//也就是说后面全都是0？
	//first := math.MinInt32
	var first uint32 = 0b1000_0000_0000_0000_0000_0000_0000_0000
	for i := 31; i >= 0 && uint32(m)&first == uint32(n)&first; i-- {
		//求第一位
		res += int(uint32(m) & first)
		first >>= 1
	}
	return res
}

//每日一题：767. 重构字符串
func reorganizeString(S string) string {
	res := ""
	//记录所有字符串的个，
	m := make([]int, 26)
	for _, v := range S {
		m[v-'a'] += 1
		//如果某一个大于一半的话，直接返回""
		if m[v-'a'] > (len(S)+1)/2 {
			return res
		}
	}
	//拼接结果返回
	nextLetter := func(cur int) int {
		res := -1
		max := 0
		for k, v := range m {
			if v > max && k != cur {
				max = v
				res = k
			}
		}
		if res != -1 {
			m[res] -= 1
		}
		return res
	}
	//从里面找最多的哪个作为下一个字母
	next := nextLetter(-1)

	for next != -1 {
		res += string(next + 'a')
		next = nextLetter(next)
	}

	return res
}
