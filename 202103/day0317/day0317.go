package main

func main() {
}

//面试题 05.07. 配对交换
func exchangeBits(num int) int {
	bytes := make([]int8, 32)
	for i := 0; i < 32; i++ {

		if (num & (1 << (31 - i))) != 0 {
			if i%2 == 0 {
				bytes[i+1] = 1
			} else {
				bytes[i-1] = 1
			}

		}

	}

	var res int32 = 0
	for i := 0; i < 32; i++ {
		res <<= 1
		if bytes[i] != 0 {
			res += 1
		}
	}

	return int(res)
}

//面试题 08.09. 括号
func generateParenthesis(n int) []string {
	//n最大为11
	res := []string{}

	var next func(leftCount, rightCount int, str string)
	next = func(leftCount, rightCount int, str string) {
		//生成了合法的括号
		if leftCount == 0 && rightCount == 0 {
			res = append(res, str)
			return
		}

		//左括号比右括号少，所以，有两种放置方法
		if leftCount > 0 && leftCount < rightCount {
			next(leftCount-1, rightCount, str+"(")
			next(leftCount, rightCount-1, str+")")
		} else if leftCount == 0 {
			next(leftCount, rightCount-1, str+")")
		} else if rightCount == leftCount {
			next(leftCount-1, rightCount, str+"(")
		}
	}

	next(n-1, n, "(")
	return res
}

//面试题 08.11. 硬币
func waysToChange(n int) int {

	//level代表后面选择的钱要至少大于等于level
	var wayToChange2 func(n, level int) int
	wayToChange2 = func(n, level int) int {
		if n < 0 {
			return 0
		}
		if n == 0 || n == level {
			return 1
		}
		if level == 1 {
			return wayToChange2(n-1, 1) +
				wayToChange2(n-5, 5) +
				wayToChange2(n-10, 10) +
				wayToChange2(n-25, 25)
		} else if level == 5 {
			return wayToChange2(n-5, 5) +
				wayToChange2(n-10, 10) +
				wayToChange2(n-25, 25)
		} else if level == 10 {
			return wayToChange2(n-10, 10) +
				wayToChange2(n-25, 25)
		}
		return wayToChange2(n-25, 25)
	}

	return wayToChange2(n, 1)
}

//面试题 08.11. 硬币
func waysToChange2(n int) int {
	//优化：f(n,1)=f(n-1,1)+f(n-5,5)+f(n-10,10)+f(n-25,25)

	twoFive := func(num int) int {
		if num < 0 {
			return 0
		}
		if num%25 == 0 {
			return 1
		}
		return 0
	}
	//动态规划算法
	ten := make([]int, n+1)
	ten[0] = 1
	for i := 10; i < n+1; i += 5 {
		ten[i] = ten[i-10]
		ten[i] += twoFive(i - 25)
	}

	five := make([]int, n+1)
	five[0] = 1
	for i := 5; i < n+1; i += 5 {
		five[i] += five[i-5]
		if i >= 10 {
			five[i] += ten[i-10]
		}
		five[i] += twoFive(i - 25)

		if five[i] > 1000000007 {
			five[i] = five[i] % 1000000007
		}
	}

	one := make([]int, n+1)
	one[0] = 1
	for i := 1; i < n+1; i++ {
		one[i] += one[i-1]
		if i >= 5 {
			one[i] += five[i-5]
		}
		if i >= 10 {
			one[i] += ten[i-10]
		}
		one[i] += twoFive(i - 25)

		if one[i] > 1000000007 {
			one[i] = one[i] % 1000000007
		}

	}
	return one[n]
}
