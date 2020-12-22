package main

import (
	"fmt"
	"math/rand"
)

func main() {
	m := []int{0, 0, 0, 0, 0, 0, 0}
	r := Constructor()
	fmt.Println(r.Insert(9))
	fmt.Println(r.Insert(9))
	fmt.Println(r.Insert(1))
	fmt.Println(r.Insert(1))
	fmt.Println(r.Insert(2))
	fmt.Println(r.Insert(1))
	fmt.Println(r.Remove(2))
	fmt.Println(r.Remove(1))
	fmt.Println(r.Remove(1))
	fmt.Println(r.Insert(9))
	fmt.Println(r.Remove(1))
	fmt.Print(r.GetRandom())
	//for i := 0; i < 10000; i++ {
	//	m[r.GetRandom()]++
	//}
	fmt.Println(m)
}

//381. O(1) 时间插入、删除和获取随机元素 - 允许重复
type RandomizedCollection struct {
	//key -->count
	Data []int
	M    map[int]map[int]bool
}

func Constructor() RandomizedCollection {
	return RandomizedCollection{
		Data: []int{},
		M:    make(map[int]map[int]bool),
	}
}

func (this *RandomizedCollection) Insert(val int) bool {
	indexMap, ok := this.M[val]
	if ok {
		this.Data = append(this.Data, val)
		indexMap[len(this.Data)-1] = true
		this.M[val] = indexMap
		return false
	} else {
		this.Data = append(this.Data, val)
		this.M[val] = make(map[int]bool)
		this.M[val][len(this.Data)-1] = true
		return true
	}
}

func (this *RandomizedCollection) Remove(val int) bool {
	indexMap, _ := this.M[val]
	res := len(indexMap)
	if res > 0 {
		//如果与总数组中的最后一个交换,索引也需要改变！
		//更新最后一个数的索引位置。
		if val != this.Data[len(this.Data)-1] {
			//得到待删除元素的一个索引
			oneIndex := 0
			for k, _ := range indexMap {
				oneIndex = k
				break
			}
			//与最后一个元素交换位置
			this.Data[oneIndex] = this.Data[len(this.Data)-1]

			//删去最后一个元素的索引。
			delete(this.M[this.Data[oneIndex]], len(this.Data)-1)
			//增加一个新索引
			this.M[this.Data[oneIndex]][oneIndex] = true
			//删去最后一个数
			this.Data = this.Data[:len(this.Data)-1]

			//更新val的索引
			delete(this.M[val], oneIndex)
			if len(this.M[val]) == 0 {
				delete(this.M, val)
			}
		} else {
			//要删除的和最后一个数相等
			//删去最后一个元素
			this.Data = this.Data[:len(this.Data)-1]
			//更新val的索引
			delete(this.M[val], len(this.Data))
			if len(this.M[val]) == 0 {
				delete(this.M, val)
			}
		}
	}
	return res > 0
}

func (this *RandomizedCollection) GetRandom() int {
	return this.Data[rand.Intn(len(this.Data))]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//每日一题：103. 二叉树的锯齿形层序遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	flag := true
	for queueLen := len(queue); queueLen > 0; queueLen = len(queue) {
		item := make([]int, queueLen)
		for i := 0; i < queueLen; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			if flag {
				item[i] = queue[i].Val
			} else {
				item[i] = queue[queueLen-1-i].Val
			}
		}
		res = append(res, item)
		flag = !flag
		queue = queue[queueLen:]
	}
	return res
}
