package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumEffortPath([][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}}))
}

//705. 设计哈希集合
type MyHashSet struct {
	//数据集合
	nums [][]int
}

func Constructor() MyHashSet {
	nums := make([][]int, 100)
	return MyHashSet{
		nums: nums,
	}
}

//默认分成100份
func (this *MyHashSet) hash(key int) int {
	return key % 100
}

func (this *MyHashSet) Add(key int) {
	nums := this.nums[this.hash(key)]
	for i := 0; i < len(nums); i++ {
		if nums[i] == key {
			return
		}
	}
	this.nums[this.hash(key)] = append(this.nums[this.hash(key)], key)
}

func (this *MyHashSet) Remove(key int) {
	if this.Contains(key) {
		nums := this.nums[this.hash(key)]
		for i := 0; i < len(nums); i++ {
			if nums[i] == key {
				nums[i] = nums[len(nums)-1]
				break
			}
		}
		this.nums[this.hash(key)] = this.nums[this.hash(key)][:len(nums)-1]
	}
}

func (this *MyHashSet) Contains(key int) bool {
	nums := this.nums[this.hash(key)]
	for i := 0; i < len(nums); i++ {
		if nums[i] == key {
			return true
		}
	}
	return false
}

//每日一题：1631. 最小体力消耗路径
func minimumEffortPath(heights [][]int) int {

	//采用动态规划算法，每更新一个位置的体力消耗后，都要看能否更新它周围的，直到它周围的不能改变了为止

	//初始化动态规划表
	dp := make([][]int, len(heights))
	for i := 0; i < len(heights); i++ {
		dp[i] = make([]int, len(heights[0]))
		for j := 0; j < len(heights[0]); j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 0

	getAbstractNum := func(n1, n2 int) int {
		if n1 > n2 {
			return n1 - n2
		}
		return n2 - n1
	}

	max := func(n1, n2 int) int {
		if n1 > n2 {
			return n1
		}
		return n2
	}

	var updateDp func(x, y int)
	updateDp = func(i, j int) {

		//它的上方是已经可以到达的
		curNum := heights[i][j]
		if i > 0 && dp[i-1][j] != math.MaxInt32 {
			if max(getAbstractNum(curNum, heights[i-1][j]), dp[i][j]) < dp[i-1][j] {
				dp[i-1][j] = max(getAbstractNum(curNum, heights[i-1][j]), dp[i][j])
				//递归更新它周围的
				updateDp(i-1, j)
			}
		}
		//从下方
		if i < len(dp)-1 && dp[i+1][j] != math.MaxInt32 {
			if max(getAbstractNum(curNum, heights[i+1][j]), dp[i][j]) < dp[i+1][j] {
				dp[i+1][j] = max(getAbstractNum(curNum, heights[i+1][j]), dp[i][j])
				//递归更新它周围的
				updateDp(i+1, j)
			}
		}

		//从左方
		if j > 0 && dp[i][j-1] != math.MaxInt32 {
			if max(getAbstractNum(curNum, heights[i][j-1]), dp[i][j]) < dp[i][j-1] {
				dp[i][j-1] = max(getAbstractNum(curNum, heights[i][j-1]), dp[i][j])
				//递归更新它周围的
				updateDp(i, j-1)
			}
		}

		//从右方
		if j < len(dp[0])-1 && dp[i][j+1] != math.MaxInt32 {
			if max(getAbstractNum(curNum, heights[i][j+1]), dp[i][j]) < dp[i][j+1] {
				dp[i][j+1] = max(getAbstractNum(curNum, heights[i][j+1]), dp[i][j])
				//递归更新它周围的
				updateDp(i, j+1)
			}
		}
	}

	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[0]); j++ {
			//与矩阵前后左右比较，如果位置已经到达，

			curRes := dp[i][j]
			curNum := heights[i][j]

			//它的上方是已经可以到达的
			if i > 0 && dp[i-1][j] != math.MaxInt32 {
				if max(getAbstractNum(curNum, heights[i-1][j]), dp[i-1][j]) < curRes {
					curRes = max(getAbstractNum(curNum, heights[i-1][j]), dp[i-1][j])
				}
			}
			//从下方
			if i < len(dp)-1 && dp[i+1][j] != math.MaxInt32 {
				if max(getAbstractNum(curNum, heights[i+1][j]), dp[i+1][j]) < curRes {
					curRes = max(getAbstractNum(curNum, heights[i+1][j]), dp[i+1][j])
				}
			}

			//从左方
			if j > 0 && dp[i][j-1] != math.MaxInt32 {
				if max(getAbstractNum(curNum, heights[i][j-1]), dp[i][j-1]) < curRes {
					curRes = max(getAbstractNum(curNum, heights[i][j-1]), dp[i][j-1])
				}
			}

			//从右方
			if j < len(dp[0])-1 && dp[i][j+1] != math.MaxInt32 {
				if max(getAbstractNum(curNum, heights[i][j+1]), dp[i][j+1]) < curRes {
					curRes = max(getAbstractNum(curNum, heights[i][j+1]), dp[i][j+1])
				}
			}

			dp[i][j] = curRes

			//更新curRes后，还需要更新它周围的dp
			updateDp(i, j)
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]

}

//784. 字母大小写全排列
func letterCasePermutation(S string) []string {
	//总共有多少种是已知的，数量为2^字母数量，通过调整字母的变化返回即可。使用递归形式完成。
	res := []string{}

	var selectLetter func(index int, s string)
	selectLetter = func(index int, s string) {
		//如果选择的字母是最后一个了，递归结束
		if index == len(S) {
			res = append(res, s)
			return
		}

		//该位置是数字
		if S[index] <= '9' && S[index] >= '0' {
			selectLetter(index+1, s+string(S[index]))
		} else {
			//大写与小写
			if S[index] <= 'Z' && S[index] >= 'A' {
				selectLetter(index+1, s+string(S[index]))
				selectLetter(index+1, s+string(S[index]+'a'-'A'))
			} else {
				selectLetter(index+1, s+string(S[index]))
				selectLetter(index+1, s+string(S[index]+'A'-'a'))
			}
		}
	}
	selectLetter(0, "")
	return res
}
