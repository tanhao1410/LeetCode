package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isPossible([]int{1, 2, 3, 3, 4, 5}))
}

//329. 矩阵中的最长递增路径
func longestIncreasingPath(matrix [][]int) int {
	//思路：采用动态规划的算法，dp[i][j],代表以自身开头的最长的递增路径的长度
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	res := 1
	//初始化
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}

	//获取以某个点开始的最长的递增路径大小
	var getMaxDisPath func(i, j int) int
	getMaxDisPath = func(i, j int) int {
		if dp[i][j] != 0 {
			return dp[i][j]
		}
		max := 1
		//可以向下
		if i+1 < len(matrix) {
			if matrix[i][j] < matrix[i+1][j] {
				down := getMaxDisPath(i+1, j)
				if down+1 > max {
					max = down + 1
				}
			}
		}
		//可以向右
		if j+1 < len(matrix[0]) {
			if matrix[i][j] < matrix[i][j+1] {
				right := getMaxDisPath(i, j+1)
				if right+1 > max {
					max = right + 1
				}
			}
		}
		//可以向左
		if j-1 >= 0 {
			if matrix[i][j] < matrix[i][j-1] {
				left := getMaxDisPath(i, j-1)
				if left+1 > max {
					max = left + 1
				}
			}
		}
		//可以向上
		if i-1 >= 0 {
			if matrix[i][j] < matrix[i-1][j] {
				up := getMaxDisPath(i-1, j)
				if up+1 > max {
					max = up + 1
				}
			}
		}
		dp[i][j] = max
		return dp[i][j]
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if res < getMaxDisPath(i, j) {
				res = getMaxDisPath(i, j)
			}
		}
	}
	return res
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
	//长度小于3
	if len(nums) < 3 {
		return false
	}
	//对于数组中的元素，如果存在一个子序列以x-1结尾，则可以将x加入到该子序列中
	m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}
	//mm := make(map[int]int)//key 为尾部元素，value为数量 // 不行，如果已经有了已key结尾的呢？
	//用一个二维来记录
	mm := [][]int{}
	//可以放在哪一个的尾部
	whichTail := func(num int) int {
		res := -1
		for i := 0; i < len(mm); i++ {
			//优先放在短的那边
			if mm[i][len(mm[i])-1] == num-1 {
				if res == -1 {
					res = i
				} else if len(mm[i]) < len(mm[res]) {
					res = i
				}
			}
		}
		return res
	}
	for _, v := range nums {
		index := whichTail(v)
		if index != -1 {
			mm[index] = append(mm[index], v)
		} else {
			mm = append(mm, []int{v})
		}
	}
	for _, v := range mm {
		if len(v) < 3 {
			return false
		}
	}
	return true
}
