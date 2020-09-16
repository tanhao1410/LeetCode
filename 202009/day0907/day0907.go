package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}

	fmt.Println(topKFrequent(nums, 2))
}

//返回频率最高的前k个数
func topKFrequent(nums []int, k int) []int {
	res := make([]int, k)
	m := make(map[int]int, k)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}

	for _, count := range m {
		SetSmallNum(count, res)
	}
	smallCount, i := res[0], 0
	for num, count := range m {
		if count >= smallCount {
			res[i] = num
			i++
		}
	}

	return res
}

func SetSmallNum(num int, nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		if num > nums[i] {
			j := 0
			for ; j < i && nums[j] == 0; j++ {
			}
			//从j--i都要往前挪
			if j != 0 {
				for ; j <= i; j++ {
					nums[j-1] = nums[j]
				}
			} else {
				for j = 1; j <= i; j++ {
					nums[j-1] = nums[j]
				}
			}
			nums[i] = num //前面的都要往前挪
			break
		}
	}
}
