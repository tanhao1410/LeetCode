package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
	fmt.Println(maximumGap([]int{3, 4, 5, 8, 12, 12, 3, 4, 45, 3, 6, 9, 1}))
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
	//采用桶排序的方式，每个桶中记录两个数，最大的和最小的，然后比较前后两个桶之间的差距，
	//需要多少个桶？每个桶中存放的范围？
	//需要 nums + 1 )/2个，即每个桶平均可以放两个
	//范围的话，先求最大值

	//桶的第一个数为最大值，第二个数为最小值。

	max, min := 0, math.MaxInt32
	for _, v := range nums {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	dis := max - min
	if dis <= 0 {
		return 0
	}
	//桶的数量,当桶的数量够多的情况下，桶内部的数据之间的距离不可能是最大的了。
	tCount := len(nums)
	ts := make([][]int, tCount)

	//区间大小
	qLen := dis/tCount + 1

	for _, v := range nums {
		//应该放在第几个桶里面
		tNum := (v - min) / qLen
		//还没有桶，生成桶，放进数
		if ts[tNum] == nil {
			t := []int{v, v}
			ts[tNum] = t
		} else {
			t := ts[tNum]
			if v > t[1] {
				if t[1] < t[0] {
					t[0] = t[1]
				}
				t[1] = v
			} else if v < t[0] {
				t[0] = v
			}
		}
	}

	res := 0
	preMax := math.MaxInt32
	for _, v := range ts {
		if v != nil {
			if v[0]-preMax > res {
				res = v[0] - preMax
			}
			preMax = v[1]
		}
	}

	return res
}
