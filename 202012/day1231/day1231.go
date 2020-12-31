package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(eraseOverlapIntervals([][]int{{0, 2}, {1, 3}, {1, 3}, {2, 4}, {3, 5}, {3, 5}, {4, 6}}))

	arr := []int{1, 4, 7, 11, 15, 2, 5, 8, 12, 19, 3, 6, 9, 16, 22, 10, 13, 14, 17, 24, 18, 21, 23, 26, 30}
	for i := 0; i < 25; i++ {
		fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26,
			30}}, arr[i]))
	}
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
