package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(increasingTriplet([]int{1, 5, 3, 4}))
}

//334. 递增的三元子序列
func increasingTriplet(nums []int) bool {
	//思路：记录最小值，次小值
	//用一个变量记录需要的数值，只要碰到递增的，就更新该数，遇到比该数大的，说明找到了，返回true
	needMore := math.MaxInt32
	min := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] > needMore {
			//说明找到了第三个数
			return true
		}
		//记录数组前面的最小值
		if nums[i] < min {
			min = nums[i]
		}
		//遇到递增的就可以判断了，要么前面有比小的还小的，要么后面有比大的还大的。
		if i+1 < len(nums) && nums[i+1] > nums[i] {
			if needMore > nums[i+1] {
				//需要更小的即可
				needMore = nums[i+1]
			}
			//如果前面最小的数比前一个数要小，说明找到了
			if nums[i] > min {
				return true
			}
		}
	}
	return false
}

//每日一题：387. 字符串中的第一个唯一字符
func firstUniqChar(s string) int {
	m := make([]int, 26)
	for i := len(s) - 1; i >= 0; i-- {
		m[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if m[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
