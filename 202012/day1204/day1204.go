package main

func main() {

}

//307. 区域和检索 - 数组可修改
type NumArray struct {
	Nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		Nums: nums,
	}
}

func (this *NumArray) Update(i int, val int) {
	if len(this.Nums) > i {
		this.Nums[i] = val
	}
}

func (this *NumArray) SumRange(i int, j int) (res int) {
	for ; i < len(this.Nums) && i <= j; i++ {
		res += this.Nums[i]
	}
	return
}

//每日一题：659. 分割数组为连续子序列
func isPossible(nums []int) bool {
	//长度小于6肯定false
	//把重复的数字单独拿出来，尽量组成短的，最后剩下的拼接过去即可。

	return false
}
