package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(minSteps(1000))
}

//650. 只有两个键的键盘
func minSteps(n int) int {
	//求一千以内的素数。
	//如果n是素数，直接返回即可。否则，它肯定是某个数的倍数。
	suNum := map[int]bool{2: true, 3: true}
	suNumSeq := []int{2, 3}
	for i := 4; i <= n; i++ {
		j := 2
		for ; j*j <= i; j++ {
			if i%j == 0 {
				break
			}
		}
		if j*j > i {
			suNum[i] = true
			suNumSeq = append(suNumSeq, i)
		}
	}

	var minStep2 func(suNum map[int]bool, suNumSeq []int, n int) int
	minStep2 = func(suNum map[int]bool, suNumSeq []int, n int) int {
		if suNum[n] {
			return n
		}
		for _, v := range suNumSeq {
			if n%v == 0 {
				return minStep2(suNum, suNumSeq, n/v) + v
			}
		}
		return 0
	}

	return minStep2(suNum, suNumSeq, n)
}

//622. 设计循环队列
type MyCircularQueue struct {
	num []int
	//front指向头
	front int
	//rear指向下一个要插入的位置，
	rear int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		num:   make([]int, k+1),
		front: 0,
		rear:  0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	//插入一个元素
	this.num[this.rear] = value
	this.rear = (this.rear - 1 + len(this.num)) % len(this.num)
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	//移除一个数
	this.front = (this.front - 1 + len(this.num)) % len(this.num)
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.num[this.front]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.num[(this.rear+1)%len(this.num)]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.front == this.rear
}

func (this *MyCircularQueue) IsFull() bool {
	return this.front == (this.rear-1+len(this.num))%len(this.num)
}

//648. 单词替换
func replaceWords(dictionary []string, sentence string) string {
	res := ""
	//先把词典变成map方便查询
	dic := make(map[string]bool)
	for _, v := range dictionary {
		dic[v] = true
	}

	words := strings.Split(sentence, " ")
	for _, word := range words {
		for i := 1; i <= len(word); i++ {
			if dic[word[:i]] {
				res += word[:i]
				break
			}
			if i == len(word) {
				res += word
			}
		}
		res += " "
	}

	return strings.TrimSpace(res)
}

//每日一题：628. 三个数的最大乘积
func maximumProduct(nums []int) int {

	//全正数-最大的三个数，全负数-最大的三个数，一个正数其他为负-最小的两个负数*最大的正数，就一个负数-最大的三个数，多个正多个负-最小的两个负*最大的正或 最大的三个数
	//分析可知，记录最大的三个数，最小的两个数，共五个数
	fir, sec, thi := -10000, -10000, -10000
	min, min2 := 10000, 10000
	for _, v := range nums {
		//记录最大的三个数
		if v > fir {
			fir, sec, thi = v, fir, sec
		} else if v > sec {
			sec, thi = v, sec
		} else if v > thi {
			thi = v
		}

		//记录最小的两个数
		if v < min {
			min, min2 = v, min
		} else if v < min2 {
			min2 = v
		}

	}
	max1 := fir * sec * thi
	max2 := fir * min * min2
	if max1 > max2 {
		return max1
	}
	return max2
}
