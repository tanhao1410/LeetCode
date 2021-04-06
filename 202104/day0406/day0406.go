package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3}))

	arr := []int{1, 1}
	dst := make([]int, 3)
	copy(dst, arr)
	fmt.Println(dst)
}

//797. 所有可能的路径
func allPathsSourceTarget(graph [][]int) [][]int {
	//图的遍历 从 0 --> n - 1
	//从0 可以到[...] // 再分别从里面的元素接着往下走，广度优先遍历。
	res := [][]int{}

	//下一步走法

	var nextStep func(preStep []int)
	nextStep = func(preStep []int) {
		//看它的前一步的最后一步可以走到哪？
		last := preStep[len(preStep)-1]
		if last == len(graph)-1 {
			//走到最后了，结束
			res = append(res, preStep)
			return
		}

		for _, v := range graph[last] {
			//新建一个新的路径，开始走下一步
			newStep := make([]int, len(preStep)+1)
			copy(newStep, preStep)
			newStep[len(preStep)] = v
			nextStep(newStep)
		}
	}
	nextStep([]int{0})

	return res
}

//每日一题：80. 删除有序数组中的重复项 II
func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	//双指针的方式来
	i, j := 0, 1
	//i指向最终返回的数组的最后一个元素，j指向元数组中的元素

	for ; j < len(nums); j++ {
		//前面已经有两个相同的了，跳过该数
		if i >= 1 && nums[j] == nums[i] && nums[j] == nums[i-1] {
			fmt.Println(j, nums[j])
		} else {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
