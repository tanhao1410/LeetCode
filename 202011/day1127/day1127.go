package main

func main() {

}

//每日一题：454. 四数相加 II
func fourSumCount(A []int, B []int, C []int, D []int) int {
	//因为具有相同的长度 N，且 0 ≤ N ≤ 500，可以采用双hash，变成两数之和
	res := 0
	//把cd组成一个hash
	m := make(map[int]int)
	n := len(A)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			m[C[i]+D[j]] += 1
		}
	}

	for _, v1 := range A {
		for _, v2 := range B {
			if v, ok := m[0-v1-v2]; ok {
				res += v
			}
		}
	}

	return res
}
