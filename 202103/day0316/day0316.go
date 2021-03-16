package main

func main() {
	//
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
