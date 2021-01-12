package main

func main() {

}

//593. 有效的正方形
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {

	//思路：四个点，共可以形成6条连线，如果其中有四条是相等的，另外两条也相等
	getLineLen2 := func(p1 []int, p2 []int) int {
		return (p2[0]-p1[0])*(p2[0]-p1[0]) + (p2[1]-p1[1])*(p2[1]-p1[1])
	}

	//找两个相等的和四个相等的
	m := make(map[int]bool)
	m[getLineLen2(p1, p2)] = true
	m[getLineLen2(p1, p3)] = true
	m[getLineLen2(p1, p4)] = true
	m[getLineLen2(p2, p3)] = true
	m[getLineLen2(p2, p4)] = true
	m[getLineLen2(p3, p4)] = true

	if len(m) != 2 {
		return false
	}

	one := -1
	for k, _ := range m {
		if one == -1 {
			one = k
		} else {
			if one < k {
				return k == 2*one
			} else {
				return one == 2*k
			}
		}

	}

	return false
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

//326. 3的幂
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for ; ; n /= 3 {
		if n == 1 {
			return true
		}
		if n%3 != 0 {
			return false
		}
	}
}
