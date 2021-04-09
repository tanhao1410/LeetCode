package main

//
func main() {

}

//剑指 Offer 57 - II. 和为s的连续正数序列
func findContinuousSequence(target int) [][]int {

	//先求最多可能有多少个数
	maxN := []int{0, 1}
	for i := 2; maxN[len(maxN)-1] <= 100000; i++ {
		maxN = append(maxN, maxN[len(maxN)-1]+i)
	}
	res := [][]int{}
	//需要i个连续的数，依次减少一个数来求
	for i := len(maxN); i > 1; i-- {
		//当i为偶数时
		if i%2 == 0 && target%(i/2) == 0 {
			dHead := target*2/i - i + 1
			if dHead > 0 && dHead%2 == 0 {
				//head = tail + 1 - i
				item := make([]int, i)
				head := dHead / 2
				for j := 0; j < i; j++ {
					item[j] = head
					head++
				}
				res = append(res, item)
			}
		} else if i%2 == 1 && target%i == 0 {
			//需要奇数个，
			head := target/i - i/2
			if head > 0 {
				item := make([]int, i)
				for j := 0; j < i; j++ {
					item[j] = head
					head++
				}
				res = append(res, item)
			}

		}
	}

	return res

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
