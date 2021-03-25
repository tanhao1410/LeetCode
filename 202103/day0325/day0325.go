package main

func main() {

}

//面试题 08.10. 颜色填充
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	//颜色相同的相连的才进行替换
	var alterColor func(x, y int, old int)
	alterColor = func(x, y int, old int) {
		image[x][y] = 65536
		//改变上下左右
		if x-1 >= 0 && image[x-1][y] == old {
			alterColor(x-1, y, old)
		}
		if x+1 < len(image) && image[x+1][y] == old {
			alterColor(x+1, y, old)
		}
		if y-1 >= 0 && image[x][y-1] == old {
			alterColor(x, y-1, old)
		}
		if y+1 < len(image[0]) && image[x][y+1] == old {
			alterColor(x, y+1, old)
		}
	}

	alterColor(sr, sc, image[sr][sc])

	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[0]); j++ {
			if image[i][j] == 65536 {
				image[i][j] = newColor
			}
		}
	}

	return image
}

//面试题 08.06. 汉诺塔问题
func hanota(A []int, B []int, C []int) []int {
	//思路：当盘子中就一个时，直接放，如果是两个，先先把所有的，处最后一个外，放在中间的，剩下的这个放到C里
	if len(A) == 0 {
		return C
	}

	for len(A) > 1 {
		B = append(B, A[len(A)-1])
		A = A[:len(A)-1]
	}
	C = append(C, A[0])
	A = A[:0]
	for len(B) > 0 {
		A = append(A, B[len(B)-1])
		B = B[:len(B)-1]
	}
	return hanota(A, B, C)
}
