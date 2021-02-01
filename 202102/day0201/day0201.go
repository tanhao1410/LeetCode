package main

func main() {

}

//每日一题：888. 公平的糖果棒交换
func fairCandySwap(A []int, B []int) []int {
	res := []int{}
	//只是交换一根，思路：先求和，
	sumA, sumB := 0, 0
	for _, v := range A {
		sumA += v
	}
	m := make(map[int]bool)
	for _, v := range B {
		sumB += v
		m[v] = true
	}
	bigThen := (sumA + sumB) / 2
	//转换成了求两数之和
	for _, v := range A {
		if m[v-bigThen] {
			res = append(res, v, v-bigThen)
			return res
		}
	}

	return res
}
