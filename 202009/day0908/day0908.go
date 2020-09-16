package main

import "fmt"

//返回 1 ... n 中所有可能的 k 个数的组合
func combine(n int, k int) [][]int {
	//思路：k个数即k重循环，用递归的方式，一层递归加一个数
	//如果是加入的第k个数，那么加入返回列表中，如果不是，继续递归，组合下一个数
	//保证顺序？下一个数要比之前的数大，这样能保证顺序和不重复。
	res := &[][]int{}
	if k > n || k <= 0 || n <= 0 {
		return *res
	}

	for i := 1; i < n-k+2; i++ {
		nums := []int{}
		nums = append(nums, i)
		createNum(nums, k, n, res)
	}
	return *res
}

func createNum(nums []int, k, n int, res *[][]int) {
	if len(nums) == k {
		*res = append(*res, nums)
		return
	}
	numsLen := len(nums)
	for i := nums[numsLen-1] + 1; i < n-k+len(nums)+2; i++ {
		//nums = append(nums, i) 不能这样用，会影响原来的数组的
		//重新创建一个的话，每一层都会创建一系列数组，也不好。
		nums2 := []int{}
		for _, v := range nums {
			nums2 = append(nums2, v)
		}
		nums2 = append(nums2, i)
		createNum(nums2, k, n, res)
	}
}

//想法就错了。
func findLengthOfShortestSubarray2(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	//思路：采用动态规划呢？dp[i]代表以自己为结尾的最大连续非递减子串的长度
	dp, max := make([]int, len(arr)), 1 //最大的子串长度
	dp[0] = 1
	for i := 1; i < len(arr); i++ {
		//dp[i] dp[i-j]+1
		//找到前面比arr[i]小的数
		j := i - 1
		for ; j >= 0 && arr[j] > arr[i]; j-- {
		}
		if j < 0 {
			dp[i] = 1
		} else {
			dp[i] = dp[j] + 1
			if dp[i] > max {
				max = dp[i] //找最大的子串长度
			}
		}
	}
	return len(arr) - max
}

//对角线打印
func findDiagonalOrder(nums [][]int) []int {
	res := []int{}

	maxlen := 0

	//暴力法
	for i := 0; i < len(nums); i++ {
		for j := 0; j <= i; j++ {
			if j < len(nums[i-j]) {
				res = append(res, nums[i-j][j])
			}
		}
		if len(nums[i]) > maxlen {
			maxlen = len(nums[i])
		}
	}
	//右下部分
	for i := len(nums) - 1; i > len(nums)-1-maxlen; i-- {
		for j := 1; j < maxlen; j++ {
			if i-j+1 < len(nums) && i-j+1 >= 0 && j < len(nums[i-j+1]) {
				res = append(res, nums[i-j+1][j])
			}
		}
	}

	return res
}

func main() {
	//fmt.Println(combine(4, 4))
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(findDiagonalOrder(arr))
}
