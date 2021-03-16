package main

import "fmt"

func main() {
	//
	fmt.Println(drawLine(6, 96, 0, 95, 1))
}

//面试题 05.08. 绘制直线
func drawLine(length int, w int, x1 int, x2 int, y int) []int {
	res := make([]int, length)
	//从哪一位开始
	start := y * (w / 32)
	//先组建显示行的数 字节数组
	bytes := make([]int, w)
	for i := x1; i <= x2; i++ {
		bytes[i] = 1
	}

	createNum := func(index int) int {
		var res int32 = 0
		for i := 0; i < 32; i++ {
			res <<= 1
			res += int32(1 & bytes[index+i])
		}
		return int(res)
	}

	//形成数
	for i := 0; i < w; i += 32 {
		res[start+i/32] = createNum(i)
	}

	return res
}

//面试题 05.06. 整数转换
func convertInteger(A int, B int) int {
	res := 0
	for i := 0; i < 32; i++ {
		if A&(1<<i) != B&(1<<i) {
			res++
		}
	}
	return res
}

//面试题 05.03. 翻转数位
func reverseBits(num int) int {

	bytes := make([]int, 32)
	//先把数字变成01
	for i := 0; i < 32; i++ {
		if num&(1<<i) > 0 {
			bytes[i] = 1
		}
	}

	//求最长的1的个数
	longOne := func() int {
		res := 0

		l := 0
		for i := 0; i < 32; i++ {
			if bytes[i] == 1 {
				l++
			} else {
				l = 0
			}
			if l > res {
				res = l
			}
		}

		return res
	}

	res := longOne()
	//找出改变一个1变成最长
	for i := 0; i < 32; i++ {
		if bytes[i] == 0 {
			bytes[i] = 1
			if longOne() > res {
				res = longOne()
			}
			bytes[i] = 0
		}
	}

	return res
}
