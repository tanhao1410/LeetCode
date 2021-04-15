package main

import "fmt"

func main() {
	fmt.Print(countDigitOne(45))
}

//233. 数字 1 的个数
func countDigitOne(n int) int {

	res := 0
	if n == 0 {
		return 0
	}
	if n < 10 {
		return 1
	}

	tenN := func(n int) int {
		res := 1
		for ; n > 0; n-- {
			res *= 10
		}
		return res
	}

	dp := make([]int, 12)
	dp[1] = 1
	for i := 1; i < 12; i++ {
		dp[i] = 10*dp[i-1] + tenN(i-1)
	}

	//先求高位
	getMaxBit := func(n int) int {
		for n > 9 {
			n /= 10
			if n < 10 {
				return n
			}
		}
		return n
	}

	//求位数
	getBitNum := func(n int) int {
		res := 1
		for ; n > 9; res++ {
			n /= 10
		}
		return res
	}

	bitNum := getBitNum(n)

	res += dp[bitNum-1] //比这个数小一位的所有1的数量。

	maxBit := getMaxBit(n)

	if maxBit > 1 {
		res += tenN(bitNum - 1)
		res += (maxBit - 1) * dp[bitNum-1]
	} else {
		res += n%tenN(bitNum-1) + 1
	}

	res += countDigitOne(n % tenN(bitNum-1))

	return res
}

//162. 寻找峰值
func findPeakElement(nums []int) int {
	//最简单思路：遍历一遍
	//
	// for i:=0;i < len(nums);i ++{
	//     if (i - 1 < 0 || nums[i - 1] < nums[i]) && (i + 1 >= len(nums) || nums[i + 1] < nums[i]){
	//         return i
	//     }
	// }
	// return -1

	//二分思路，从中间开始找，如果是峰，返回，如果是谷，则前面或后面肯定有个峰，如果递增，则后面有
	start, end := 0, len(nums)-1
	middle := (end + start) / 2
	for ; ; middle = (start + end) / 2 {

		//区间的数字不超过两个时
		if middle == start {
			if nums[end] > nums[middle] {
				return end
			} else {
				return middle
			}
		}

		if nums[middle] > nums[middle-1] && nums[middle] > nums[middle+1] {
			return middle
		}

		//递增
		if nums[middle] > nums[middle-1] && nums[middle+1] > nums[middle] {
			start = middle + 1
		} else {
			end = middle - 1
		}

	}
}

//213. 打家劫舍 II
func rob(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}

	//动态规划算法：+ 用一个数组额外记录，开始的地方
	startIndex := make([]int, len(nums))
	dp := make([]int, len(nums))
	dp1 := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {

		if i == 1 {

			if nums[1] > nums[0] {
				dp[1] = nums[1]
				startIndex[1] = 1
			} else {
				dp[1] = nums[0]
			}

			dp1[1] = nums[1]

		} else if i == len(nums)-1 {

			//最后一个特殊处理
			if nums[i]+dp[i-2] > dp[i-1] {
				if startIndex[i-2] == 0 {

					if dp1[i-2]+nums[i] > dp[i-1] {
						//要最后一个数，不要第一个数了
						dp[i] = dp1[i-2] + nums[i]
					} else {
						dp[i] = dp[i-1]
					}

				} else {

					dp[i] = dp[i-2] + nums[i]
				}
			} else {
				dp[i] = dp[i-1]
			}

		} else if i > 1 {

			if nums[i]+dp[i-2] > dp[i-1] {
				dp[i] = nums[i] + dp[i-2]
				startIndex[i] = startIndex[i-2]
			} else {
				dp[i] = dp[i-1]
				startIndex[i] = startIndex[i-1]
			}

			//dp1是不能选第一个数的
			if i == 2 {
				if nums[2] > nums[1] {
					dp1[2] = nums[2]
				} else {
					dp1[2] = dp1[1]
				}
			} else {
				if nums[i]+dp1[i-2] > dp1[i-1] {
					dp1[i] = nums[i] + dp1[i-2]
				} else {
					dp1[i] = dp1[i-1]
				}
			}

		}
	}

	return dp[len(nums)-1]
}
