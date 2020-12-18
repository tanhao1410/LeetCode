package main

import (
	"fmt"
	"math/rand"
)

func main() {

	//head := ListNode{
	//	Val: 1,
	//}
	//head.Next = &ListNode{Val: 2}
	//head.Next.Next = &ListNode{Val: 3}
	//solution := Constructor3(&head)
	m := []int{0, 0, 0}
	//for i := 0; i < 10000; i++ {
	//	m[solution.GetRandom()-1]++
	//}
	//fmt.Println(m)

	r := Constructor2()
	fmt.Println(r.Insert(1))
	r.Insert(2)
	r.Insert(3)
	r.Insert(4)
	fmt.Println(r.Remove(4))
	fmt.Println(r.Remove(5))
	fmt.Println(r.Insert(2))
	for i := 0; i < 10000; i++ {
		m[r.GetRandom()-1]++
	}
	fmt.Println(m)
}

//380. 常数时间插入、删除和获取随机元素
type RandomizedSet struct {
	//思路：采用map方式，可以方便做到常数时间删除，插入。获取随机的话，可以用蓄水池问题解决。
	Set map[int]bool
}

func Constructor2() RandomizedSet {
	return RandomizedSet{
		Set: make(map[int]bool),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	res := this.Set[val]
	this.Set[val] = true
	return !res
}

func (this *RandomizedSet) Remove(val int) bool {
	res := this.Set[val]
	delete(this.Set, val)
	return res
}

func (this *RandomizedSet) GetRandom() int {
	res, n := 0, 0
	for k, _ := range this.Set {
		n++
		if rand.Int()%n == 0 {
			res = k
		}
	}
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//382. 链表随机节点
type Solution3 struct {
	//Vals []int
	List *ListNode
}

func Constructor3(head *ListNode) Solution3 {
	//vals := []int{}
	//for ; head != nil; head = head.Next {
	//	vals = append(vals, head.Val)
	//}
	//return Solution3{
	//	Vals: vals,
	//}
	return Solution3{
		List: head,
	}
}

func (this *Solution3) GetRandom() int {
	//return this.Vals[rand.Intn(len(this.Vals))]
	res, n := 0, 0
	for head := this.List; head != nil; head = head.Next {
		n++
		if rand.Int()%n == 0 {
			res = head.Val
		}
	}
	return res
}

//398. 随机数索引
type Solution struct {
	m map[int][]int
}

func Constructor(nums []int) Solution {
	m := make(map[int][]int)
	for k, v := range nums {
		if _, ok := m[v]; ok {
			m[v] = append(m[v], k)
		} else {
			m[v] = []int{k}
		}
	}
	return Solution{
		m: m,
	}
}

func (this *Solution) Pick(target int) int {
	ints := this.m[target]
	index := rand.Intn(len(ints))
	return this.m[target][index]
}

//每日一题：389. 找不同
func findTheDifference(s string, t string) byte {
	m := make([]int, 26)
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']--
		m[t[i]-'a']++
	}
	m[t[len(s)]-'a']++
	for i := 0; ; i++ {
		if m[i] > 0 {
			return byte(i + 'a')
		}
	}
}
