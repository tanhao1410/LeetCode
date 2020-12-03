package main

func main() {

}

//每日一题：204. 计数质数-埃氏筛法
func countPrimes(n int) int {
	res := 0
	primes := make([]bool, n)
	for i := 2; i < n; i++ {
		if !primes[i] {
			res += 1
			for j := 2; i*j < n; j++ {
				primes[i*j] = true
			}
		}
	}
	return res
}

//每日一题：204. 计数质数。时间超时
func countPrimes2(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		if isPrimes(i) {
			res += 1
		}
	}
	return res
}

//判断一个数是否是质数
func isPrimes(n int) bool {
	for i := 2; i <= n; i++ {
		if i*i > n {
			return true
		}
		if n%i == 0 {
			return false
		}
	}
	return false
}
