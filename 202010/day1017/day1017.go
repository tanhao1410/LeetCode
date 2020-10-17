package main

import "fmt"

func main() {
	fmt.Println(totalNQueens(10))
}

//52.N皇后 II
func totalNQueens(n int) int {
	//返回总数即可，用一个数组记录填充的皇后，这是关键，回溯法，填充下一个
	chees := make([]int, n)
	res := 0
	next(chees, 0, &res)
	return res
}

//chees 即棋盘，index为下一个要填入的列号，res 记录结果值
func next(chees []int, index int, res *int) {
	if index >= len(chees) {
		*res += 1
	}
	n := len(chees)
	//先确定该位置可以填几
	for i := 0; i < n; i++ {
		can := true
		//行不能重复
		for j := 0; j < index; j++ {
			if chees[j] == i {
				can = false
			}
		}
		//对角线不能重复
		for j := 0; j < index; j++ {
			if (i < chees[j] && chees[j]-i == index-j) ||
				(i > chees[j] && i-chees[j] == index-j) {
				can = false
			}
		}
		if can {
			//都满足了，可以填充了
			chees[index] = i
			//填下一个
			next(chees, index+1, res)
		}
	}
}
