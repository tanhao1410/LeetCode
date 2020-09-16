package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//旋转链表
func rotateRight(head *ListNode, k int) *ListNode {
	//一步一步走
	if head == nil || head.Next == nil {
		return head
	}
	for ; k > 0; k-- {
		temp, pre := head, head
		for ; head.Next != nil; head = head.Next {
			pre = head
		}
		head.Next, pre.Next = temp, nil
	}
	return head
}

//时间优化
func rotateRight2(head *ListNode, k int) *ListNode {
	//一步一步走
	if head == nil || head.Next == nil {
		return head
	}

	count := 0
	for p := head; p != nil; p = p.Next {
		count++
	}

	for k = k % count; k > 0; k-- {
		temp, pre := head, head
		for ; head.Next != nil; head = head.Next {
			pre = head
		}
		head.Next, pre.Next = temp, nil
	}
	return head
}

func main() {
	//nums := []int{2,2,2,0,2,2}
	//fmt.Println(search(nums,1))
	//fmt.Println(lengthOfLastWord(" t     bb  "))
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum(grid))
}

//最小路径和
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	dp := make([][]int, len(grid))
	//初始化dp
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[0]))
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < len(grid); i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < len(grid[0]); j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < len(grid); i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
		for j := 1; j < len(grid[0]); j++ {
			if dp[i-1][j] > dp[i][j-1] {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			}
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

func lengthOfLastWord(s string) int {

	flag := true
	i, j := 0, 0
	for ; i < len(s); i++ {
		if s[i] != ' ' {
			if flag {
				j = i
				flag = false
			}
		} else {
			j++
			flag = true
		}
	}
	return i - j
}

func search(nums []int, target int) bool {

	if len(nums) == 0 {
		return false
	}

	if len(nums) == 1 {
		return target == nums[0]
	}

	start, end := 0, len(nums)-1
	midddle := (end + start) / 2
	if nums[len(nums)-1] <= nums[0] {
		for start <= end {

			if midddle-1 >= 0 && nums[midddle] < nums[midddle-1] {
				//比前面的数还要小，说明此处即为断层
				break
			} else {
				//需要判断这个数和首位比谁大谁小
				if nums[midddle] >= nums[0] {
					//断层处还在后面,这边不能保证的！！
					start = midddle + 1
					midddle = (end + start) / 2
				} else {
					end = midddle - 1
					midddle = (end + start) / 2
				}
			}
		}
		//判断数字应该是在左边还是右边
		if target >= nums[0] && midddle-1 >= 0 && target <= nums[midddle-1] {
			//在这里面找
			start, end = 0, midddle-1
		} else if target >= nums[midddle] && target <= nums[len(nums)-1] {
			//在后面数组中二分查找
			start, end = midddle, len(nums)-1
		}

	} else {
		start, end = 0, len(nums)-1
	}

	for midddle := (end + start) / 2; start <= end; {
		if nums[midddle] == target {
			//比前面的数还要小，说明此处即为断层
			return true
		} else {
			//需要判断这个数和首位比谁大谁小
			if nums[midddle] > target {
				//断层处还在后面
				end = midddle - 1
				midddle = (end + start) / 2
			} else {
				start = midddle + 1
				midddle = (end + start) / 2
			}
		}
	}

	return false
}
