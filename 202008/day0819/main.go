package day0819

//T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2
//给你整数 n，请返回第 n 个泰波那契数 Tn 的值,0 <= n <= 37
func tribonacci(n int) int {

	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	}
	//地柜求法效率很低的
	//类似动态规划，从低往高求
	var nums []int = make([]int, n+1)
	nums[1], nums[2] = 1, 2
	for i := 3; i < n; i++ {
		nums[i] = nums[i-1] + nums[i-2] + nums[i-3]
	}

	return nums[n]
}

//优化空间复杂度
func tribonacci2(n int) int {

	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	}
	//地柜求法效率很低的
	//类似动态规划，从低往高求

	p, q, j := 0, 1, 1
	for i := 3; i <= n; i++ {
		j, p, q = j+p+q, q, j
	}

	return j
}

//统计一个数字在排序数组中出现的次数
func search(nums []int, target int) int {

	//思路：先用二分查找，找到这个数，
	start, end, middle := 0, len(nums), len(nums)/2
	flag := false
	for start > end {
		if target == nums[middle] {
			//找到了
			flag = true
			break
		} else if target > nums[middle] {
			//目标数比中间数要大
			start = middle + 1

		} else {
			end = middle - 1

		}
		middle = (start + end) / 2
	}

	if flag {
		res := 0
		//以middle为中心，从两边找相等的数，直到都没有了
		for i := middle + 1; i < len(nums) && nums[i] == target; i++ {
			res++
		}
		for i := middle - 1; i >= 0 && nums[i] == target; i-- {
			res++
		}
		return res
	}

	return 0
}

//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素）
func maxSubArray(nums []int) int {
	//思路：先来一个数组，记录，从前面不小于0的位置开始加到本处的值。然后从中找出最大的即可。

	if len(nums) < 1 {
		return 0
	}

	a := nums[0]

	res := nums[0]

	for i := 1; i < len(nums); i++ {
		if a <= 0 {
			a = nums[i]
		} else {
			a = a + nums[i]
		}
		//记录最大的值
		if a > res {
			res = a
		}
	}
	return res

}
