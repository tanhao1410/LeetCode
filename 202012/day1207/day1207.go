package main

import "fmt"

func main() {
	fmt.Println(matrixScore([][]int{{0, 0, 0, 1, 1, 0}, {1, 1, 0, 1, 0, 1}, {1, 0, 1, 0, 0, 1}}))
}

//1438. 绝对差不超过限制的最长连续子数组
func longestSubarray(nums []int, limit int) int {
	max, min := 0, 0
	res := 0
	curLen := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] >= nums[min] && nums[i] <= nums[max] {
			//介于最大与最小之间
			curLen++
		} else if nums[i] > nums[max] && nums[i]-nums[min] <= limit {
			//比最大的要大，但不超范围
			curLen++
			max = i
		} else if nums[i] > nums[max] {
			//比最大的要大，超过了范围
			min, max = i, i
			j := i
			//往前面找，看有多少符合
			for ; j >= 0 && nums[i]-nums[j] <= limit; j-- {
				if nums[j] < nums[min] {
					min = j
				}
			}
			if curLen > res {
				res = curLen
			}
			curLen = i - j
		} else if nums[i] < nums[min] && nums[max]-nums[i] <= limit {
			//比最小的要小，但不超范围
			curLen++
			min = i
		} else {
			//比最小的要小，超过了范围
			//往前面找，看有多少符合
			min, max = i, i
			j := i
			for ; j >= 0 && nums[j]-nums[i] <= limit; j-- {
				if nums[j] > nums[max] {
					max = j
				}
			}
			if curLen > res {
				res = curLen
			}
			curLen = i - j
		}
	}
	if curLen > res {
		res = curLen
	}
	return res
}

//每日一题：861. 翻转矩阵后的得分
func matrixScore(A [][]int) int {
	//思路：优先第一列变为1，通过行移动来完成
	for i := 0; i < len(A); i++ {
		if A[i][0] == 0 {
			//移动该行
			for j := 0; j < len(A[0]); j++ {
				if A[i][j] == 1 {
					A[i][j] = 0
				} else {
					A[i][j] = 1
				}
			}
		}
	}
	res := len(A) * 1 << (len(A[0]) - 1)
	//剩下的开始，优先通过列移动，每一列尽量多1，直到完成目标
	for i := 1; i < len(A[0]); i++ {
		oneCount := 0
		for row := 0; row < len(A); row++ {
			if A[row][i] == 1 {
				oneCount++
			}
		}
		if oneCount < (len(A)+1)/2 {
			oneCount = len(A) - oneCount
		}
		num := 1 << (len(A[0]) - i - 1) * oneCount
		res += num
	}
	return res
}
