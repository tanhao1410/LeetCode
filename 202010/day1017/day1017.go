package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(totalNQueens(10))
	fmt.Println(minMoves([]int{1, 2, 3, 4}))
}

//1.每次+1,判断是否可以
//2.优化1：每次+可以+更多
//3.优化2：如果最大和最小相等，跳出
//4.新方法最大的数比最小的数大多少，则先移动这么多步，直到最大和最小的数相等
//5.优化：少一次循环
//453.最小移动次数使数组元素相等
func minMoves(nums []int) int {
	res := 0
	//用自带的排序来排序
	sort.Ints(nums)
	//先排序
	//sortNums(nums)
	maxIndex, min := len(nums)-1, nums[0]
	dis := nums[maxIndex] - min
	for dis != 0 {
		res += dis

		nums[maxIndex] -= dis
		//每一个都加，相当于只是一个来减
		min = nums[maxIndex]
		maxIndex = maxIndex - 1
		dis = nums[maxIndex] - min
	}
	return res
}

func sortNums(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

//52.N皇后 II
func totalNQueens(n int) int {
	//返回总数即可，用一个数组记录填充的皇后，这是关键，回溯法，填充下一个
	chees := make([]int, n)
	res := 0
	next(chees, 0, &res)
	return res
}

//chees 即棋盘，index为下一个要填入的列号，res 记录结果值
func next(chees []int, index int, res *int) {
	if index >= len(chees) {
		*res += 1
	}
	n := len(chees)
	//先确定该位置可以填几
	for i := 0; i < n; i++ {
		can := true
		//行不能重复
		for j := 0; j < index; j++ {
			if chees[j] == i {
				can = false
			}
		}
		//对角线不能重复
		for j := 0; j < index; j++ {
			if (i < chees[j] && chees[j]-i == index-j) ||
				(i > chees[j] && i-chees[j] == index-j) {
				can = false
			}
		}
		if can {
			//都满足了，可以填充了
			chees[index] = i
			//填下一个
			next(chees, index+1, res)
		}
	}
}
