package main

func main() {

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
