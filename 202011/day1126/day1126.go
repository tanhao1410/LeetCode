package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
}

//224. 基本计算器
func calculate(s string) int {
	//处理字符串
	ss := []string{}
	for i := 0; i < len(s); i++ {
		//去除前面的空格
		for ; i < len(s) && s[i] == ' '; i++ {
		}
		//看s[i]是否是+-（）
		if i < len(s) && (s[i] == '+' || s[i] == '-' || s[i] == '(' || s[i] == ')') {
			ss = append(ss, string(s[i]))
		} else if i < len(s) { //s[i]为0-9
			//截取数字串
			j := i + 1
			for ; j < len(s) && s[j] >= '0' && s[j] <= '9'; j++ {
			}
			ss = append(ss, string(s[i:j]))
			i = j - 1
		}
	}
	//简单求和，没有括号的
	simpleSum := func(stack []string) int {
		res, _ := strconv.Atoi(stack[0])
		//计算没有括号，只有+-的剩余的
		for i := 1; i < len(stack); i += 2 {
			nextNum, _ := strconv.Atoi(stack[i+1])
			if stack[i] == "+" {
				res += nextNum
			} else if stack[i] == "-" {
				res -= nextNum
			}
		}
		return res
	}
	stack := []string{}
	for i := 0; i < len(ss); i++ {
		//取出ss[i]
		cur := ss[i]
		if cur == ")" {
			j := len(stack) - 1
			for ; j >= 0 && stack[j] != "("; j-- {
			}
			sum := simpleSum(stack[j+1:])
			stack[j] = strconv.Itoa(sum)
			stack = stack[:j+1]
		} else {
			stack = append(stack, ss[i])
		}
	}
	return simpleSum(stack)
}

//剑指 Offer 39. 数组中出现次数超过一半的数字
func majorityElement(nums []int) int {
	//最简单的思路是记录每个数出现的次数
	res := nums[0]
	m := make(map[int]int)
	for _, v := range nums {
		if count, ok := m[v]; ok {
			m[v]++
			if count+1 > len(nums)/2 {
				return v
			}
		} else {
			m[v] = 1
		}
	}

	return res
}

//每日一题：164. 最大间距
func maximumGap(nums []int) int {

	return 0
}
