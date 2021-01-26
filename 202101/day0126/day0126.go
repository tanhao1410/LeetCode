package main

import "sort"

func main() {

}

//面试题 16.17. 连续数列
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] = nums[i-1] + nums[i]
		}
		if nums[i] > res {
			res = nums[i]
		}
	}
	return res
}

//面试题 08.04. 幂集
func subsets(nums []int) [][]int {
	res := [][]int{[]int{}}
	//先选一个数
	for i := 0; i < len(nums); i++ {
		re := subsets(nums[i+1:])
		for _, v := range re {
			res = append(res, append(v, nums[i]))
		}
	}
	return res
}

//每日一题：1128. 等价多米诺骨牌对的数量
func numEquivDominoPairs(dominoes [][]int) int {
	//可以由数组改为大的*10 + 小的，相等的就是可替换的。问题就变成了普通的数组，就相等的数的数量了。
	nums := make([]int, len(dominoes))
	for k, dominoe := range dominoes {
		if dominoe[0] > dominoe[1] {
			nums[k] = dominoe[0]*10 + dominoe[1]
		} else {
			nums[k] = dominoe[1]*10 + dominoe[0]
		}
	}

	//排序下
	sort.Ints(nums)
	res := 0
	pre := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == pre {
			count++
			//最后一个数特殊处理
			if i == len(nums)-1 {
				res += (count * (count - 1) / 2)
			}
		} else {
			res += (count * (count - 1) / 2)
			count = 1
			pre = nums[i]
		}
	}

	return res
}
