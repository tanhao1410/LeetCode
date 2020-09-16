package day0818

import (
	"fmt"
)

type String string

func (s String) Len() int {
	return len(s)
}

func main() {
	var s String = "123"
	//strings.Contains()
	fmt.Println(s.Len())
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//将非负整数转换为其对应的英文表示。可以保证给定输入小于 231 - 1
//1234567891
//输出: "One Billion Two Hundred Thirty Four Million
//Five Hundred Sixty Seven Thousand Eight Hundred Ninety One

func numberToWords(num int) string {
	m := make(map[int]string)
	m[0] = "Zero"
	m[1] = "One"
	m[2] = "Two"
	m[3] = "Three"
	m[4] = "Four"
	m[5] = "Five"
	m[6] = "Six"
	m[7] = "Seven"
	m[8] = "Eight"
	m[9] = "Nine"

	return ""
}

//判断递增递减数列
func isMonotonic(A []int) bool {
	if len(A) < 2 {
		return true
	}
	i := 1
	d := 0
	incre, decre := false, false
	for ; i < len(A); i++ {
		d = A[i] - A[i-1]
		if d > 0 {
			incre = true
		} else if d < 0 {
			incre = true
		}
	}

	return !(incre && decre)
}

//转小写
func toLowerCase(str string) string {
	bytes := []byte(str)
	var dis int = 'a' - 'A'
	for i, v := range bytes {
		if 'A' <= v && v <= 'Z' {
			bytes[i] = byte(int(v) + dis)
		}
	}
	return string(bytes)
}

//给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次
func deleteDuplicates(head *ListNode) *ListNode {
	res := head
	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return res
}

//反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
//1 ≤ m ≤ n ≤ 链表长度。
func reverseBetween(head *ListNode, m int, n int) *ListNode {

	//再翻转

	//再拼接
	return nil

}

//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回
func reversePrint(head *ListNode) []int {
	res := []int{}
	//最简单的思路是先统一放到数组切片中，然后倒置切片
	//难点在于倒置数组。一个额外的空间，首尾交换，直到两者相等，或者，首大于尾。
	lenth := -1
	for ; head != nil; head = head.Next {
		res = append(res, head.Val)
		lenth++
	}

	for i := 0; lenth > i; i++ {
		res[i], res[lenth] = res[lenth], res[i]
		lenth--
	}

	return res
}
