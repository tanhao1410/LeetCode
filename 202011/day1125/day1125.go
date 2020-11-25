package main

import (
	"fmt"
	"sort"
)

func main() {
	board := [][]int{{1, 1, 0}, {1, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(board)
}

//274. H 指数
func hIndex(citations []int) int {
	res := 0
	sort.Ints(citations)
	for i := len(citations) - 1; i >= 0; i-- {
		if citations[i] <= len(citations)-i {
			if citations[i] > res {
				return citations[i]
			}
			return res
		}
		res = len(citations) - i
	}
	return res
}

//289. 生命游戏
func gameOfLife(board [][]int) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	m, n := len(board), len(board[0])
	//思路：可以用置2来解决本地的问题,2代表原来是死的现在变活了，3代表原来是活的，现在死了
	liveCount := func(x, y int) int {
		count := 0

		if x-1 >= 0 && (board[x-1][y] == 1 || board[x-1][y] == 3) {
			count += 1
		}
		if x-1 >= 0 && y-1 >= 0 && (board[x-1][y-1] == 1 || board[x-1][y-1] == 3) {
			count++
		}
		if x-1 >= 0 && y+1 < n && (board[x-1][y+1] == 1 || board[x-1][y+1] == 3) {
			count++
		}
		if y-1 >= 0 && (board[x][y-1] == 1 || board[x][y-1] == 3) {
			count++
		}
		if y+1 < n && (board[x][y+1] == 1 || board[x][y+1] == 3) {
			count++
		}
		if x+1 < m && (board[x+1][y] == 1 || board[x+1][y] == 3) {
			count++
		}
		if x+1 < m && y-1 >= 0 && (board[x+1][y-1] == 1 || board[x+1][y-1] == 3) {
			count++
		}
		if x+1 < m && y+1 < n && (board[x+1][y+1] == 1 || board[x+1][y+1] == 3) {
			count++
		}
		return count
	}

	for x, row := range board {
		for y, v := range row {
			count := liveCount(x, y)
			//如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
			//活细胞周围八个位置有超过三个活细胞
			if (count < 2 || count > 3) && v == 1 {
				board[x][y] = 3
			} else if v == 0 && count == 3 {
				board[x][y] = 2
			}
		}
	}
	for x, row := range board {
		for y, v := range row {
			if v == 2 {
				board[x][y] = 1
			} else if v == 3 {
				board[x][y] = 0
			}
		}
	}
}

//299. 猜数字游戏
func getHint(secret string, guess string) string {
	//思路：先记录全对的个数。然后统计其他数字的个数，找共同的个数即可
	a, b := 0, 0
	m1, m2 := make([]int, 10), make([]int, 10)
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			a++
		} else {
			m1[secret[i]-'0'] += 1
			m2[guess[i]-'0'] += 1
		}
	}
	for i := 0; i < 10; i++ {
		if m1[i] < m2[i] {
			b += m1[i]
		} else {
			b += m2[i]
		}
	}

	return fmt.Sprintf("%s%s%s%s", a, "A", b, "B")
}

//每日一题：1370. 上升下降字符串
func sortString(s string) string {
	res := ""
	//思路：用一个26个大小的数组，记录所有的字母的数量
	m := make([]int, 26)
	for _, v := range s {
		m[v-'a'] += 1
	}
	//开始拼接字符串
	//从中选择最小的
	selectMin := func() {
		for i := 0; i < 26; i++ {
			if m[i] > 0 {
				m[i]--
				res += string('a' + i)
			}
		}
	}
	selectMax := func() {
		for i := 25; i >= 0; i-- {
			if m[i] > 0 {
				m[i]--
				res += string('a' + i)
			}
		}
	}
	for len(res) < len(s) {
		//先选最小的
		selectMin()
		selectMax()
	}
	return res
}
