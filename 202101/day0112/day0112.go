package main

func main() {

}

//231. 2的幂
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	for ; ; n /= 2 {
		if n == 1 {
			return true
		}
		if n%2 != 0 {
			return false
		}
	}
}
