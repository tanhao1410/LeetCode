package main

import "fmt"

func main() {
	//nums := []int{1,3,1,1,1}
	//fmt.Println(search(nums, 3))
	board := [][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}
	fmt.Println(exist(board,"ABCB"))
}

//79.单词搜索
func exist(board [][]byte, word string) bool {
	//用完的时候，置零，回溯后，回复
	first := []int{}
	firstLetter := word[0]
	//找可以的
	for i := 0;i < len(board);i ++{
		for j := 0;j < len(board[0]);j ++{
			if board[i][j] == word[0]{
				first = append(first, i,j)
			}
		}
	}

	for i:=0;i < len(first);i+=2{
		board[first[i]][first[i+1]] = '0'
		if !next(board,first[i],first[i+1],word[1:]){
			board[first[i]][first[i+1]] = firstLetter
		}else{
			return true
		}
	}

	return false
}

func next(board [][]byte,i,j int,word string)bool{
	if len(word) == 0{
		return true
	}
	//看它周围有多少可以的，最多三个
	can := []int{}
	//左边
	fisrtLetter := word[0]
	if i + 1 < len(board) && board[i+1][j] == fisrtLetter {
		can = append(can, i+1,j)
	}
	if i - 1 >= 0 && board[i-1][j] == fisrtLetter {
		can = append(can, i-1,j)
	}
	if j + 1 < len(board[0]) && board[i][j+1] == fisrtLetter {
		can = append(can, i,j+1)
	}
	if j - 1 >= 0 && board[i][j-1] == fisrtLetter {
		can = append(can, i,j-1)
	}

	for i:=0;i < len(can);i+=2{
		board[can[i]][can[i+1]] = '0'
		if !next(board,can[i],can[i+1],word[1:]){
			board[can[i]][can[i+1]] = fisrtLetter
		}else{
			return true
		}
	}
	return false
}



//81.搜索旋转排序数组 II
func search(nums []int, target int) bool {

	index := 0
	//思路：先找断层，通过二分法
	start, end := 0, len(nums)-1
	middle := (start + end) / 2
	for start <= end {

		if nums[start] < nums[end] || start == end {
			//说明是有序的了
			index = start
			break
		} else {
			if nums[middle] > nums[start] {
				start = middle
			} else if nums[middle] < nums[start] {
				end = middle
			} else {
				//end向左移动，直到碰到不相等的
				temp:=middle
				for ; nums[start] == nums[temp] && temp != start; temp-- {
				}
				if temp == start {
					//说明前面全是相等的
					start = middle+1
				} else {
					//前面有不相等的
					end = middle
				}
			}
			middle = (start + end) / 2
		}
	}

	search2 := func(nums []int) bool {
		start, end := 0, len(nums)
		middle := (start + end) / 2
		for start < len(nums) && start <= end {
			if nums[middle] == target {
				return true
			} else if nums[middle] > target {
				end = middle - 1
			} else {
				start = middle + 1
			}
			middle = (start + end) / 2
		}
		return false
	}

	if nums[0] == target {
		return true
	} else if nums[0] > target {
		return search2(nums[index:])
	} else {
		if index == 0 {
			return search2(nums)
		}
		return search2(nums[:index])
	}

	return false
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect2(root *Node) *Node {
	queue := []*Node{root}
	prehead, curhead := root, root
	for len(queue) > 0 {
		//# 当前队列长度
		fmt.Println(len(queue))
		length := len(queue)
		for i := 0; i < length; i++ {
			curhead = queue[0]
			queue = queue[1:]
			if i == 0 {
				prehead = curhead
			} else if i < length-1 {
				prehead.Next = curhead
				prehead = curhead
			} else {
				prehead.Next = curhead
				curhead.Next = nil
			}
			if curhead.Left != nil {
				queue = append(queue, curhead.Left)
			}
			if curhead.Right != nil {
				queue = append(queue, curhead.Right)
			}
		}
	}
	return root

}

//116.填充每个节点的下一个右侧节点指针
func connect(root *Node) *Node {
	//思路：左孩子-->右孩子，右孩子指向父节点右边的左孩子（没有？指向它的右）
	if root == nil || (root.Right == nil && root.Left == nil) {
		return root
	}
	if root.Left != nil {
		if root.Right != nil {
			root.Left.Next = root.Right
			if root.Next != nil {
				if root.Next.Left != nil {
					root.Right.Next = root.Next.Left
				} else if root.Next.Right != nil {
					root.Right.Next = root.Next.Right
				}
			}
		} else {
			if root.Next != nil {
				if root.Next.Left != nil {
					root.Left.Next = root.Next.Left
				} else if root.Next.Right != nil {
					root.Left.Next = root.Next.Right
				}
			}
		}
	} else if root.Right != nil {
		if root.Next != nil {
			if root.Next.Left != nil {
				root.Right.Next = root.Next.Left
			} else if root.Next.Right != nil {
				root.Right.Next = root.Next.Right
			}
		}
	}
	connect(root.Right)
	connect(root.Left)
	return root
}
