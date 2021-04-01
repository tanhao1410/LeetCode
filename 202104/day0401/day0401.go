package main

func main() {

}

//每日一题：1006. 笨阶乘
func clumsy(N int) int {
	//思路：按照*/+-顺序来
	//共有 几组 N - 1 / 4 组，余 N - 1 % 4项
	//先算*/,以
	if N == 2 {
		return 2
	} else if N == 1 {
		return 1
	} else if N == 3 {
		return 6
	}
	res := N*(N-1)/(N-2) + N - 3
	//用for循环来做
	for N = N - 4; N >= 3; N -= 4 {
		res -= N * (N - 1) / (N - 2)
		//这个是加上
		res += N - 3
	}
	if N > 0 {
		res -= N
	}
	return res
}
