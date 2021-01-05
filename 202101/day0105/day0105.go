package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(largeGroupPositions("bbaaabbbc"))
	for i := math.MaxInt32; i >= math.MaxInt32-1000; i-- {
		fmt.Println(findClosedNumbers(i))
	}
}

//面试题 03.01. 三合一
type TripleInOne struct {
	Nums  []int
	One   int
	Two   int
	Three int
}

func Constructor(stackSize int) TripleInOne {
	return TripleInOne{
		Nums:  make([]int, stackSize*3),
		One:   -1,
		Two:   -1 + stackSize,
		Three: -1 + 2*stackSize,
	}
}

func (this *TripleInOne) Push(stackNum int, value int) {
	size := len(this.Nums) / 3
	switch stackNum {
	case 0:
		if this.One < size-1 {
			this.Nums[this.One+1] = value
			this.One++
		}
	case 1:
		if this.Two < size-1+size {
			this.Nums[this.Two+1] = value
			this.Two++
		}
	case 2:
		if this.Three < size-1+2*size {
			this.Nums[this.Three+1] = value
			this.Three++
		}
	}
}

func (this *TripleInOne) Pop(stackNum int) int {
	size := len(this.Nums) / 3
	switch stackNum {
	case 0:
		if this.One > -1 {
			this.One--
			return this.Nums[this.One+1]
		}
	case 1:
		if this.Two > -1+size {
			this.Two--
			return this.Nums[this.Two+1]
		}
	case 2:
		if this.Three > -1+2*size {
			this.Three--
			return this.Nums[this.Three+1]
		}
	}
	return -1
}

func (this *TripleInOne) Peek(stackNum int) int {
	size := len(this.Nums) / 3
	switch stackNum {
	case 0:
		if this.One > -1 {
			return this.Nums[this.One]
		}
	case 1:
		if this.Two > -1+size {
			return this.Nums[this.Two]
		}
	case 2:
		if this.Three > -1+2*size {
			return this.Nums[this.Three]
		}
	}
	return -1
}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
	size := len(this.Nums) / 3
	switch stackNum {
	case 0:
		if this.One > -1 {
			return false
		}
	case 1:
		if this.Two > -1+size {
			return false
		}
	case 2:
		if this.Three > -1+2*size {
			return false
		}
	}
	return true
}

//面试题 05.04. 下一个数
func findClosedNumbers(num int) []int {
	//求n的位数
	getOneCount := func(n int) int {
		res := 0
		for ; n != 0; n >>= 4 {
			v := n & 0xf
			if v == 1 || v == 2 || v == 4 || v == 8 {
				res += 1
			} else if v == 3 || v == 6 || v == 5 || v == 9 || v == 10 || v == 12 {
				res += 2
			} else if v == 7 || v == 11 || v == 14 || v == 13 {
				res += 3
			} else if v == 15 {
				res += 4
			}
		}
		return res
	}

	getMaxCount := func(n int) int {
		//最多31个1
		res := 31
		max := 0x80000000
		for ; n&max == 0 && max != 0; max >>= 1 {
			res--
		}
		return res
	}

	res := []int{}
	oneCount := getOneCount(num)
	if num == math.MaxInt32 {
		res = append(res, -1)
	} else {
		for i := num + 1; ; i++ {
			if getOneCount(i) == oneCount {
				res = append(res, i)
				break
			}
			if i == math.MaxInt32 {
				res = append(res, -1)
				break
			}
		}
	}
	//最多就那么多1的时候就要退出了，不一定要循环到0才结束。
	for i := num - 1; ; i-- {
		if getOneCount(i) == oneCount {
			return append(res, i)
		}
		//前面数最多能有多少个1，如果小于oneCount，那么直接退出即可。
		if getMaxCount(i) < oneCount {
			return append(res, -1)
		}
	}
}

//每日一题：830. 较大分组的位置
func largeGroupPositions(s string) [][]int {
	res := [][]int{}
	start, end := 0, 0
	for pre, i := s[0], 1; i < len(s); i++ {
		if s[i] == pre {
			end++
		} else {
			//遇到不相等的了
			if end-start >= 2 {
				res = append(res, []int{start, end})
			}
			start, end, pre = i, i, s[i]
		}
	}
	if end-start >= 2 {
		res = append(res, []int{start, end})
	}
	return res
}
