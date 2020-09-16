package main

import "fmt"

//给你一个整数数组nums，每次 操作会从中选择一个元素并 将该元素的值减少1。
//
//如果符合下列情况之一，则数组A就是 锯齿数组：
//
//每个偶数索引对应的元素都大于相邻的元素，即A[0] > A[1] < A[2] > A[3] < A[4] > ...
//或者，每个奇数索引对应的元素都大于相邻的元素，即A[0] < A[1] > A[2] < A[3] > A[4] < ...
//返回将数组nums转换为锯齿数组所需的最小操作次数。

func main() {
	//nums := []int{2, 7, 10, 9, 8, 9} //[2,7,10,9,8,9]
	//nums := []int{1,2,3}
	//n := movesToMakeZigzag2(nums)

	nums := []int{6, 8, 6, 8, 0, 4, 1, 2, 9, 9}
	n := findMaxAverage2(nums, 2)
	fmt.Println(n)
}

//给定 n 个整数，找出平均数最大且长度为 k 的连续子数组，并输出该最大平均数
func findMaxAverage2(nums []int, k int) float64 {
	//最简单的思路是一趟循环，返回最大的
	//当往前走的时候，并不是每一次都要计算，当增加的那个数与减去掉的那个数小的话，就没有必要重新计算了

	dp := make([]int, len(nums))
	count, max := 0, 0
	//先计算第一个子数组的大小
	for i := 0; i < k; i++ {
		count += nums[i]
	}
	dp[k-1], max = count, count
	for i := k; i < len(nums); i++ {
		dp[i] = dp[i-1] + nums[i] - nums[i-k] //这个时候的i-k可能已经被改变了。
		if dp[i] > max {
			max = dp[i]
		}
	}
	return float64(max) / float64(k)
}

//给定 n 个整数，找出平均数最大且长度为 k 的连续子数组，并输出该最大平均数
func findMaxAverage(nums []int, k int) float64 {
	//最简单的思路是一趟循环，返回最大的
	//当往前走的时候，并不是每一次都要计算，当增加的那个数与减去掉的那个数小的话，就没有必要重新计算了
	count := 0
	//先计算第一个子数组的大小
	for i := 0; i < k; i++ {
		count += nums[i]
	}
	for i := k; i < len(nums); i++ {
		//问题所在，在走的过程中，可能最大的几个数值在很前面，，通过这样的排除方法不行的。所以这种方法不行。
		if nums[i] > nums[i-k] {
			count += nums[i]
			count -= nums[i-k]
		}
	}
	return float64(count) / float64(k)
}

//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置
func searchInsert(nums []int, target int) int {
	//思路：二分法
	start, end := 0, len(nums)-1
	middle := (start + end) / 2
	for start <= end {
		if nums[middle] > target {
			end = middle - 1
			middle = (start + end) / 2
		} else if target > nums[middle] {
			start = middle + 1
			middle = (start + end) / 2
		} else {
			return middle
		}
	}
	return end + 1
}

//每次减少，而不是增加
func movesToMakeZigzag2(nums []int) int {

	res1, res2 := 0, 0
	pre := 0
	flag := false
	var i int = 0
	for ; i < len(nums); i += 2 {
		if i-1 > 0 && nums[i-1] >= nums[i] {
			res1 += 1 + nums[i-1] - nums[i]
		}
		if flag {
			nums[i-1] = pre
		}
		if i+1 < len(nums) && nums[i+1] >= nums[i] {
			res1 += 1 + nums[i+1] - nums[i]
			pre = nums[i+1]
			nums[i+1] = nums[i] - 1
			flag = true
		} else {
			flag = false
		}
	}

	if flag {
		nums[i-1] = pre
	}

	for i := 1; i < len(nums); i += 2 {
		if i-1 >= 0 && nums[i-1] >= nums[i] {
			res2 += 1 + nums[i-1] - nums[i]
		}
		if i+1 < len(nums) && nums[i+1] >= nums[i] {
			res2 += 1 + nums[i+1] - nums[i]
			nums[i+1] = nums[i] - 1
		}
	}

	if res1 > res2 {
		return res2
	}
	return res1
}

func movesToMakeZigzag(nums []int) int {

	//思路：两次遍历，一次看偶数位，一个看奇数位，返回小的
	//问题：变化后，数值改变了？可以不用改变数值的，只是记录应该增大多少次
	var res1, res2 int = 0, 0

	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			//偶数位，要比前一个大，要比后一个大
			//先确认后面有没有了，以及前面有没有了
			addNum := 0
			if i+1 < len(nums) && nums[i+1] >= nums[i] {
				addNum = 1 + nums[i+1] - nums[i]
				res1 += addNum
			}
			if i-1 > 0 && nums[i-1] >= nums[i]+addNum {
				res1 += 1 + nums[i-1] - nums[i] - addNum
			}
			//上面问题所在，加了两次了,

		} else {
			addNum := 0
			if i+1 < len(nums) && nums[i+1] >= nums[i] {
				addNum = 1 + nums[i+1] - nums[i]
				res2 += addNum
			}
			if i-1 >= 0 && nums[i-1] >= nums[i]+addNum {
				res2 += 1 + nums[i-1] - nums[i] - addNum
			}
		}
	}
	if res1 > res2 {
		return res2
	}
	return res1
}
