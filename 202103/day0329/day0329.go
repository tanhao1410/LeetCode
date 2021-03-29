package main

import (
	"fmt"
)

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)

	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}

//367. 有效的完全平方数
func isPerfectSquare(num int) bool {
	for i := 1; ; i++ {
		if i*i < num {
			continue
		} else if i*i == num {
			return true
		} else {
			return false
		}
	}
}

//146. LRU 缓存机制
type LRUCache struct {
	M        map[int]*Node
	Capacity int
	Tail     *Node
	Head     *Node
}

//双向链表节点
type Node struct {
	Val  int
	Key  int
	Pre  *Node
	Next *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		M:        make(map[int]*Node),
		Tail:     nil,
		Head:     nil,
	}
}

func (this *LRUCache) Get(key int) int {
	//先看存不存在
	if node, ok := this.M[key]; ok {
		//需要将它移动至头部
		if this.Head != node {
			//它的前一个指向自己的后一个
			node.Pre.Next = node.Next

			//如果当前节点不是最后一个
			if node.Next != nil {
				//它后面还有节点，它后面的指向它的前面
				node.Next.Pre = node.Pre
			} else {
				//说明节点是尾，更新尾部
				this.Tail = node.Pre
				this.Tail.Next = nil
			}

			//当前节点指向原来的头节点
			node.Next = this.Head

			//原来的头节点需要指向当前
			this.Head.Pre = node

			//作为新的头
			this.Head = node
			node.Pre = nil

		}
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	//先看存不存在
	old := this.Get(key)

	if old != -1 && old != value {
		//存在但值不相等，直接更新值即可
		this.M[key].Val = value
	} else if old == -1 {
		//不存在

		//看大小是否超出界限
		if this.Capacity == len(this.M) {
			//超出界限了，直接就用最后一个节点来充当新节点

			//删除最后一个key
			newHead := this.Tail
			delete(this.M, newHead.Key)
			newHead.Val = value
			newHead.Key = key
			//新增加一个key
			this.M[key] = newHead

			this.Get(key)
		} else if len(this.M) == 0 {
			//第一个节点
			newHead := &Node{
				Val: value,
				Key: key,
			}
			this.Head = newHead
			this.Tail = newHead

			this.M[key] = newHead
		} else {
			newHead := &Node{
				Val:  value,
				Key:  key,
				Next: this.Head,
			}
			this.Head.Pre = newHead
			this.Head = newHead

			this.M[key] = newHead
		}

	}
	//如果存在，并且值相等，直接返回即可
}

//120. 三角形最小路径和
func minimumTotal(triangle [][]int) int {
	//如果长度为1,直接就返回了
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			//它可以从上面过来，要么两条路，要么一条路
			if j < 1 {
				triangle[i][j] += triangle[i-1][0]
			} else if j == len(triangle[i])-1 {
				triangle[i][j] += triangle[i-1][len(triangle[i])-2]
			} else {
				//从哪边小从哪边走
				if triangle[i-1][j-1] > triangle[i-1][j] {
					triangle[i][j] += triangle[i-1][j]
				} else {
					triangle[i][j] += triangle[i-1][j-1]
				}
			}
		}
	}

	//从最后一行中找最小的数
	lastRow := triangle[len(triangle)-1]
	res := lastRow[0]
	for i := 1; i < len(lastRow); i++ {
		if lastRow[i] < res {
			res = lastRow[i]
		}
	}

	return res
}
