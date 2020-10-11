package main

import "fmt"

func main() {
	nums := []int{4, 4, 4, 4, 4, 4, 4, 4,
		8, 8, 8, 8, 8, 8, 8, 8,
		12, 12, 12, 12, 12, 12, 12, 12, 16, 16, 16, 16, 16, 16, 16, 16,
		20, 20, 20, 20, 20, 20, 20, 20, 24, 24, 24, 24, 24, 24, 24, 24,
		28, 28, 28, 28, 28, 28, 28, 28, 32, 32, 32, 32, 32, 32, 32, 32,
		36, 36, 36, 36, 36, 36, 36, 36, 40, 40, 40, 40, 40, 40, 40, 40,
		44, 44, 44, 44, 44, 44, 44, 44, 48, 48, 48, 48, 48, 48, 48, 48,
		52, 52, 52, 52, 52, 52, 52, 52, 56, 56, 56, 56, 56, 56, 56, 56,
		60, 60, 60, 60, 60, 60, 60, 60, 64, 64, 64, 64, 64, 64, 64, 64,
		68, 68, 68, 68, 68, 68, 68, 68, 72, 72, 72, 72, 72, 72, 72, 72,
		76, 76, 76, 76, 76, 76, 76, 76, 80, 80, 80, 80, 80, 80, 80, 80,
		84, 84, 84, 84, 84, 84, 84, 84, 88, 88, 88, 88, 88, 88, 88, 88,
		92, 92, 92, 92, 92, 92, 92, 92, 96, 96, 96, 96, 96, 96, 96, 96,
		97, 99}
	fmt.Println(canPartition2(nums))
}

//416.分割等和子集
func canPartition2(nums []int) bool {
	//思路：先求总数，奇书直接返回false.然后，从数组中取出一部分数，使它们的和等于总数的一半
	//回溯法：取第一个数，然后取第二个数，。。。若大于总数了，回溯前一个
	count := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
	}
	if count%2 != 0 {
		return false
	}

	//排序数组
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}

	//第一个数必选
	sum := nums[0]
	return nextNum(sum, 1, count/2, nums)
}

func nextNum(sum, index, target int, nums []int) bool {

	if sum == target {
		return true
	}
	for i := index; i < len(nums); i++ {
		if sum+nums[i] > target {
			return false
		} else if sum+nums[i] == target {
			return true
		} else {
			next := nextNum(sum+nums[i], i+1, target, nums)
			if next {
				return true
			}
			//否则尝试下一个,如果下一个和自己相等，跳过去
			j := i + 1
			for ; j < len(nums) && nums[j] == nums[i]; j++ {
			}
			i = j - 1

		}
	}
	return false
}
