package main

func main() {

}

//120. 三角形最小路径和
func minimumTotal(triangle [][]int) int {
	//如果长度为1,直接就返回了
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			//它可以从上面过来，要么两条路，要么一条路
			if j < 1 {
				triangle[i][j] += triangle[i-1][0]
			} else if j == len(triangle[i])-1 {
				triangle[i][j] += triangle[i-1][len(triangle[i])-2]
			} else {
				//从哪边小从哪边走
				if triangle[i-1][j-1] > triangle[i-1][j] {
					triangle[i][j] += triangle[i-1][j]
				} else {
					triangle[i][j] += triangle[i-1][j-1]
				}
			}
		}
	}

	//从最后一行中找最小的数
	lastRow := triangle[len(triangle)-1]
	res := lastRow[0]
	for i := 1; i < len(lastRow); i++ {
		if lastRow[i] < res {
			res = lastRow[i]
		}
	}

	return res
}
