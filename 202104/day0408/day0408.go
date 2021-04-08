package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

func main() {
	//
	for i := 0; i < 200; i++ {
		fmt.Print(rand.Intn(1000), ",")
	}
}

//18. 四数之和
func fourSum(nums []int, target int) [][]int {
	// 先排序，从小到大
	sort.Ints(nums)
	res := [][]int{}
	//从中选出两个
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {

			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			preSum := nums[i] + nums[j]
			//在接下来的空间中找到 target - preSum

			m := make(map[int]int)
			//记录上一个
			pre := math.MaxInt32
			for k := j + 1; k < len(nums); k++ {
				if v, ok := m[target-preSum-nums[k]]; ok && k > v { //比前面的存在的下标要大才可以
					//不能和上一步用的数字相同
					if pre == math.MaxInt32 || pre != nums[k] {
						//找到一个
						item := []int{nums[i], nums[j], target - preSum - nums[k], nums[k]}
						res = append(res, item)
					}
					pre = nums[k]
				} else {
					m[nums[k]] = k
				}
			}
		}
	}
	return res
}
