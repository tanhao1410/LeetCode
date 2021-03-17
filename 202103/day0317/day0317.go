package main

func main() {
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
