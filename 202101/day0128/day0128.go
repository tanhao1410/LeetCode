package main

import "fmt"

func main() {
	fmt.Println(pyramidTransition("BCD", []string{"BCG", "CDE", "GEA", "FFF"}))

	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6, 4, 6, 77, 7, 100, 12, 111, 12, 12, 3, 1, 6}, 100))
}

//713. 乘积小于K的子数组
func numSubarrayProductLessThanK(nums []int, k int) int {
	res := 0

	dp := make([]int, len(nums))
	dp2 := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] >= k {
			dp[i] = -1
			dp2[i] = -1
			continue
		}
		//前面没有的情况，即前一个数大于k或没有数
		if i < 1 || dp[i-1] == -1 {
			j := i + 1
			num := nums[i]
			for ; j < len(nums) && num*nums[j] < k; j++ {
				num *= nums[j]
			}
			dp[i] = j - 1
			dp2[i] = num
		} else {

			num := 0
			j := dp[i-1] + 1

			if dp[i-1] >= i {
				num = dp2[i-1] / nums[i-1]
			} else {
				num = nums[i]
				j = i + 1
			}
			for ; j < len(nums) && num*nums[j] < k; j++ {
				num *= nums[j]
			}
			dp[i] = j - 1
			dp2[i] = num
		}
	}
	for i := 0; i < len(nums); i++ {
		if dp[i] >= i {
			res += (dp[i] - i + 1)
		}
	}
	return res
}

//756. 金字塔转换矩阵
func pyramidTransition(bottom string, allowed []string) bool {

	m := make(map[string]bool)
	for _, s := range allowed {
		m[s] = true
	}
	//是否存在以某两个字母开头的
	m2 := make(map[string][]string)
	for _, s := range allowed {
		if _, ok := m2[s[:2]]; ok {
			m2[s[:2]] = append(m2[s[:2]], s[2:])
		} else {
			m2[s[:2]] = []string{s[2:]}
		}
	}
	//可以用递归的形式，找出第二层，直到找到首层

	//根据底层推断出上一层可以的
	if len(bottom) == 2 {
		//如果能找到，就返回
		_, ok := m2[bottom]
		return ok
	}

	//上一层各位置可以选择的字母
	preCan := make([][]string, len(bottom)-1)

	for i := 1; i < len(bottom); i++ {
		if letters, ok := m2[bottom[i-1:i+1]]; ok {
			preCan[i-1] = letters
		} else {
			return false
		}
	}

	//组合成可以的上一层
	preStrs := []string{}

	var getPreCanStr func(index int, preStr string)

	getPreCanStr = func(index int, preStr string) {

		if index == len(preCan)-1 {
			for i := 0; i < len(preCan[index]); i++ {
				if len(preStr) > 1 {
					if len(m2[preStr[len(preStr)-2:]]) > 0 {
						preStrs = append(preStrs, preStr+preCan[index][i])
					}
				} else {
					preStrs = append(preStrs, preStr+preCan[index][i])
				}
			}
			return
		}

		for i := 0; i < len(preCan[index]); i++ {
			s := preStr + preCan[index][i]
			if len(preStr) >= 2 {
				if len(m2[preStr[len(preStr)-2:]]) > 0 {
					getPreCanStr(index+1, s)
				}
			} else {
				getPreCanStr(index+1, s)
			}
		}
	}

	getPreCanStr(0, "")

	//可以组成的集合
	for i := 0; i < len(preStrs); i++ {
		if pyramidTransition(preStrs[i], allowed) {
			return true
		}
	}
	return false
}

//每日一题：724. 寻找数组的中心索引
func pivotIndex(nums []int) int {
	//先求总和，然后从左边，依次求和，若和+下一个数==总和-所有，则找到了
	sum := 0
	for _, v := range nums {
		sum += v
	}

	preSum := 0
	for i := 0; i < len(nums); i++ {
		//前面的数的和
		if i > 0 {
			preSum += nums[i-1]
		}
		//后面数的和
		tailSum := sum - preSum - nums[i]
		if tailSum == preSum {
			return i
		}
	}
	return -1
}
