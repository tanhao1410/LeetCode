package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := [][]int{{2, 3, 8, 11, 15, 19, 20, 20}, {4, 8, 12, 15, 18, 21, 25, 28}, {5, 8, 17, 20, 22, 23, 30, 34}, {6, 12, 18, 20, 25, 25, 34, 34}, {9, 14, 21, 24, 25, 29, 39, 40}}
	fmt.Println(findNumberIn2DArray(matrix, 12))
}

//二维数组中的查找
//思路：先找对应的行，再查对应的列
func findNumberIn2DArray(matrix [][]int, target int) bool {
	//1.确定该数可能在哪一行？数比头一个大，比最后一个小，即是有可能的。
	m := len(matrix) // 行数
	if m == 0 {
		return false
	}
	n := len(matrix[0]) // 列数
	if n == 0 {
		return false
	}

	//找可能的行 ，行首小于 等于target,行尾大于target
	rowStart, rowEnd := math.MaxInt32, 0
	for i := 0; i < m; i++ {

		if matrix[i][0] == target || matrix[i][n-1] == target {
			return true
		}

		if matrix[i][0] < target && matrix[i][n-1] > target && rowStart > i {
			rowStart, rowEnd = i, i
		}

		if matrix[i][0] > target {
			rowEnd = i - 1
			break
		} else {
			rowEnd = i
		}

	}

	//开始二分查找
	for ; rowStart <= rowEnd; rowStart++ {

		start, end := 0, n-1
		for middle := (start + end) / 2; start <= end; middle = (start + end) / 2 {
			if matrix[rowStart][middle] == target {
				return true
			}
			if matrix[rowStart][middle] < target {
				start = middle + 1
			}
			if matrix[rowStart][middle] > target {
				end = middle - 1
			}
		}

	}

	return false
}
