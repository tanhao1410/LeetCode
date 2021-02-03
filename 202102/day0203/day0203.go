package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{1, 0, 3}))
}

//268. 丢失的数字
func missingNumber(nums []int) int {
	//最简单思路：空间复杂度o(n) ，或者排序完成。
	temp := make([]bool, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		temp[nums[i]] = true
	}

	for i := 0; ; i++ {
		if !temp[i] {
			return i
		}
	}

}

//292. Nim 游戏
func canWinNim(n int) bool {
	return n%4 == 0
}
