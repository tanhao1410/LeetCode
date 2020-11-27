package main

func main() {

}

//275. H 指数 II
func hIndex(citations []int) int {
	//已经是排好序的了
	for i := len(citations) - 1; i >= 0; i-- {
		//至少多少篇被引用了
		count := len(citations) - i
		hNum := citations[i]
		if hNum == count {
			return count
		} else if count > hNum {
			//即引用的人数高于当前访问的这个文献的访问次数了，即再往下走，没有意义了。而前面都是符合的，不然也不会走到这。
			return count - 1
		}
	}
	return len(citations)
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
