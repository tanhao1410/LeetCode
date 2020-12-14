package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(compress([]byte{'a', 'a', 'b', 'b', 'b', 'c', 'a', 'b', 'b', 'c'}))
}

type Node struct {
	Val      int
	Children []*Node
}

//443. 压缩字符串
func compress(chars []byte) int {

	numToStr := func(count int) []byte {
		res := []byte{}
		for ; count > 0; count = count / 10 {
			res = append(res, []byte(strconv.Itoa(count%10))...)
		}

		return res
	}

	pre, count := chars[0], 1
	//原地修改的话，需要一个index指向要写入的位置
	index := 1
	for i := 1; i < len(chars); i++ {
		if chars[i] == pre {
			count++
		} else if count == 1 {
			//如果和上一个不相等了，说明上一个计算结束了，但是因为count=1，因此，不用设置数字。但该位置需要设置为新的当前字母
			chars[index] = chars[i]
			index++
		} else {

			//chars[index] = byte(count + '0')
			countBytes := numToStr(count)
			for j := 0; j < len(countBytes); j++ {
				chars[index] = countBytes[len(countBytes)-1-j]
				index++
			}

			//设置当前字母
			chars[index] = chars[i]
			//指向下一个应该设置的地方
			index += 1
			count = 1
		}
		pre = chars[i]
	}

	if count > 1 {
		countBytes := numToStr(count)
		for j := 0; j < len(countBytes); j++ {
			chars[index] = countBytes[len(countBytes)-1-j]
			index++
		}
	}

	return index
}

//429. N 叉树的层序遍历
func levelOrder(root *Node) [][]int {
	res := [][]int{}

	if root == nil {
		return res
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		item := []int{}
		itemLen := len(queue)
		for i := 0; i < itemLen; i++ {
			item = append(item, queue[i].Val)
			if queue[i].Children != nil {
				queue = append(queue, queue[i].Children...)
			}
		}
		queue = queue[itemLen:]
		res = append(res, item)
	}
	return res
}
