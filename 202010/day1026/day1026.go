package main

func main() {
	nums := []int{5, 0, 10, 0, 10, 6}
	print(smallerNumbersThanCurrent(nums))
}

//1365.有多少小于当前数字的数字
func smallerNumbersThanCurrent(nums []int) []int {
	//思路，用一个数组统计每个数字出现的次数，
	//从小到大，一次遍历，得到每个数字前面有多少小于自己的数的数组
	//改变入参的数组，返回即可
	numCount := make([]int, 101)
	for _, v := range nums {
		numCount[v] += 1
	}

	//比当前数小的=比前一个数小的数量+前一个数的数量
	preCount := numCount[0]
	numCount[0] = 0
	for i := 1; i < 101; i++ {
		numCount[i], preCount = numCount[i-1]+preCount, numCount[i]
	}

	//改变参数，返回结果
	for i := 0; i < len(nums); i++ {
		nums[i] = numCount[nums[i]]
	}

	return nums
}
