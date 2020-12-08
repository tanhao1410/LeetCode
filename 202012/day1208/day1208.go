package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(splitIntoFibonacci("539834657215398346785398346991079669377161950407626991734534318677529701785098211336528511"))
	fmt.Println(restoreIpAddresses("101023"))
	fmt.Println(isInterleave(
		"aacaac",
		"aacaaeaac",
		"aacaacaaeaacaac"))
}

//474. 一和零。递归思路，时间超时。。。
func findMaxForm(strs []string, m int, n int) int {
	//动态规划做法：
	//递归思路：可以选择包括或不包括
	//结果包含i
	if len(strs) == 0 {
		return 0
	}
	zeroNum, oneNum := 0, 0
	for i := 0; i < len(strs[0]); i++ {
		if strs[0][i] == '0' {
			zeroNum++
		} else {
			oneNum++
		}
	}
	//这个肯定不能选
	if zeroNum > m || oneNum > n {
		return findMaxForm(strs[1:], m, n)
	}
	include := findMaxForm(strs[1:], m-zeroNum, n-oneNum) + 1
	noInclude := findMaxForm(strs[1:], m, n)
	if include > noInclude {
		return include
	}
	return noInclude
}

//97. 交错字符串
func isInterleave(s1 string, s2 string, s3 string) bool {
	//时间复杂度不行。
	//思路：
	//字母可以来自哪
	i, j, k := 0, 0, 0
	for (i < len(s1) || j < len(s2)) && k < len(s3) {
		//两边都可以取的情况。//用递归呢
		if i < len(s1) && j < len(s2) && s3[k] == s1[i] && s3[k] == s2[j] {
			//如果从A取可以的话，可以取几个，一次性放进去所有的

			//这样不行。
			l := 1
			for ; i+l < len(s1) && j+l < len(s2) && k+l < len(
				s3) && s3[k+l] == s1[i+l] && s3[k+l] == s2[j+l] && s1[i] == s1[i+l]; l++ {
			}

			if isInterleave(s1[i+l:], s2[j:], s3[k+l:]) {
				return true
			}
			return isInterleave(s1[i:], s2[j+l:], s3[k+l:])
		} else if i < len(s1) && s3[k] == s1[i] {
			k++
			i++
		} else if j < len(s2) && s3[k] == s2[j] {
			k++
			j++
		} else {
			return false
		}
	}
	return i == len(s1) && j == len(s2) && k == len(s3)
}

//93. 复原IP地址
func restoreIpAddresses(s string) []string {
	//思路：第一个可以是三个数字，也可以两个，或一个。如果是0开头的话，只能是一个数字
	can := func(s string) bool {
		if len(s) == 0 || len(s) > 3 || len(s) > 1 && s[0] == '0' {
			return false
		}
		num, _ := strconv.Atoi(s)
		return num <= 255
	}
	res := []string{}
	//第一个数字可以用几个
	for i := 1; i < 4; i++ {
		for j := 1; i < len(s) && can(s[:i]) && j < 4; j++ {
			for k := 1; i+j < len(s) && can(s[i:i+j]) && k < 4; k++ {
				//第三个确定的话，第四个也就确定了。
				if i+j+k < len(s) && can(s[i+j:i+j+k]) && can(s[i+j+k:]) {
					res = append(res, s[:i]+"."+s[i:i+j]+"."+s[i+j:i+j+k]+"."+s[i+j+k:])
				}
			}
		}
	}
	return res
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
