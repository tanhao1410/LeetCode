package main

func main() {

}

//剑指 Offer 39. 数组中出现次数超过一半的数字
func majorityElement(nums []int) int {
	//最简单的思路是记录每个数出现的次数
	res := nums[0]
	m := make(map[int]int)
	for _, v := range nums {
		if count, ok := m[v]; ok {
			m[v]++
			if count+1 > len(nums)/2 {
				return v
			}
		} else {
			m[v] = 1
		}
	}

	return res
}

//每日一题：164. 最大间距
func maximumGap(nums []int) int {

	return 0
}
