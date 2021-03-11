package main

import (
	"fmt"
	"strings"
)

func main() {
	oper := Constructor()
	fmt.Println(oper.Minus(2, 6))
	fmt.Println(oper.Divide(1909390, 3))
	fmt.Println(oper.Multiply(12345, -1234))
}

//面试题 01.09. 字符串轮转
func isFlipedString(s1 string, s2 string) bool {
	//拼接两个s2，看是否包含s1
	return strings.Contains(s2+s2, s1) && len(s2) == len(s1)
}

//面试题 16.11. 跳水板
func divingBoard(shorter int, longer int, k int) []int {
	if k == 0 {
		return nil
	}
	if shorter == longer {
		return []int{shorter * k}
	}
	res := make([]int, k+1)
	for i := 0; i <= k; i++ {
		res[i] = shorter*(k-i) + longer*i
	}

	return res
}

//面试题 16.10. 生存人数
func maxAliveYear(birth []int, death []int) int {
	//分别计算每年的人数，复杂度 年数 * 人数

	//遍历人数，给对应年份增加人数，复杂度不变。速度快些
	years := make([]int, 101)

	for i := 0; i < len(birth); i++ {
		for j := birth[i] - 1900; j <= death[i]-1900; j++ {
			years[j]++
		}
	}

	max := 0
	res := 0
	for i := 0; i < 101; i++ {
		if years[i] > max {
			res = i
			max = years[i]
		}
	}
	return res + 1900
}

//面试题 16.09. 运算
type Operations struct {
}

func Constructor() Operations {
	return Operations{}
}

func (this *Operations) Minus(a int, b int) int {
	//反码 + 1
	c := ^b + 1
	return a + c
}

func (this *Operations) Multiply(a int, b int) int {

	res := 0

	isZheng := (a > 0 && b > 0) || (a < 0 && b < 0)

	aa := a
	if a < 0 {
		aa = ^a + 1
	}
	bb := b
	if b < 0 {
		bb = ^b + 1
	}
	//用for循环来求
	if aa < bb {
		for i := 0; i < aa; i++ {
			res += bb
		}
	} else {
		for i := 0; i < bb; i++ {
			res += aa
		}
	}

	if isZheng {
		return res
	}
	return ^res + 1
}

func (this *Operations) Divide(a int, b int) int {

	res := 0

	isZheng := (a > 0 && b > 0) || (a < 0 && b < 0)

	aa := a
	if a < 0 {
		aa = ^a + 1
	}
	bb := b
	if b < 0 {
		bb = ^b + 1
	}
	//用for循环来求
	if aa < bb {
		return 0
	} else {

		for count := bb; count < aa; res++ {
			count += bb
		}
	}

	if isZheng {
		return res
	}
	return ^res + 1
}
