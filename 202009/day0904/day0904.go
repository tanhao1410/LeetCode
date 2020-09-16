package main

import (
	"fmt"
	"strconv"
)

func main() {
	//nums:= []int{3,2,1,2,1,7}

	//fmt.Println(minIncrementForUnique(nums))
	grid := [][]int{{2, 7, 6}, {1, 5, 9}, {4, 3, 8}}
	//fmt.Println(grid[0:2][1:4])
	fmt.Println(numMagicSquaresInside(grid))
}

//检测幻方
func numMagicSquaresInside(grid [][]int) int {
	//检测所有的子矩阵
	res := 0
	for i := 0; i < len(grid)-2; i++ {
		for j := 0; j < len(grid[0])-2; j++ {
			//子矩阵为
			//行 grid[i-2:i+1]
			if isMagicSquares(grid, i, j) {
				res++
			}
		}
	}

	return res
}

func isMagicSquares(grid [][]int, row, col int) bool {
	//必须要包含1-9
	nums := make([]int, 16)

	//检测对角线
	c1 := grid[row][col] + grid[row+1][col+1] + grid[row+2][col+2]
	c2 := grid[row][col+2] + grid[row+1][col+1] + grid[row+2][col]
	if c1 != c2 {
		return false
	}
	//每一行之和都相等
	rowNum := c1
	for i := 0; i < 3; i++ {
		rowNum2 := 0
		colNum := 0
		for j := 0; j < 3; j++ {
			nums[grid[row+i][col+j]] = 1
			rowNum2 += grid[row+i][col+j]
			colNum += grid[row+j][col+i]
		}
		if rowNum2 != rowNum || colNum != rowNum {
			return false
		}
	}

	for i := 1; i < 10; i++ {
		if nums[i] != 1 {
			return false
		}
	}

	return true
}

//单词搜索
func exist(board [][]byte, word string) bool {
	//

	return false
}

//二倍数对数组
func canReorderDoubled(A []int) bool {
	//一个数要找到它对应的二倍数或者1/2，当一个数既有二倍又有1/2时，不用管

	return false
}

//每次 move 操作将会选择任意 A[i]，并将其递增 1，使数组中所有的数唯一的最小增量
func minIncrementForUnique(A []int) int {
	//思路：需要将有重复的数进行增长到 数组没有的最小的比自己大的数
	//问题1：需要知道哪些数重复了，需要知道哪些数没有，因为数组中的数字的大小是确定的，0-40000,可以用个数组来表示
	n := make([]int, 40001)
	for _, v := range A {
		n[v]++ //说明这个数出现过，值为出现的次数
	}

	//因为是从小到大遍历的，因此不会出现多增长的情况。
	res := 0
	isPreNum := false //
	for i, j := 0, 0; i < 40001; {
		if n[i] > 1 {
			//说明出现了超过一次
			//找到后面是0的位置处放置
			if isPreNum {
				j = j + 1
			} else {
				j = i + 1
			}
			for ; j < 40001 && n[j] != 0; j++ {
			}
			if j < 40001 {
				res += (j - i)
				n[j] = 1
				n[i]--
				isPreNum = true //下次j不用再一点点跳了，直接跳到原来的j后面即可
			}
		} else {
			i++
		}
	}

	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//返回所有从根节点到叶子节点的路径
func binaryTreePaths(root *TreeNode) []string {
	res := &[]string{}
	MiddleReadTree(root, "", res)
	return *res
}

//中序遍历 1->2
func MiddleReadTree(root *TreeNode, pre string, res *[]string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		//说明是叶子节点加入返回列表中
		if pre == "" {
			*res = append(*res, strconv.Itoa(root.Val))
		} else {
			*res = append(*res, pre+"->"+strconv.Itoa(root.Val))
		}
		return
	}

	if root.Left != nil {
		if pre == "" {
			MiddleReadTree(root.Left, strconv.Itoa(root.Val), res)
		} else {
			MiddleReadTree(root.Left, pre+"->"+strconv.Itoa(root.Val), res)
		}

	}
	if root.Right != nil {
		if pre == "" {
			MiddleReadTree(root.Right, strconv.Itoa(root.Val), res)
		} else {
			MiddleReadTree(root.Right, pre+"->"+strconv.Itoa(root.Val), res)
		}

	}
}
