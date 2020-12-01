package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{8, 8, 8}, 8))
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
