package main

import "sort"

func main() {

}

//287. 寻找重复数
func findDuplicate(nums []int) int {
	//思路：如何判断存在一个重复的，排序后，最后一个数一定是小于1+  n，而不是n+1
	sort.Ints(nums)
	start, end := 0, len(nums)-1
	for middle := (start + end) / 2; ; middle = (start + end) / 2 {
		if end > start && nums[end] == nums[start] {
			return nums[end]
		}
		//重复值在后面
		if middle-start == nums[middle]-nums[start] || end-middle > nums[end]-nums[middle] {
			start = middle
		} else if end-middle == nums[end]-nums[middle] || middle-start > nums[middle]-nums[start] {
			end = middle
		}

	}
}
