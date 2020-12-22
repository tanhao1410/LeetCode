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

//547. 朋友圈
func findCircleNum(M [][]int) int {
	//思路，第一次遍历，形成学生-->朋友s
	//第二次遍历学生，形成学生-->完整朋友圈集合
	//剩下的朋友中继续，直到没有剩余的学生了
	m := make(map[int]map[int]bool) //学生id -->{朋友id:true}
	for i := 0; i < len(M); i++ {
		m[i] = make(map[int]bool)
		for j := 0; j < len(M); j++ {
			if M[i][j] == 1 {
				m[i][j] = true
			}
		}
	}

	//用一个map记录哪些学生已经处理了
	already := make(map[int]bool)
	res := 0
	for i := 0; i < len(M); i++ {
		//没有加入任何朋友圈的才会进入
		if !already[i] {
			res++
			already[i] = true
			for flag := true; flag; {
				//先处理学生0，直到新加入的朋友数为0退出
				flag = false
				for friend, _ := range m[i] {
					already[friend] = true
					//遍历朋友的朋友
					for friendFriend, _ := range m[friend] {
						//如果朋友的朋友不是自己的朋友，那么加入进来，并记录
						if !m[i][friendFriend] {
							m[i][friendFriend] = true
							flag = true
						}
					}
				}
			}
		}
	}

	return res
}

//497. 非重叠矩形中的随机点
type Solution struct {
	Rects   [][]int
	Areas   []int
	AreaSum int
}

func Constructor2(rects [][]int) Solution {
	areaSum := 0
	areas := []int{}

	for _, v := range rects {
		//思路：第一，非重叠，总面积易求。每个矩形的面积也易求。可以求出每个矩形的概率。根据矩形的概率就可以求出对应坐标的概率
		//修改，用面积不好，改为用里面的点的数量
		area := (v[2] - v[0] + 1) * (v[3] - v[1] + 1)
		areaSum += area
		areas = append(areas, areaSum)
	}

	return Solution{
		Rects:   rects,
		Areas:   areas,
		AreaSum: areaSum,
	}
}

func (this *Solution) Pick() []int {
	res := make([]int, 2)
	//先查找落在了哪一个矩形中
	randNum := rand.Intn(this.AreaSum)
	for i := 0; i < len(this.Areas); i++ {
		if randNum < this.Areas[i] {
			//找到了对应的矩形了
			rect := this.Rects[i]
			//随机话 x,y,可以包括边
			res[0] = rand.Intn(rect[2]-rect[0]+1) + rect[0]
			res[1] = rand.Intn(rect[3]-rect[1]+1) + rect[1]
			return res
		}
	}

	return res
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
