package main

//
func main() {

}

//每日一题：154. 寻找旋转排序数组中的最小值 II
func findMin(nums []int) int {
	//最后一位大于首位，说明是有序的。
	if len(nums) == 1 {
		return nums[0]
	}

	if nums[len(nums)-1] > nums[0] {
		return nums[0]
	}

	middle := (len(nums) - 1) / 2
	//从中间切开，肯定有一半是有序的。
	left := findMin(nums[:middle+1])
	right := findMin(nums[middle+1:])

	if left > right {
		return right
	}
	return left
}
