package main

import "fmt"

func main() {
	//nums := []int{1, 0, 1, 2, 0, 0}
	//sortColors(nums)
	//fmt.Println(nums)
	fmt.Println(maxPoints([][]int{{-6, -65}, {-18, 26}, {-6, -65}, {10, -2}, {10, -2}}))
}

//剑指 Offer 29. 顺时针打印矩阵
func spiralOrder(matrix [][]int) []int {
	res := []int{}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}
	r, c := len(matrix), len(matrix[0])
	for x, y := 0, -1; len(res) < len(matrix)*len(matrix[0]); {

		//向右走 c 步
		for i := 0; i < c; i++ {
			y++
			res = append(res, matrix[x][y])
		}

		//向下走 r - 1步
		for i := 0; i < r-1; i++ {
			x++
			res = append(res, matrix[x][y])
		}

		//向左走 c - 1步
		for i := 0; i < c-1; i++ {
			y--
			res = append(res, matrix[x][y])
		}

		//向上走r - 2步
		for i := 0; i < r-2; i++ {
			x--
			res = append(res, matrix[x][y])
		}

		r -= 2
		c -= 2
	}
	return res[:len(matrix)*len(matrix[0])]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//每日一题：222. 完全二叉树的节点个数
func countNodes(root *TreeNode) int {
	//最简单的思路就是遍历即可。但效率？
	//直接求第n层个数即可。
	res := 0
	if root != nil {
		res = countNodes(root.Left) + countNodes(root.Right) + 1
	}
	return res
}

//75. 颜色分类
func sortColors(nums []int) {
	//思路：计数法
	zero, one := 0, 0
	for _, v := range nums {
		if v == 0 {
			zero++
		} else if v == 1 {
			one++
		}
	}

	for i := 0; i < len(nums); i++ {
		if i < zero {
			nums[i] = 0
		} else if i < zero+one {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}
}

//149. 直线上最多的点数
func maxPoints(points [][]int) int {
	//思路：
	//两个点相同的时候，会出现，第三个点一定在同一直线上的问题。
	isOneLine := func(p1, p2, p3 []int) bool {
		return (p3[1]-p1[1])*(p2[0]-p1[0]) ==
			(p2[1]-p1[1])*(p3[0]-p1[0])
	}
	if len(points) < 3 {
		return len(points)
	}
	res := 0
	//用一个数组记录哪个线段被用了
	//先从第一个点开始
	for k, first := range points {

		//找下一个点，以凑成线
		samePointNum := 0
		for i := k + 1; i < len(points); i++ {
			//即这个点已经和刚才形成了线了，不用重复计算了
			second := points[i]
			//如果第一条线和第二条线是重合的话
			if second[0] == first[0] && second[1] == first[1] {
				samePointNum += 1
				continue
			}
			//所有的点都是同一个点的话，直接就结束了
			count := 2 + samePointNum
			//samePointNum = 0
			for third := i + 1; third < len(points); third++ {
				if isOneLine(first, second, points[third]) {
					count++
				}
			}
			if count > res {
				res = count
			}
		}
		if samePointNum+1 > res {
			res = samePointNum + 1
		}
	}
	return res
}
