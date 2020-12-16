package main

import (
	"strings"
)

func main() {

}

//397. 整数替换
func integerReplacement(n int) int {
	res := 0
	//如果是偶数，可以直接减半，如果是奇数，可以选择+1，也可以选择-1，至于选择哪个？
	for n%2 == 0 {
		n = n / 2
		res++
	}
	//结束了
	if n == 1 {
		return res
	}
	////选择+1或-1
	res++
	plus := integerReplacement(n + 1)
	dec := integerReplacement(n - 1)
	if plus < dec {
		return plus + res
	}
	return res + dec
}

type NestedInteger struct{}

func (this NestedInteger) IsInteger() bool { return false }

func (this NestedInteger) GetInteger() int { return 0 }

func (n *NestedInteger) SetInteger(value int)      {}
func (this *NestedInteger) Add(elem NestedInteger) {}

func (this NestedInteger) GetList() []*NestedInteger { return nil }

//341. 扁平化嵌套列表迭代器
type NestedIterator struct {
	Data  []int
	point int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	res := NestedIterator{
		Data:  []int{},
		point: 0,
	}
	if len(nestedList) != 0 {
		stack := []*NestedInteger{}
		for i := len(nestedList) - 1; i >= 0; i-- {
			stack = append(stack, nestedList[i])
		}
		stackLen := len(stack)
		for stackLen > 0 {
			//处理栈顶的元素
			top := stack[stackLen-1]
			stack = stack[:len(stack)-1]
			if top.IsInteger() {
				res.Data = append(res.Data, top.GetInteger())
			} else {
				for i := len(top.GetList()) - 1; i >= 0; i-- {
					stack = append(stack, top.GetList()[i])
				}
			}
			stackLen = len(stack)
		}
	}
	return &res
}

func (this *NestedIterator) Next() int {
	this.point++
	return this.Data[this.point-1]
}

func (this *NestedIterator) HasNext() bool {
	return len(this.Data) > this.point
}

//每日一题：290. 单词规律
func wordPattern(pattern string, s string) bool {
	m := make(map[byte]string)
	ss := strings.Split(s, " ")
	if len(s) == 0 || len(pattern) != len(ss) {
		return false
	}
	for i := 0; i < len(ss); i++ {
		if pre, ok := m[pattern[i]]; ok {
			if pre != ss[i] {
				return false
			}
		} else {
			m[pattern[i]] = ss[i]
		}
	}
	m2 := make(map[string]bool)
	for _, v := range m {
		m2[v] = true
	}

	return len(m2) == len(m)
}
