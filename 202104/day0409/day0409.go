package main

//
func main() {

}

//剑指 Offer 11. 旋转数组的最小数字
func minArray(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}
	if numbers[len(numbers)-1] > numbers[0] {
		return numbers[0]
	}
	left := minArray(numbers[:(len(numbers)-1)/2+1])
	right := minArray(numbers[(len(numbers)-1)/2+1:])
	if left > right {
		return right
	}
	return left
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
