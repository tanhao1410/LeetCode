package main

func main() {
	//
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
