package main

import "fmt"

func main() {
	//nums := []int{2,3,5}
	//fmt.Println(combinationSum(nums,8))
	fmt.Println(modifyString(""))
}

//数组中重复的数据，不用额外空间，时间复杂度o(1)，找到所有出现两次的元素
func findDuplicates(nums []int) []int {

	return nums
}

func modifyString(s string) string {
	bytes := []byte(s)
	cs := []byte{'a', 'b', 'c'}
	for i := 0; i < len(s); i++ {
		if s[i] == '?' {
			//需要替换
			if i > 0 && i < len(s)-1 {
				//不和两边相同即可
				for j := 0; j < 3; j++ {
					if cs[j] != bytes[i-1] && cs[j] != bytes[i+1] {
						bytes[i] = cs[j]
						break
					}
				}
			} else if i > 0 {
				//i是最后一个字母了
				for j := 0; j < 3; j++ {
					if cs[j] != bytes[i-1] {
						bytes[i] = cs[j]
						break
					}
				}
			} else if i < len(s)-1 {
				//i第一个字母
				for j := 0; j < 3; j++ {
					if cs[j] != bytes[i+1] {
						bytes[i] = cs[j]
						break
					}
				}
			} else if i == 0 {
				bytes[0] = 'a'
			}
		}
	}
	return string(bytes)
}

//组合总数
func combinationSum(candidates []int, target int) [][]int {
	res := &[][]int{}
	//思路：不重复，可以先一个个来，先选择第一个数，后面的还是不好选啊
	for k, v := range candidates {
		//先选一个，再递归调用选择另一个
		item := []int{v}
		combinationSum2(candidates[k:], target, item, res)
	}
	return *res
}

func combinationSum2(candidates []int, target int, item []int, res *[][]int) {
	count := 0
	for _, v := range item {
		count += v
	}
	if count == target {
		*res = append(*res, item)
		return
	}

	if count > target {
		return
	}

	//继续选择下一个
	for k, v := range candidates {
		//先选一个，再递归调用选择另一个
		nitem := []int{v}
		for _, v2 := range item {
			nitem = append(nitem, v2)
		}
		combinationSum2(candidates[k:], target, nitem, res)
	}

}
