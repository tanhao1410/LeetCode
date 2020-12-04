package main

import "math"

func main() {

}

//313. 超级丑数
func nthSuperUglyNumber(n int, primes []int) int {
	//1是所有的数的超级丑数
	uglyNums := []int{1}
	primesCount := make([]int, len(primes))
	m := make(map[int]bool)
	for len(uglyNums) < n {
		min := math.MaxInt32
		minIndex := -1
		for k2, v2 := range primes {
			if v2*uglyNums[primesCount[k2]] < min {
				min = v2 * uglyNums[primesCount[k2]]
				minIndex = k2
			}
		}
		if !m[min] {
			uglyNums = append(uglyNums, min)
			m[min] = true
		}
		primesCount[minIndex] += 1
	}
	return uglyNums[n-1]
}

//307. 区域和检索 - 数组可修改
type NumArray struct {
	Nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		Nums: nums,
	}
}

func (this *NumArray) Update(i int, val int) {
	if len(this.Nums) > i {
		this.Nums[i] = val
	}
}

func (this *NumArray) SumRange(i int, j int) (res int) {
	for ; i < len(this.Nums) && i <= j; i++ {
		res += this.Nums[i]
	}
	return
}

//每日一题：659. 分割数组为连续子序列
func isPossible(nums []int) bool {
	//长度小于6肯定false
	//把重复的数字单独拿出来，尽量组成短的，最后剩下的拼接过去即可。

	return false
}
