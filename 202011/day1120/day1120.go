package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//剑指 Offer 25. 合并两个排序的链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var p *ListNode
	p1, p2 := l1, l2
	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			if head == nil {
				head, p = p2, p2
			} else {
				p.Next = p2
				p = p2
			}
			p2 = p2.Next
		} else {
			if head == nil {
				head, p = p1, p1
			} else {
				p.Next = p1
				p = p1
			}
			p1 = p1.Next
		}
	}
	if p1 != nil {
		if p == nil {
			return p1
		}
		p.Next = p1
	}
	if p2 != nil {
		if p == nil {
			return p2
		}
		p.Next = p2
	}

	return head
}

//面试题 03.05. 栈排序
type SortedStack struct {
	Stack []int
}

func Constructor2() SortedStack {
	return SortedStack{
		Stack: []int{},
	}
}

func (this *SortedStack) Push(val int) {
	if len(this.Stack) == 0 || this.Peek() >= val {
		this.Stack = append(this.Stack, val)
	} else {
		//需要将该数插入到合适的位置
		this.Stack = append(this.Stack, 0)
		index := len(this.Stack) - 2 //原来的倒数第一个数
		for ; index >= 0 && this.Stack[index] < val; index-- {
			this.Stack[index+1] = this.Stack[index]
		}
		this.Stack[index+1] = val
	}
}

func (this *SortedStack) Pop() {
	if len(this.Stack) > 0 {
		this.Stack = this.Stack[:len(this.Stack)-1]
	}
}

func (this *SortedStack) Peek() int {
	res := -1
	if len(this.Stack) > 0 {
		res = this.Stack[len(this.Stack)-1]
	}
	return res
}

func (this *SortedStack) IsEmpty() bool {
	return len(this.Stack) == 0
}

//面试题 03.03. 堆盘子
type StackOfPlates struct {
	Cap    int
	Stacks [][]int
}

func Constructor(cap int) StackOfPlates {
	return StackOfPlates{
		Cap:    cap,
		Stacks: [][]int{},
	}
}

func (this *StackOfPlates) Push(val int) {
	if this.Cap <= 0 {
		return
	}
	plateCount := len(this.Stacks)
	//新加入的数应该放进去总后一个盘子中
	//可能为0
	if plateCount == 0 || len(this.Stacks[plateCount-1]) == this.Cap {
		newPlate := []int{val}
		this.Stacks = append(this.Stacks, newPlate)
	} else {
		this.Stacks[plateCount-1] = append(this.Stacks[plateCount-1], val)
	}
}

func (this *StackOfPlates) Pop() int {
	//从最后一个盘子中弹出数据
	plateCount := len(this.Stacks)
	if plateCount == 0 || this.Cap <= 0 {
		return -1
	}
	res := this.Stacks[plateCount-1][len(this.Stacks[plateCount-1])-1]
	if len(this.Stacks[plateCount-1]) == 1 {
		this.Stacks = this.Stacks[:plateCount-1]
	} else {
		this.Stacks[plateCount-1] = this.Stacks[plateCount-1][:len(this.Stacks[plateCount-1])-1]
	}
	return res
}

func (this *StackOfPlates) PopAt(index int) int {
	plateCount := len(this.Stacks)
	//指定盘子不存在
	if index >= plateCount || this.Cap <= 0 {
		return -1
	}
	curPlate := this.Stacks[index]
	res := curPlate[len(curPlate)-1]
	if len(curPlate) == 1 {
		//如果只有一个的话，会删除此盘子
		tail := this.Stacks[index+1:]
		pre := this.Stacks[:index]
		this.Stacks = append(pre, tail...)
	} else {
		//只是删除此盘子中的这个数
		this.Stacks[index] = this.Stacks[index][:len(curPlate)-1]
	}
	return res
}
