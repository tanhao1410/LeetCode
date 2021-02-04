package main

import "fmt"

func main() {
	for i := 1; i < 1000; i++ {
		fmt.Print(i, ",")
	}
}

//873. 最长的斐波那契子序列的长度-时间超时。。。
func lenLongestFibSubseq2(arr []int) int {
	//
	res := 0
	for i := 0; i < len(arr)-2; i++ {

		//后面的数不用考虑了
		if len(arr)-i < res {
			break
		}
		for j := i + 1; j < len(arr)-1; j++ {

			size := 2
			fir, sec := arr[i], arr[j]
			for k := i + 2; k < len(arr); k++ {
				if arr[k] > fir+sec {
					break
				}
				if arr[k] == fir+sec {
					size++
					fir, sec = sec, arr[k]
				}
			}

			if size > 2 && size > res {
				res = size
			}
		}
	}

	return res
}

//每日一题：643. 子数组最大平均数 I
func findMaxAverage(nums []int, k int) float64 {

	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	res := sum

	for i := k; i < len(nums); i++ {

		sum = sum - nums[i-k] + nums[i]
		if sum > res {
			res = sum
		}

	}

	return float64(res) / float64(k)
}
