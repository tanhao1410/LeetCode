package main

//从一个数字数字中找出两个相加等于一个数，假设这是数存在
func main() {
	//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
	//
	//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

}

func twoSum(nums []int, target int) []int {

	numMap := map[int]int{} //key为值，value为数的下表
	for i, v := range nums {
		find := target - v
		if i2, ok := numMap[find]; ok {
			return []int{i2, i}
		}
		numMap[v] = i
	}
	return nil
}

func twoSum3(nums []int, target int) []int {
	//首先不能用判断存不存在来解决了。因为可能这个8,4+4

	m := map[int]int{}

	for i, v := range nums {
		r := target - v

		//if v,ok := m[r];ok{
		//
		//}
		//

		if m[r] != 0 {
			return []int{m[r], i}
		}

		m[v] = i
	}

	return nil
}

func twoSum2(nums []int, target int) []int {
	var res []int
	//首先不能用判断存不存在来解决了。
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res = append(res, i, j)
				return res
			}
		}
	}

	return res
}

//问题：存在相同数字怎么办？
func twoSum1(nums []int, target int) []int {

	//res:= make([]int,2)
	var res []int
	numsMap := make(map[int]int, len(nums))
	for i, v := range nums {
		numsMap[v] = i
	}

	for i := 0; i < len(nums)-1; i++ {
		if numsMap[target-nums[i]] != 0 {
			res = append(res, i, numsMap[target-nums[i]])
		}

	}

	return res

}

func getAnswer(nums []int, sum int) (int, int) {
	//将所有的数存入一个map
	numsMap := make(map[int]bool, len(nums))
	for _, v := range nums {
		numsMap[v] = true
	}

	for i := 0; i < len(nums); i++ {
		if numsMap[sum-nums[i]] {
			return nums[i], sum - nums[i]
		}
	}
	return 0, 0

}
