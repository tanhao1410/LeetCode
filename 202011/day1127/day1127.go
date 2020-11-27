package main

func main() {

}

//剑指 Offer 09. 用两个栈实现队列
type CQueue struct {
	Stack1 []int
	Stack2 []int
	//思路，两个栈，进队时，直接进push进1，出队时，先将栈1全部pop到栈2，然后弹出栈2的顶部数据，剩余的再弹回栈1
}

func Constructor() CQueue {
	return CQueue{
		[]int{}, []int{},
	}
}

func (this *CQueue) AppendTail(value int) {
	this.Stack1 = append(this.Stack1, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.Stack1) == 0 && len(this.Stack2) == 0 {
		return -1
	}
	if len(this.Stack2) == 0 {
		for len(this.Stack1) > 0 {
			popNum := this.Stack1[len(this.Stack1)-1]
			this.Stack1 = this.Stack1[:len(this.Stack1)-1]
			this.Stack2 = append(this.Stack2, popNum)
		}
	}
	res := this.Stack2[len(this.Stack2)-1]
	this.Stack2 = this.Stack2[:len(this.Stack2)-1]
	return res
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
