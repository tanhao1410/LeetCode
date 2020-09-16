package main

import (
	"fmt"
)

func main() {
	nums := []int{
		24, 12, 71, 33, 5, 87, 10, 11, 3, 58, 2, 97, 97,
		36, 32, 35, 15, 80, 24, 45, 38, 9, 22, 21, 33, 68, 22, 85, 35, 83, 92, 38, 59, 90, 42, 64, 61, 15,
		4, 40, 50, 44, 54, 25, 34, 14, 33,
		94, 66, 27, 78, 56,
		3,
		29, 3, 51, 19, 5, 93, 21, 58, 91, 65, 87, 55, 70, 29, 81, 89, 67, 58, 29, 68, 84,
		4, 51, 87, 74, 42, 85, 81, 55, 8, 95, 39}
	//fmt.Println(PredictTheWinner(nums))
	fmt.Println(longestSubarray(nums, 87))
}

//最长连续子数组的长度    5
func longestSubarray(nums []int, limit int) int {
	//窗口开始为1，i向前走，若满足条件，窗口+1，否则，j++，看是否满足，不满足，继续j++。i向前走，直到结束
	res := 1
	for i, j := 1, 0; i < len(nums); i++ {
		//如果nums[i]符合要求，j与i之间的数都
		for k := j; k < i; k++ {
			if !IsTwoNumLimitTarget(nums[k], nums[i], limit) {
				//有不符合条件的，则j++ //不是j++ 而是j直接走到不符合位置之后,这样也不行//k+1和k相同。
				j = k + 1 //j=12
				continue  //问题出在这。j会++的
			}
		}
		if i-j+1 > res {
			res = i - j + 1
		}
	}
	return res
}

func longestSubarray2(nums []int, limit int) int {
	//窗口开始为1，i向前走，若满足条件，窗口+1，否则，j++，看是否满足，不满足，继续j++。i向前走，直到结束
	res := 1
	for i, j := 1, 0; i < len(nums); i++ {
		//如果nums[i]符合要求，j与i之间的数都
		for k := j; k < i; k++ {
			if !IsTwoNumLimitTarget(nums[k], nums[i], limit) {
				//有不符合条件的，则j++ //不是j++ 而是j直接走到不符合位置之后,这样也不行//k+1和k相同。
				j = k + 1 //j=12
				continue  //问题出在这。j会++的
			}
		}
		if i-j+1 > res {
			res = i - j + 1
		}
	}
	return res
}

func IsTwoNumLimitTarget(num1, num2, target int) bool {
	if num1 > num2 {
		return num1-num2 <= target
	}
	return num2-num1 <= target
}

//预测第一个玩家的输赢 [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15]
func PredictTheWinner(nums []int) bool {
	//思路：m[i]选了i之后会怎样？
	//递归思路：当只有一个数字时，两个数字时必赢，当有三个数字时，若 中间大于 两端之后，false,四个数字的话 相当于 必赢
	if len(nums)%2 == 0 || len(nums) == 1 {
		return true
	}

	//否则第一个先选一个，最后一个，或者第一个
	selectFirst := MoreThanSecond_Max(nums[1:]) - nums[0]
	selectLast := MoreThanSecond_Max(nums[:len(nums)-1]) - nums[len(nums)-1]

	return selectFirst <= 0 || selectLast <= 0
}

//新思路：当我挑选走一个后，就变成了2号先拿的 偶数形式了。也就是比较 2号先拿的话，拿到的比我多拿的一个还多

//当偶数个情况下，第一个先拿的比第二个最少多多少 2,3,4,5,6,7,8,9,10,11,12,13,14,15
func MoreThanSecond_Max(nums []int) int {
	if len(nums) == 2 && nums[0] > nums[1] {
		return nums[0] - nums[1]
	} else if len(nums) == 2 && nums[1] >= nums[0] {
		return nums[1] - nums[0]
	}

	//可以要里面任意一个，因此肯定要拿最大的，但最大的旁边的肯定是拿不到了

	///四个的时候，是拿最大的，和
	//选一个，然后对方选一个，就变成两个了
	selectFirst := nums[0] - nums[len(nums)-1] + MoreThanSecond_Max(nums[1:len(nums)-1])
	selectFirst2 := nums[0] - nums[1] + MoreThanSecond_Max(nums[2:])

	//如果我选了第一个，对方肯定要从中选择最大的收益
	max := 0
	if selectFirst > selectFirst2 {
		max = selectFirst2
	} else {
		max = selectFirst
	}

	selectLast1 := nums[len(nums)-1] - nums[0] + MoreThanSecond_Max(nums[1:len(nums)-1])
	selectLast2 := nums[len(nums)-1] - nums[len(nums)-2] + MoreThanSecond_Max(nums[:len(nums)-2])
	if selectLast1 > selectLast2 {
		//最小的都大于之前的，
		if selectLast2 > max {
			return selectLast2
		} else {
			return max
		}
	} else {
		if selectLast1 > max {
			return selectLast1
		} else {
			return max
		}
	}
}

//递归的话，掐头去尾，中间可以相当于一个，第一个人应该拿的数。如果是偶数，必胜。 0 3 2 0
//考虑奇数的情况。 1 20 3 20 5
//只有这一种输的情况，中间的大于两边之和，
//
func TakeMiddleSum(nums []int) int {
	if len(nums) == 3 {
		return nums[1]
	}

	return TakeMiddleSum(nums[1 : len(nums)-1])
}
