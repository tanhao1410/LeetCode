package main

import "fmt"

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}

//215. 数组中的第K个最大元素
func findKthLargest(nums []int, k int) int {
	//以第一个数将两边的数隔开，前面是比它小的或相等的，后面是比它大的。
	//如果前面的数刚好有 k - 1个，那么，这个中间的数就是结果。否则
	//递归下去。
	first := nums[0]
	i, j := 0, len(nums)-1
	for i < j {
		for ; i < len(nums) && nums[i] <= first; i++ {
		}
		for ; j >= 0 && nums[j] > first; j-- {
		}
		//交换
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[0], nums[j] = nums[j], first
	//找的是k个最大元素
	if len(nums)-j == k {
		return nums[j]
	} else if len(nums)-j < k {
		return findKthLargest(nums[:j], k+j-len(nums))
	}
	return findKthLargest(nums[j+1:], k)
}

//每日一题：204. 计数质数-埃氏筛法
func countPrimes(n int) int {
	res := 0
	primes := make([]bool, n)
	for i := 2; i < n; i++ {
		if !primes[i] {
			res += 1
			for j := 2; i*j < n; j++ {
				primes[i*j] = true
			}
		}
	}
	return res
}

//每日一题：204. 计数质数。时间超时
func countPrimes2(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		if isPrimes(i) {
			res += 1
		}
	}
	return res
}

//判断一个数是否是质数
func isPrimes(n int) bool {
	for i := 2; i <= n; i++ {
		if i*i > n {
			return true
		}
		if n%i == 0 {
			return false
		}
	}
	return false
}
