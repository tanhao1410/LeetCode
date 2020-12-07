package main

import "fmt"

func main() {
	fmt.Println(matrixScore([][]int{{0, 0, 0, 1, 1, 0}, {1, 1, 0, 1, 0, 1}, {1, 0, 1, 0, 0, 1}}))
}

//每日一题：861. 翻转矩阵后的得分
func matrixScore(A [][]int) int {
	//思路：优先第一列变为1，通过行移动来完成
	for i := 0; i < len(A); i++ {
		if A[i][0] == 0 {
			//移动该行
			for j := 0; j < len(A[0]); j++ {
				if A[i][j] == 1 {
					A[i][j] = 0
				} else {
					A[i][j] = 1
				}
			}
		}
	}
	res := len(A) * 1 << (len(A[0]) - 1)
	//剩下的开始，优先通过列移动，每一列尽量多1，直到完成目标
	for i := 1; i < len(A[0]); i++ {
		oneCount := 0
		for row := 0; row < len(A); row++ {
			if A[row][i] == 1 {
				oneCount++
			}
		}
		if oneCount < (len(A)+1)/2 {
			oneCount = len(A) - oneCount
		}
		num := 1 << (len(A[0]) - i - 1) * oneCount
		res += num
	}
	return res
}
