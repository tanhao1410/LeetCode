package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(increasingTriplet([]int{1, 5, 3, 4}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//337. 打家劫舍 III
func rob(root *TreeNode) int {
	//思路：递归，头节点很重要，关系到
	if root == nil {
		return 0
	}
	sum1 := root.Val + rob2(root.Left) + rob2(root.Right)
	sum2 := rob(root.Left) + rob(root.Right)
	//问题：sum2虽然没取根节点，但是如果左右子树也没有取根节点的话，答案就不对了。但是这样的情况下，sum1肯定要大于sum2，即不会取这种情况。
	if sum1 > sum2 {
		return sum1
	}
	return sum2
}

//不能取根节点
func rob2(root *TreeNode) int {
	if root == nil || (root.Right == nil && root.Left == nil) {
		return 0
	}
	return rob(root.Left) + rob(root.Right)
}

//334. 递增的三元子序列
func increasingTriplet(nums []int) bool {
	//思路：记录最小值，次小值
	//用一个变量记录需要的数值，只要碰到递增的，就更新该数，遇到比该数大的，说明找到了，返回true
	needMore := math.MaxInt32
	min := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] > needMore {
			//说明找到了第三个数
			return true
		}
		//记录数组前面的最小值
		if nums[i] < min {
			min = nums[i]
		}
		//遇到递增的就可以判断了，要么前面有比小的还小的，要么后面有比大的还大的。
		if i+1 < len(nums) && nums[i+1] > nums[i] {
			if needMore > nums[i+1] {
				//需要更小的即可
				needMore = nums[i+1]
			}
			//如果前面最小的数比前一个数要小，说明找到了
			if nums[i] > min {
				return true
			}
		}
	}
	return false
}

//每日一题：387. 字符串中的第一个唯一字符
func firstUniqChar(s string) int {
	m := make([]int, 26)
	for i := len(s) - 1; i >= 0; i-- {
		m[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if m[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
