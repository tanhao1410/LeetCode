package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(splitIntoFibonacci("539834657215398346785398346991079669377161950407626991734534318677529701785098211336528511"))
}

//每日一题：842. 将数组拆分成斐波那契序列
func splitIntoFibonacci(S string) []int {
	res := []int{}
	//思路：先从一位开始算，然后找 下一个，用一位来处理，不行的话，二位，直到越界。然后，第一个数开始增加位数。如果可以的话就一直往前，失败了，位数增加一个，再重新开始。
	//判断一个字符串数是否越界
	isOver := func(s string) bool {
		max := "2147483647"
		if len(s) < 10 {
			//一个数不能以0开头
			if len(s) > 1 && s[0] == '0' {
				return false
			}
			return true
		}
		if len(s) == 10 {
			for i := 0; i < 10; i++ {
				if s[i] > max[i] {
					return false
				} else if s[i] < max[i] {
					return true
				}
			}
		}
		return false
	}

	tailTwoNumSumStr := func(nums []int) string {
		return strconv.Itoa(nums[len(nums)-1] + nums[len(nums)-2])
	}

	//第一个数先是一位
	for i := 1; i < len(S); i++ {
		for j := 1; i+j < len(S); j++ {
			//每一次重新设置首位和第二位时都要重置结果集
			res = []int{}

			//判断数是否越界。
			if !isOver(S[:i]) {
				return res
			}
			if !isOver(S[i : i+j]) {
				break
			}
			firstNum, _ := strconv.Atoi(S[:i])
			secondNum, _ := strconv.Atoi(S[i : i+j])
			sumStr := strconv.Itoa(firstNum + secondNum)
			if !isOver(sumStr) {
				break
			}
			//看后面还有数吗
			if len(S)-i-j == 0 || i+j+len(sumStr) > len(S) {
				//没有数了,或数不够用了
				break
			} else {
				//切割出来一个数
				nextNum := S[i+j : i+j+len(sumStr)]
				if sumStr == nextNum {
					//说明组成了
					res = append(res, firstNum, secondNum, firstNum+secondNum)
					//需要循环接下来的所有数看是否满足
					index := i + j + len(sumStr)
					for index < len(S) {
						//从结果集中拿出两个数之和
						sumStr := tailTwoNumSumStr(res)
						if !isOver(sumStr) {
							break
						}
						//在这里面还需要判断是否越界
						if index+len(sumStr) <= len(S) && sumStr == S[index:index+len(sumStr)] {
							newNum, _ := strconv.Atoi(sumStr)
							res = append(res, newNum)
							index = index + len(sumStr)
						} else {
							break
						}
					}
					if index == len(S) {
						return res
					}
				}
			}
		}
	}

	return res
}
