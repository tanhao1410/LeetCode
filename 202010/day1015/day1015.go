package main

import "fmt"

func main() {
	nums := []int{1,3,1,1,1}
	fmt.Println(search(nums, 3))
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
