package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//nums:=[]int{2,5,2,1,2}
	//fmt.Println(combinationSum2(nums,5))

	//nums2:= []int{1,2,3}
	//fmt.Println(minMoves(nums2))
	fmt.Println(numSquares(53))
}

//完全平方数
func numSquares(n int) int {
	//第一，肯定要先找离n最近的那个平方数，也不一定，12 = 4 + 4+ 4找不到才会这样的 18 = 9+9 16 +1 +1
	//先找所有可能的平方数吧
	squares := []int{}
	for i := 1; i*i <= n; i++ {
		squares = append(squares, i*i)
	}
	res := math.MaxInt32
	selectedNum(squares, n, 0, &res)
	return res
}

func selectedNum(candidates []int, target int, count int, res *int) {
	if target == 0 {
		//符合条件了
		if count < *res {
			*res = count
		}
		return
	}

	if target < 0 {
		return
	}

	count++
	for i := 0; i < len(candidates) && target-candidates[i] >= 0; i++ {
		selectedNum(candidates, target-candidates[i], count, res)
	}

}

func minMoves(nums []int) int {
	res := 0
	for !addMore2(nums, &res) {
		//除了最大的数，其它全都加1
		//优化：如果说最大的比第二大的数都大了很多，那么res可以直接加上这些数
		//addOne(nums)
		//res+=addMore(nums)
		//继续优化，在相加的过程中判断是否已经全都相同了
	}
	return res
}

func addMore2(nums []int, res *int) bool {
	max, second, min := 0, 0, 0
	for k, v := range nums {
		if v > nums[max] {
			max, second = k, max
		}
		if v < nums[min] {
			min = k
		}
	}

	if min == max {
		return true
	}

	count := nums[max] - nums[second]
	if count == 0 {
		count = 1 //最少增加1
	}

	for i := 0; i < len(nums); i++ {
		if i != max {
			nums[i] += count
		}
	}
	*res += count

	return false
}

func addMore(nums []int) int {
	max, second := 0, 0
	for k, v := range nums {
		if v > nums[max] {
			max, second = k, max
		}
	}
	count := nums[max] - nums[second]
	if count == 0 {
		count = 1 //最少增加1
	}

	for i := 0; i < len(nums); i++ {
		if i != max {
			nums[i] += count
		}
	}
	return count
}

func isOk(nums []int) bool {
	temp := nums[0]
	for _, v := range nums {
		if v != temp {
			return false
		}
	}
	return true
}

func combinationSum2(candidates []int, target int) [][]int {
	//难点在于如何避免重复的出现
	//1.找到所有的后，从结果中去重
	//2.在找的过程中去重。先排序，就不会出现重复情况了，但会出现。1，1 1 2
	res := &[][]int{}

	//先对原数组进行排序
	slice := sort.IntSlice(candidates)
	slice.Sort()

	selected := []int{}
	combination(selected, 0, candidates, target, res)

	return *res
}

//index代表只能从哪一位开始找
func combination(selected []int, index int, candidates []int, target int, res *[][]int) {

	if target == 0 {
		*res = append(*res, selected)
		return
	}

	//如果没有候选数了，target不为0,直接退出
	if len(candidates) == index {
		return
	}
	//如果候选的中最小的都超过了想要的数了，直接返回
	if candidates[index] > target {
		return
	}
	//如果能从中找到,就返回了
	preNum := -1
	for i := index; i < len(candidates) && candidates[i] <= target; i++ {
		if candidates[i] == preNum {
			continue
		}
		preNum = candidates[i]
		selected2 := []int{}
		for _, v := range selected {
			selected2 = append(selected2, v)
		}
		selected2 = append(selected2, candidates[i])
		combination(selected2, i+1, candidates, target-candidates[i], res)
	}
}
