package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(eraseOverlapIntervals([][]int{{0, 2}, {1, 3}, {1, 3}, {2, 4}, {3, 5}, {3, 5}, {4, 6}}))

	arr := []int{1, 4, 7, 11, 15, 2, 5, 8, 12, 19, 3, 6, 9, 16, 22, 10, 13, 14, 17, 24, 18, 21, 23, 26, 30}
	for i := 0; i < 25; i++ {
		fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26,
			30}}, arr[i]))
	}
}

type NestedInteger struct {
}

func (n NestedInteger) IsInteger() bool { return true }

func (n NestedInteger) GetInteger() int { return 0 }

func (n *NestedInteger) SetInteger(value int) {}

func (n *NestedInteger) Add(elem NestedInteger) {}

func (n NestedInteger) GetList() []*NestedInteger { return nil }

//385. 迷你语法分析器
func deserialize(s string) *NestedInteger {
	if len(s) == 0 {
		return new(NestedInteger)
	}
	res := new(NestedInteger)
	//如果s是一个纯数字
	if (s[0] >= '0' && s[0] <= '9') || s[0] == '-' {
		num, _ := strconv.Atoi(s)
		res.SetInteger(num)
	} else {
		//说明是以[开头的数组类
		for _, v := range processS(s) {
			res.Add(*deserialize(v))
		}
	}
	return res
}

//把一个[xx,xxx,[xx],xx]分成数组
func processS(s string) []string {
	res := []string{}
	for i := 1; i < len(s)-1; {
		if s[i] == ',' {
			i++
			continue
		}
		if s[i] == '[' {
			stack := 1
			j := i + 1
			for ; stack > 0; j++ {
				if s[j] == '[' {
					stack++
				} else if s[j] == ']' {
					stack--
				}
			}
			res = append(res, s[i:j])
			i = j
		} else {
			//碰到数字了
			j := i + 1
			for ; j < len(s)-1; j++ {
				if (s[j] > '9' || s[j] < '0') && s[j] != '-' {
					break
				}
			}
			res = append(res, s[i:j])
			i = j
		}
	}
	return res
}

//面试题 17.08. 马戏团人塔
func bestSeqAtIndex(height []int, weight []int) int {
	if len(height) == 0 {
		return 0
	}
	//思路：先排序，然后类似于俄罗斯套娃问题，动态规划的算法，dp[i]自己在顶部的最大人数。dp[i+1]等于前面的可以放的+1，前面 到某几位呢？
	//不能简单的排序，因为身高和体重是一一对应的。
	hw := make([][]int, len(height))
	for i := 0; i < len(height); i++ {
		hw[i] = []int{height[i], weight[i]}
	}
	sort.Slice(hw, func(i, j int) bool {
		//首先按身高排序，应该矮的在后
		return hw[i][0] > hw[j][0]

	})

	//它的前面可以叠多少人
	dp := make([]int, len(hw))
	//记录能形成的最长人数
	dpMax := make([]int, len(hw))
	dp[0] = 1
	dpMax[0] = 1

	for i := 1; i < len(dp); i++ {
		//从后往前找，找到第一个它可以放上去的。同身高下，轻的放在了前面，轻的dp肯定要比重的大或相等。前面的身高肯定要比后面的要矮或相等。
		max := 0
		for j := i - 1; j >= 0; j-- {
			if hw[j][0] > hw[i][0] && hw[j][1] > hw[i][1] && dp[j] > max {
				max = dp[j]
			}
			if max >= dpMax[j] {
				break
			}
		}
		dp[i] = max + 1
		if dpMax[i-1] < dp[i] {
			dpMax[i] = dp[i]
		} else {
			dpMax[i] = dpMax[i-1]
		}
	}
	return dpMax[len(dpMax)-1]
}

//面试题 10.09. 排序矩阵查找
func searchMatrix(matrix [][]int, target int) bool {
	//矩阵为空的时候直接返回
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	//二分法查找数字
	binarySearch := func(nums []int, target int) bool {
		s, e := 0, len(nums)-1
		for m := (s + e) / 2; s <= e; m = (s + e) / 2 {
			if nums[m] == target {
				return true
			} else if nums[m] > target {
				e = m - 1
			} else {
				s = m + 1
			}
		}
		return false
	}

	var search func(rowStart, rowEnd, colStart, colEnd int) bool
	search = func(rowStart, rowEnd, colStart, colEnd int) bool {
		if rowEnd < rowStart || colEnd < colStart {
			return false
		}
		if rowStart == rowEnd {
			//就剩下一行了。二分查找即可。
			return binarySearch(matrix[rowEnd][colStart:colEnd+1], target)
		}
		//查看中间的是否相等
		x, y := (rowStart+rowEnd)/2, (colStart+colEnd)/2
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] > target {
			//中间的数比目标值还大，答案可能在上面，本行，左下
			return search(rowStart, x-1, colStart, colEnd) || binarySearch(matrix[x], target) || search(x+1,
				rowEnd, colStart, colEnd-1)
		} else {
			//中间的数比目标值小，答案可能在本行，下面，或者右上方
			return search(x+1, rowEnd, colStart, colEnd) || binarySearch(matrix[x], target) || search(rowStart,
				x-1, y+1, colEnd)
		}
		return false
	}

	return search(0, len(matrix)-1, 0, len(matrix[0])-1)
}

//每日一题：435. 无重叠区间
func eraseOverlapIntervals(intervals [][]int) int {
	res := 0
	//思路：先找到重叠的区间，根据区间的开始来排序，如果区间的结束大于后面的开始，则说明，有重叠
	//移除最少的区间，使之无重叠？先确定谁和谁之间有重叠了。谁重叠的多，去掉谁？怎么确定多？根据与之重叠的数量排序。去掉一个区间后，与之重叠的数也会减少一个，需要注意，直到无重叠区间。
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	m := make(map[int]map[int]bool)
	for i := 0; i < len(intervals); i++ {
		//将与它对应的重叠区间index加入进来，
		for j := i + 1; j < len(intervals) && intervals[j][0] < intervals[i][1]; j++ {
			if m[i] == nil {
				item := make(map[int]bool)
				item[j] = true
				m[i] = item
			} else {
				m[i][j] = true
			}

			//重叠区间应该是相互的。
			if m[j] == nil {
				item := make(map[int]bool)
				item[i] = true
				m[j] = item
			} else {
				m[j][i] = true
			}
		}
	}

	//删除最大的哪个，直接找最大的哪一个就好了
	for {

		//找重叠数量最多的一个？重叠数最多，不一定就要删它。
		delIndex, delLen := 0, 0
		for k, v := range m {
			if len(v) > delLen {
				delIndex = k
				delLen = len(v)
			}
		}
		if delLen == 0 {
			return res
		}
		res++

		for k, _ := range m[delIndex] {
			delete(m[delIndex], k)
			//删除kk中对应的k
			delete(m[k], delIndex)
		}
	}

	return res
}
