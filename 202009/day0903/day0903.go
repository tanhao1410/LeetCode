package main

import "fmt"

func main2() {
	//用一个数组来记录即可，代表每一列的皇后放在第几行。
	//var balance = [8]int{0, 0, 0, 0, 0, 0, 0, 0}
	//queen(balance, 0)

	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	//两个数的组合放置于map中，
	m := make(map[int]int, 1)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	res := [][]int{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			//两数之和为
			sum := nums[i] + nums[j]
			if mi, ok := m[-sum]; ok {
				//如果里面有，
				if mi != i && mi != j { // 只能说明mi和i,j不是同一个
					//不代表这三个数没出现过
					//判断这三个数没出现过，按大小顺序,问题 0,0,0,0
					//即使是都按顺序，也会出现重复的。
					if -sum >= nums[j] && nums[j] >= nums[i] {
						item := []int{nums[i], nums[j], -sum}
						res = append(res, item)
					}
				}
			}
		}
	}

	return res
}

//三数之和[[-1, 0, 1],[-1, -1, 2]]
func threeSum2(nums []int) [][]int {
	//两个数的组合放置于map中，
	m := make(map[int]int, 1)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			m[nums[i]+nums[j]] = i*len(nums) + j
		}
	}

	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		if ms, ok := m[-nums[i]]; ok {
			//求 原来的两个下标
			first, second := ms/len(nums), ms%len(nums)
			if first == i || second == i {
				//说明重复了，不能算

				//还有一种可能性没考虑，就是这三个数已经在之前已经放进去了。因此解法应该和之前双数求和一样，边遍历，边放入
			} else {
				item := []int{i, first, second}
				res = append(res, item)
			}
		}
	}

	return res
}

func queen(a [8]int, cur int) {

	//最后一个皇后放好后，调用该方法时，cur ==8,即8个皇后都放好了，因此可以打印数组了
	if cur == len(a) {
		fmt.Print(a)
		fmt.Println()
		return
	}

	//i现在代表的是行数，cur代表列数
	for i := 0; i < len(a); i++ {
		//从第0行开始放
		a[cur] = i

		//代表这一行能不能放
		flag := true

		//验证可不可以放
		//后面几列还没放，所以不用考虑，j<cur
		for j := 0; j < cur; j++ {

			//i 代表本皇后放置的行号，j代表之前的列号，a[j]代表的是 j列的皇后放在了第几行
			//若果 a[j] == i 肯定不能放，flag = false，因为同一行有了两个了
			//ab代表的是两个皇后之间的行数差距，即纵向距离，如果纵向距离和横向距离相等的话，说明在一条斜线上
			ab := i - a[j]
			temp := 0
			//求纵向距离，变成正数，方便比较
			if ab > 0 {
				temp = ab
			} else {
				temp = -ab
			}
			//a[j] == i 处于同一行了，temp == cur-j 处于同一斜线了，该位置不能放
			if a[j] == i || temp == cur-j {
				flag = false
				break
			}
		}

		//满足的话，开始放下一列的皇后，不满足，行数往下走一行
		if flag {
			queen(a, cur+1)
		}
	}
}

//4皇后回溯法解答 [
// [".Q..",  // 解法 1
//  "...Q",
//  "Q...",
//  "..Q."],【】】
func solve4Queens() [][]string {
	res := make([][]string, 4)

	//棋盘
	nn := [][]int{}
	for i := 1; i < 4; i++ {
		nn[i] = make([]int, 4)
	}

	for i := 0; i < 4; i++ {
		//第一个可能的位置

	}

	return res
}

func main() {
	fmt.Println(solveNQueens(8))
}

//H皇后问题
func solveNQueens(n int) [][]string {
	//代表每一列中的皇后放置的行号
	res := &[][]string{}
	rows := make([]int, n)
	setNextQueen(0, rows, res)
	//回溯法 的思想：
	return *res
}

func setNextQueen(n int, rows []int, res *[][]string) {
	if n == len(rows) {
		resItem := make([]string, len(rows))
		for i := 0; i < len(rows); i++ {

			//生成“....Q......"
			strBytes := make([]byte, len(rows))
			for j := 0; j < len(rows); j++ {
				strBytes[j] = '.'
			}
			strBytes[rows[i]] = 'Q'

			resItem[i] = string(strBytes)
		}
		//加入到返回结果中
		*res = append(*res, resItem)
	} else {
		//看下一个放在哪,前n列都已经放好了
		for i := 0; i < len(rows); i++ {
			//先放0行
			rows[n] = i

			isOk := true
			//要判断行不行
			for j := 0; j < n; j++ {
				if rows[j] == i || //同一行了，不行
					(rows[j] > i && rows[j]-i == n-j) || //纵向距离与横向距离相等
					(rows[j] < i && i-rows[j] == n-j) {
					isOk = false
					break
				}
			}

			if isOk {
				setNextQueen(n+1, rows, res)
			}
		}
	}
}
