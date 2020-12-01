package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(searchRange([]int{8, 8, 8}, 8))
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(numberToWords(102000))
}

//273. 整数转换英文表示
func numberToWords(num int) string {
	res := ""
	//对应关系
	m := map[int]string{
		0: "Zero", 1: "One", 2: "Two", 3: "Three", 4: "Four", 5: "Five", 6: "Six", 7: "Seven", 8: "Eight", 9: "Nine",
		10: "Ten", 11: "Eleven", 12: "Twelve", 13: "Thirteen", 14: "Fourteen", 15: "Fifteen", 16: "Sixteen", 17: "Seventeen", 18: "Eighteen", 19: "Nineteen",
		20: "Twenty", 30: "Thirty", 40: "Forty", 50: "Fifty", 60: "Sixty", 70: "Seventy", 80: "Eighty", 90: "Ninety",
	}
	//先确定最高位是多少
	if num >= 1000000000 {
		billionNum := num / 1000000000
		res += m[billionNum]
		res += " Billion "
		num = num % 1000000000
	}

	//确定million
	if num >= 1000000 {
		millionNum := num / 1000000
		res += numberToWords(millionNum)
		res += " Million "
		num = num % 1000000
	}

	//确定千位
	if num >= 1000 {
		thousandNum := num / 1000
		res += numberToWords(thousandNum)
		res += " Thousand "
		num = num % 1000
	}

	//确定百位
	if num >= 100 {
		hundredNum := num / 100
		res += m[hundredNum]
		res += " Hundred "
		num = num % 100
	}

	//确定十位
	if num <= 20 {
		res += m[num]
	} else {
		//
		tenNum := num / 10
		res += m[tenNum*10]
		if num%10 != 0 {
			res += " "
			res += m[num%10]
		}
	}
	return strings.ReplaceAll(res, " Zero", "")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//230. 二叉搜索树中第K小的元素
func kthSmallest(root *TreeNode, k int) int {
	//思路：判断左边和右边有多少元素
	//看左边有多少比它小的
	smallerCount := nodeCount(root.Left)
	if smallerCount < k-1 {
		return kthSmallest(root.Right, k-smallerCount-1)
	} else if smallerCount == k-1 {
		return root.Val
	} else {
		return kthSmallest(root.Left, k)
	}
}

//求树的节点个数
func nodeCount(root *TreeNode) int {
	if root != nil {
		return 1 + nodeCount(root.Right) + nodeCount(root.Left)
	}
	return 0
}

//209. 长度最小的子数组
func minSubArrayLen(s int, nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	res := len(nums)
	//思路：
	//nums[0]加到什么地方 > s,然后记录一个res
	//dp[i].val 满足条件的子数组之和
	//dp[i].count 满足条件的子数组大小
	//inums[i] 为起点的子数组
	//dp[i+1] dp[i]有关系
	//dp[0]易求
	dp := make([][2]int, len(nums))
	//求dp[0]
	i := 0
	for ; i < len(nums) && dp[0][0] < s; i++ {
		dp[0][0] += nums[i]
	}
	if dp[0][0] < s {
		return 0
	}
	dp[0][1] = i
	if i < res {
		res = i
	}
	//求剩余的其它dp
	for i = 1; i < len(nums); i++ {
		//首先判断是否减去了一个，还能满足>=s
		if dp[i-1][0]-nums[i-1] >= s {
			dp[i][0] = dp[i-1][0] - nums[i-1]
			dp[i][1] = dp[i-1][1] - 1
		} else {
			//加一个数，减一个数，看是否还大于s
			//减去前一个数的
			sum := dp[i-1][0] - nums[i-1]
			j := i - 1 + dp[i-1][1]
			for ; j < len(nums) && sum < s; j++ {
				sum += nums[j]
			}
			//没有符合的了，直接返回最小的即可
			if sum < s {
				return res
			}
			dp[i][0] = sum
			dp[i][1] = j - i
		}
		if dp[i][1] < res {
			res = dp[i][1]
		}
	}
	return res
}

//每日一题：34. 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	if nums == nil || len(nums) == 0 {
		return []int{-1, -1}
	}
	res := []int{}

	//先二分查找找到开始
	low, high := 0, len(nums)-1
	middle := (high + low) / 2
	start := -1
	for low <= high {
		if nums[middle] > target {
			high = middle - 1
		} else if nums[middle] == target {
			start = middle
			high = middle - 1
		} else {
			low = middle + 1
		}
		middle = (high + low) / 2
	}
	if start != -1 {
		res = append(res, start)
	} else {
		//没找到
		return []int{-1, -1}
	}

	low, high = 0, len(nums)-1
	middle = (high + low) / 2
	end := -1
	for low <= high {
		if nums[middle] > target {
			high = middle - 1
		} else if nums[middle] == target {
			end = middle
			low = middle + 1
		} else {
			low = middle + 1
		}
		middle = (high + low) / 2
	}
	res = append(res, end)
	return res
}
