package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(Game24Points([]int{7, 1, 6, 2}))
	largestNumber([]int{1})
}

//每日一题：179. 最大数
func largestNumber(nums []int) string {
	//若都为0，直接返回0
	allZero := true
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		return "0"
	}

	sort.Slice(nums, func(i, j int) bool {
		//判断两者谁大谁小，最简单的方式就是先拼接，看谁在前比较大
		str1, str2 := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		parseInt, _ := strconv.ParseInt(str1+str2, 10, 64)
		parseInt2, _ := strconv.ParseInt(str2+str1, 10, 64)
		return parseInt > parseInt2
	})

	res := ""
	for i := 0; i < len(nums); i++ {
		res += strconv.Itoa(nums[i])
	}

	return res
}

/**
 *
 * @param arr int整型一维数组
 * @return bool布尔型
 */
func Game24Points(arr []int) bool {

	//计算两个数的表达式
	calTwo := func(calType int, num1, num2 int) int {
		if calType == 0 {
			return num1 + num2
		} else if calType == 1 {
			return num1 - num2
		} else if calType == 2 {
			return num1 * num2
		}
		return num1 / num2
	}

	//计算三个数表达式，先算乘除，再算加减 1 +  2 * 3
	calThree := func(type1, type2 int, one, two, three int) int {
		//如果首项是*/或尾是+-，先计算前面的
		if type1 >= 2 || type2 < 2 {
			return calTwo(type2, calTwo(type1, one, two), three)
		}
		return calTwo(type1, one, calTwo(type2, two, three))
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				if i >= 2 || j < 2 {
					if calThree(j, k, calTwo(i, arr[0], arr[1]), arr[2], arr[3]) == 24 {
						//找到了
						fmt.Printf("%d%d%d", i, j, k)
						return true
					}
				} else {
					if calThree(i, k, arr[0], calTwo(j, arr[1], arr[2]), arr[3]) == 24 {
						//找到了
						fmt.Printf("%d%d%d", i, j, k)
						return true
					}
				}
			}
		}
	}

	fmt.Println("false")
	return false
}
