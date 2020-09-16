package main

import "fmt"

func main() {
	//fmt.Println(combinationSum3(0,2))

	nums := []int{1, 2, 2, 3, 4, 5}
	fmt.Println(containsNearbyAlmostDuplicate(nums, 3, 0))
}

// i 和 j，使得 nums [i] 和 nums [j] 的差的绝对值小于等于 t ，且满足 i 和 j 的差的绝对值也小于等于 ķ
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	//思路：暴力法的话，便是o(n * k)
	//问题的关键在于如何不回溯，每一次向前比较k个数之后，还需要回去，两个指针呢？想法一样，思路没变
	for i := 1; i < len(nums); i++ {
		for j := i - k; j < i; j++ {
			if j < 0 {
				continue
			}
			if (nums[i]-nums[j] <= t && nums[i] >= nums[j]) || (nums[j]-nums[i] <= t && nums[j] >= nums[i]) {
				//符合条件的
				return true
			}
		}
	}
	return false
}

//组合总数
//找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复
func combinationSum3(k int, n int) [][]int {
	//思路：遍历，大于9的不要，选N个数，依旧是N重递归
	res := &[][]int{}
	//为了保证唯一，选择的数的方式，先选小的，下一个选的必须比前一个大
	selectNextNum(make([]int, 0), k, 0, n, res)
	return *res
}

func selectNextNum(nums []int, k, pre, target int, res *[][]int) {
	if target == 0 && len(nums) == k {
		//说明找到了
		*res = append(*res, nums)
		return
	}
	if target < 0 || len(nums) >= k {
		//该组合找不到
		return
	}
	for i := pre + 1; i <= 9; i++ {
		nums2 := []int{}
		for _, v := range nums {
			nums2 = append(nums2, v)
		}
		nums2 = append(nums2, i)
		selectNextNum(nums2, k, i, target-i, res)
	}
}
