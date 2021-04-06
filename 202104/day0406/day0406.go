package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3}))
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
