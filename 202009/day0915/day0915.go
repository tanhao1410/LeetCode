package main

import "fmt"

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	solveSudoku(board)

	for i := 0; i < 9; i++ {
		for k := 0; k < 9; k++ {
			fmt.Print(int(board[i][k])-'0', " ")
		}
		fmt.Println()
	}

	fmt.Println("....")
}

func solveSudoku(board [][]byte) {
	solveSudoku2(board)
}

//解数独1-9和. 里面存的是字符1-9
func solveSudoku2(board [][]byte) bool {
	//思路：先从第一行开始，
	for i := 0; i < 9; i++ {

		rowNums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
		//board[i]代表第一行，先填第一行可以填的
		rowFlag := false
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' { //说明是空位
				rowNums[board[i][j]-'0'] = 0 //代表这个数出现过
			} else {
				rowFlag = true
			}
		}

		//这一行都填满了，直接看下一行了
		if !rowFlag {
			continue
		}

		//往空位上填数字。只能填rowNums和smallSquare中没出现过的数
		//要填的位置在哪呢？
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				//要填的位置所在的小矩形中
				smallSquare := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
				for k := (i / 3) * 3; k < (i/3)*3+3; k++ {
					for l := (j / 3) * 3; l < (j/3)*3+3; l++ {
						if board[k][l] != '.' {
							smallSquare[board[k][l]-'0'] = 0 //说明在校小矩形中数字出现了
						}
					}
				}

				//要填的位置所在的列中
				colNums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
				for k := 0; k < 9; k++ {
					if board[k][j] != '.' {
						colNums[board[k][j]-'0'] = 0 //在列中出现了
					}
				}

				//即该位置只能填三个数组中都为出现过的，即值为1的下标值
				canSelected := []int{}
				for i := 1; i < 10; i++ {
					if rowNums[i] == 1 && colNums[i] == 1 && smallSquare[i] == 1 {
						canSelected = append(canSelected, i)
					}
				}

				//难点：回退时，有些数已经被更改了？？，如果要填的数没了，但还有空，说明，前面的肯定填错了。开始回退
				//if len(canSelected)== 0{
				//	return false
				//}

				flag := false
				//否则开始填下一个数字
				for _, v := range canSelected {
					board[i][j] = byte(v + '0')
					//开始填下一个
					flag = solveSudoku2(board)
					if flag {
						return flag
					}
				}

				//判断条件是什么呢？
				if !flag {
					board[i][j] = '.' // 都循环结束了，还是不行的话，需要将这个位置重新改为“."
				}
				return false

			}
		}

	}
	return true
}
