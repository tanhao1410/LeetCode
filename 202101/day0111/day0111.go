package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(smallestStringWithSwaps("dcab", [][]int{{0, 3}, {1, 2}, {0, 2}}))
}

//225. 用队列实现栈
type MyStack struct {
	queues [][]int
	cur    int
}

func Constructor() MyStack {
	return MyStack{
		queues: make([][]int, 2),
		cur:    0,
	}
}

func (this *MyStack) Push(x int) {
	this.queues[this.cur] = append(this.queues[this.cur], x)
}

func (this *MyStack) Pop() int {
	//如果q1的长度为1,则直接返回即可
	res := -1
	if len(this.queues[this.cur]) == 0 {
		for len(this.queues[1-this.cur]) > 1 {
			head := this.queues[1-this.cur][0]
			this.queues[this.cur] = append(this.queues[this.cur], head)
			this.queues[1-this.cur] = this.queues[1-this.cur][1:]
		}
		res = this.queues[1-this.cur][0]
		this.queues[1-this.cur] = this.queues[1-this.cur][1:]
		return res
	}
	for len(this.queues[this.cur]) > 1 {
		head := this.queues[this.cur][0]
		this.queues[1-this.cur] = append(this.queues[1-this.cur], head)
		this.queues[this.cur] = this.queues[this.cur][1:]
	}
	res = this.queues[this.cur][0]
	this.queues[this.cur] = this.queues[this.cur][1:]
	return res
}

func (this *MyStack) Top() int {
	res := -1
	if len(this.queues[this.cur]) == 0 {
		for len(this.queues[1-this.cur]) > 1 {
			head := this.queues[1-this.cur][0]
			this.queues[this.cur] = append(this.queues[this.cur], head)
			this.queues[1-this.cur] = this.queues[1-this.cur][1:]
		}
		res = this.queues[1-this.cur][0]
		this.cur = 1 - this.cur
		return res
	}
	for len(this.queues[this.cur]) > 1 {
		head := this.queues[this.cur][0]
		this.queues[1-this.cur] = append(this.queues[1-this.cur], head)
		this.queues[this.cur] = this.queues[this.cur][1:]
	}
	res = this.queues[this.cur][0]
	return res
}

func (this *MyStack) Empty() bool {
	return len(this.queues[1-this.cur]) == 0 && len(this.queues[this.cur]) == 0
}

//232. 用栈实现队列
type MyQueue struct {
	Stack1 []int
	Stack2 []int
}

func Constructor2() MyQueue {
	return MyQueue{
		Stack1: []int{},
		Stack2: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.Stack1 = append(this.Stack1, x)
}

func (this *MyQueue) Pop() int {

	if len(this.Stack2) == 0 { //需要将栈1依次弹出，放入栈2中
		for len(this.Stack1) > 0 {
			top := this.Stack1[len(this.Stack1)-1]
			this.Stack2 = append(this.Stack2, top)
			this.Stack1 = this.Stack1[:len(this.Stack1)-1]
		}
	}

	res := this.Stack2[len(this.Stack2)-1]
	this.Stack2 = this.Stack2[:len(this.Stack2)-1]
	return res
}

func (this *MyQueue) Peek() int {
	if len(this.Stack2) == 0 { //需要将栈1依次弹出，放入栈2中
		for len(this.Stack1) > 0 {
			top := this.Stack1[len(this.Stack1)-1]
			this.Stack2 = append(this.Stack2, top)
			this.Stack1 = this.Stack1[:len(this.Stack1)-1]
		}
	}
	res := this.Stack2[len(this.Stack2)-1]
	return res
}

func (this *MyQueue) Empty() bool {
	return len(this.Stack2) == 0 && len(this.Stack1) == 0
}

//每日一题：1202. 交换字符串中的元素
func smallestStringWithSwaps(s string, pairs [][]int) string {
	res := []byte(s)
	//思路：寻找可以互相交换的位置集合。集合内先最小，然后组合起来即可，类似于图了
	c := make([]map[int]bool, 0)

	for _, pair := range pairs {
		noOne := true
		firstC := -1
		for i := 0; i < len(c); i++ {

			//如果又在另一个集合中了，说明两者可以合并
			if firstC != -1 && (c[i][pair[0]] || c[i][pair[1]]) {
				//该位置可以合并了。
				for k, _ := range c[i] {
					c[firstC][k] = true
					delete(c[i], k)
				}
				break
			}

			//有一个在集合中，都加入集合
			if c[i][pair[0]] || c[i][pair[1]] {
				c[i][pair[0]], c[i][pair[1]] = true, true
				firstC = i
				noOne = false
			}

		}

		if noOne {
			m := make(map[int]bool)
			m[pair[0]], m[pair[1]] = true, true
			c = append(c, m)
		}
	}
	//最后形成的集合中存有重复的，需要去重

	//c的数量即最后的可以互换的数量。集合内先统计各字母的数量，然后拼接形成最小的。
	for i := 0; i < len(c); i++ {
		byteN := make([]int, 26)
		sortIndex := []int{}
		for k, _ := range c[i] {
			byteN[s[k]-'a']++
			sortIndex = append(sortIndex, k)
		}

		temp := make([]byte, len(sortIndex))
		t := 0
		for j := 0; j < 26; j++ {
			for k := 0; k < byteN[j]; k++ {
				temp[k+t] = byte(j + 'a')
			}
			t += byteN[j]
		}

		//排序过后的下标集合。
		sort.Ints(sortIndex)
		for j := 0; j < len(sortIndex); j++ {
			res[sortIndex[j]] = temp[j]
		}

	}

	return string(res)
}
