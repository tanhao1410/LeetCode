package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 2, 4, 4, 3}
	fmt.Println(candy(nums))
}

//423. 从英文中重建数字
func originalDigits(s string) string {
	//思路：先统计各个字母的次数，拼出对应数字的大小
	//one two three four five six seven eight nine  zero
	//w的数量代表2的数量。u代表4的数量。4 已知的话，根据f可以求出5的量，x代表6，z可以代表0
	//每一个字母对应的数量
	m := make([]int, 26)
	for _, v := range s {
		m[v-'a']++
	}
	//0，2,3,4,5,6,8已知数量。
	numCount := make([]int, 10)
	numCount[0] = m['z'-'a']
	numCount[2] = m['w'-'a']
	numCount[4] = m['u'-'a']
	numCount[6] = m['x'-'a']
	numCount[5] = m['f'-'a'] - numCount[4]
	numCount[7] = m['v'-'a'] - numCount[5]
	numCount[3] = m['r'-'a'] - numCount[4] - numCount[0]
	numCount[8] = m['t'-'a'] - numCount[2] - numCount[3]
	numCount[1] = m['o'-'a'] - numCount[2] - numCount[0] - numCount[4]
	numCount[9] = m['i'-'a'] - numCount[5] - numCount[6] - numCount[8]
	res := []byte{}
	for i := 0; i < 10; i++ {
		for j := 0; j < numCount[i]; j++ {
			res = append(res, byte(i+'a'))
		}
	}
	return string(res)
}

//每日一题：135. 分发糖果
func candy(ratings []int) int {
	if len(ratings) < 2 {
		return len(ratings)
	}
	//思路：
	dp := make([]int, len(ratings))

	count := 0
	//如果一个数比两边都小，则它分配1个
	for i := 1; i < len(ratings)-1; i++ {
		if ratings[i] <= ratings[i-1] && ratings[i] <= ratings[i+1] {
			dp[i] = 1
			count++
		}
	}
	if ratings[1] >= ratings[0] {
		dp[0] = 1
		count++
	}
	if ratings[len(ratings)-1] <= ratings[len(ratings)-2] {
		dp[len(ratings)-1] = 1
		count++
	}

	for count < len(ratings) {
		//先给分数最小的分配一个，存在多个最小的key
		min := math.MaxInt32
		for k, v := range ratings {
			if dp[k] == 0 && v < min {
				min = v
			}
		}
		for k, v := range ratings {
			if v == min && dp[k] == 0 {
				//dp[k] = 1
				if (k+1 < len(ratings) && dp[k+1] != 0) || (k-1 > -1 && dp[k-1] != 0) {
					//它的两边已经有分配了
					//看是否有相等的，如果有相等的，先按相等的来
					if k+1 < len(ratings) && dp[k+1] != 0 && ratings[k+1] == ratings[k] {
						dp[k] = dp[k+1]
						if k-1 > -1 && ratings[k] > ratings[k-1] && (dp[k] <= dp[k-1] || dp[k] > dp[k-1]+1) {
							dp[k] = dp[k-1] + 1
						}
					} else if k-1 > -1 && dp[k-1] != 0 && ratings[k-1] == ratings[k] {
						dp[k] = dp[k-1]
						if k+1 < len(ratings) && ratings[k] > ratings[k+1] && (dp[k] <= dp[k+1] || dp[k] > dp[k+1]+1) {
							dp[k] = dp[k+1] + 1
						}
					} else {
						value := 0
						if k+1 < len(ratings) {
							value = dp[k+1]
						}
						if k-1 > -1 && dp[k-1] > value {
							value = dp[k-1]
						}
						dp[k] = value + 1
					}
					count++
				} else {
					dp[k] = 1
					count++
				}
			}
		}
	}

	res := 0
	for _, v := range dp {
		res += v
	}
	return res
}
