package main

func main() {

}

//529. 扫雷游戏
func updateBoard(board [][]byte, click []int) [][]byte {
	//如果一个地雷（'M'）被挖出，游戏就结束了- 把它改为 'X'。
	//如果一个没有相邻地雷的空方块（'E'）被挖出，修改它为（'B'），并且所有和其相邻的未挖出方块都应该被递归地揭露。
	//如果一个至少与一个地雷相邻的空方块（'E'）被挖出，修改它为数字（'1'到'8'），表示相邻地雷的数量。
	//如果在此次点击中，若无更多方块可被揭露，则返回面板。

	x, y := click[0], click[1]

	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
		return nil
	}

	haveBomb := func(x, y int) int {
		res := 0
		if x+1 < len(board) && board[x+1][y] == 'M' {
			res++
		}
		if x-1 > -1 && board[x-1][y] == 'M' {
			res++
		}
		if y-1 > -1 && board[x][y-1] == 'M' {
			res++
		}
		if y+1 < len(board[0]) && board[x][y+1] == 'M' {
			res++
		}
		if x+1 < len(board) && y+1 < len(board[0]) && board[x+1][y+1] == 'M' {
			res++
		}
		if x-1 > -1 && y-1 > -1 && board[x-1][y-1] == 'M' {
			res++
		}
		if x+1 < len(board) && y-1 > -1 && board[x+1][y-1] == 'M' {
			res++
		}
		if x-1 > -1 && y+1 < len(board[0]) && board[x-1][y+1] == 'M' {
			res++
		}
		return res
	}

	//看点击中的是什么
	switch board[x][y] {
	case 'E':
		//点击了空方块，如果空方块不和雷相邻，与它相邻的所有E都要改成B，与地雷相邻的要改成带数字的
		if haveBomb(x, y) == 0 {
			board[x][y] = 'B'
			//递归调用周围的格
			updateBoard(board, []int{x + 1, y})
			updateBoard(board, []int{x + 1, y + 1})
			updateBoard(board, []int{x + 1, y - 1})
			updateBoard(board, []int{x - 1, y})
			updateBoard(board, []int{x - 1, y + 1})
			updateBoard(board, []int{x - 1, y - 1})
			updateBoard(board, []int{x, y + 1})
			updateBoard(board, []int{x, y - 1})
		} else {
			//相邻的话，要改成 数字
			board[x][y] = byte('0' + haveBomb(x, y))
		}
	case 'M':
		//点击了地雷，雷位置要改成X,结束游戏
		board[x][y] = 'X'
	}
	return board
}

//每日一题：547. 省份数量
func findCircleNum(isConnected [][]int) int {
	res := 0
	m := make(map[int]map[int]bool)
	//第一次遍历形成直接相连的两城市
	for i := 0; i < len(isConnected); i++ {
		m[i] = make(map[int]bool)
		for j := 0; j < len(isConnected); j++ {
			if i != j && isConnected[i][j] == 1 {
				m[i][j] = true
			}
		}
	}
	//已经加入到某个省份的城市集合
	alreadyCitys := make(map[int]bool)

	//从第一个城市开始算起
	for i := 0; i < len(isConnected); i++ {
		//没有加入集合的才会计算
		if !alreadyCitys[i] {
			res++
			for noAdd := false; !noAdd; {
				//遍历它所有的能连接的城市
				noAdd = true
				for city, _ := range m[i] {
					//没有遍历过的城市
					if !alreadyCitys[city] {
						alreadyCitys[city] = true
						for city2, _ := range m[city] {
							if !m[i][city2] {
								//非直接相连的城市，将其加入进来，下次遍历时就会遍历到它
								m[i][city2] = true
								noAdd = false
							}
						}
					}
				}
			}
		}
	}

	return res
}
