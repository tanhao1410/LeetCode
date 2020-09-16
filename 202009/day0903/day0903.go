package main

import "fmt"

func main() {
	arr := []int{2}
	fmt.Println(countTriplets(arr))
}

//n皇后问题暴力解法
func solveNQueens(n int) [][]string {
	res := [][]string{}

	//初始化棋盘
	nn := make([][]int, n)
	for i := 0; i < n; i++ {
		nn[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		nn[0][i] = 1 //第一行的第i个位置放上

	}

	return res
}

func SetQueen(nn [][]int, row int) {
	if row == len(nn)-1 { //说明是最后一个字了
		//按位置放好，看OK不OK，如果OK，打印结果，并清空棋盘，开始下一轮循环
		for i := 0; i < len(nn); i++ {

		}
	}

	for i := 0; i < row; {
		//按可能的地方摆放
		SetQueen(nn, row+1)
	}
}

func IsQueenOk(nn [][]int) {

}

//种花问题
func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		//可不可以种花需要看它的前面和后面
		if flowerbed[i] == 0 &&
			(i-1 < 0 || flowerbed[i-1] == 0) &&
			(i+1 >= len(flowerbed) || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
		}
	}
	return n <= 0
}

//先水平反转元素，再反转元素1<-->0
func flipAndInvertImage(A [][]int) [][]int {
	//水平反转并改变数字
	for i := 0; i < len(A); i++ {
		for j, k := 0, len(A[0])-1; j <= k; {
			A[i][j], A[i][k] = 1&^A[i][k], 1&^A[i][j] //交换值，并且1<-->0
			j, k = j+1, k-1
		}
	}
	return A
}

//两球之间的磁力，先把两个放在最大处和最小处，第三个，放置在中间（最大+最小/2）离的最近的
//如果有四个呢？
func maxDistance(position []int, m int) int {
	return 0
}

//i <j <=k
//a = arr[i] ^ arr[i + 1] ^ ... ^ arr[j - 1]
//b = arr[j] ^ arr[j + 1] ^ ... ^ arr[k]
//a==b
func countTriplets(arr []int) int {

	//思路：dp[i][j] = arr[i] ^ arr[i + 1] ^ ... ^ arr[j]
	dp := make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		row := make([]int, len(arr))
		dp[i] = row
	}

	for i := 0; i < len(arr); i++ {
		dp[i][i] = arr[i]
		for j := i + 1; j < len(arr); j++ {
			dp[i][j] = dp[i][j-1] ^ arr[j]
		}
	}

	res := 0
	//找符合条件的，遍历dp数组，即找dp[i][j-1] == dp[j][k]
	for i := 0; i < len(dp); i++ {
		for j := i + 1; j < len(dp); j++ {
			for k := j; k < len(dp); k++ {
				if dp[j][k] == dp[i][j-1] {
					res++
				}
			}
		}
	}

	return res

}

//高效法2dp[i][k] =0;res += k-j
