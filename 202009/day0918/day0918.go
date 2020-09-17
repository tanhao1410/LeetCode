package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{2, 1, 1, 3}
	fmt.Println(permuteUnique(nums))
}

//全排列2
//给定一个可包含重复数字的序列，返回所有不重复的全排列
func permuteUnique(nums []int) [][]int {
	//思路：先排序，再递归
	//先选择一个数，再在剩下的里面选一个。依次循环
	res := &[][]int{}
	slice := sort.IntSlice(nums)
	slice.Sort()
	permute2(nums, []int{}, res)
	return *res
}

func permute2(nums []int, parts []int, res *[][]int) {
	//所有的数字都选完了，即递归结束
	if 0 == len(nums) {
		*res = append(*res, parts)
		return
	}
	//从剩下的数字中选择，剩下的数字为方法传递过来的数，不选择相同的数
	prev := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] == prev {
			continue
		}
		nums2 := []int{}
		for j := 0; j < len(nums); j++ {
			if j != i {
				nums2 = append(nums2, nums[j])
			}
		}
		part2 := []int{}
		for _, v := range parts {
			part2 = append(part2, v)
		}
		part2 = append(part2, nums[i])
		prev = nums[i]
		permute2(nums2, part2, res)
	}
}
