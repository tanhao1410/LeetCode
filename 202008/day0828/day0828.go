package main

import (
	"math"
	"strings"
)

func main() {
	//fmt.Println(myAtoi("   0000000000012345678"))
	letterCombinations("223")
}

//电话号码的组合，n重for循环
func letterCombinations(digits string) []string {

	var res []string = []string{}

	var m map[string][]string = map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"x", "y", "z"},
	}
	//第一个字母，第二个字母，。。。。第n个字母
	CreateLetter("", digits, &res, m)
	//for _,v := range res{
	//	fmt.Println(v)
	//}

	return res
}

func CreateLetter(pre string, tail string, res *[]string, m map[string][]string) {
	if len(tail) == 1 { //说明是最后一位了
		//根据tail获得对应的字母
		cs := m[tail]
		for i := 0; i < len(cs); i++ {
			s := pre + cs[i]
			*res = append(*res, s)
		}
	} else {
		//不是最后一位
		cs := m[tail[0:1]]
		for i := 0; i < len(cs); i++ {
			s := pre + cs[i]
			CreateLetter(s, tail[1:], res, m)
		}
	}
}

//用常数空间解决
//用第0行0列来记录
func setZeroes2(matrix [][]int) {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	//先判断第0行0列是否出现过0
	colFlag := false
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			//出现过
			colFlag = true
			break
		}
	}
	rowFlag := false
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			rowFlag = true
			break
		}
	}

	//遍历矩阵
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ { //需要考虑matrix[0]不存在
			if matrix[i][j] == 0 {
				//fmt.Println("....")
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	//设0
	for k, v := range matrix[0] {
		if k == 0 {
			continue
		}
		//设置值
		if v == 0 {
			//fmt.Println("列",k)
			//整列置零
			for i := 1; i < len(matrix); i++ {
				matrix[i][k] = 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			//整行置零
			for j := 1; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}

	if rowFlag {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}

	if colFlag {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}

//矩阵置零
//m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法
func setZeroes(matrix [][]int) {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	//难点在于，直接置零会影响其他的判断。
	//记录哪些行和列应该被清空，
	mRow := make(map[int]int, 1)
	mCol := make(map[int]int, 1)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ { //需要考虑matrix[0]不存在
			if matrix[i][j] == 0 {
				mRow[i] = i
				mCol[j] = j
			}
		}
	}
	for _, v := range mRow {
		//设置值
		for i := 0; i < len(matrix[0]); i++ {
			matrix[v][i] = 0
		}
	}

	for _, v := range mCol {
		//设置值
		for i := 0; i < len(matrix); i++ {
			matrix[i][v] = 0
		}
	}

	//用常数空间解决
}

func myAtoi(str string) int {
	//先去前面空格
	//for  ; len(str)>0;{
	//	if str[0] == ' '{
	//		str = str[1:]
	//	}else{
	//		break
	//	}
	//}
	str = strings.TrimLeft(str, " ")

	//看正负号
	isZ := 1
	if len(str) > 0 {
		if str[0] == '-' {
			isZ = -1
			str = str[1:]
		} else if str[0] == '+' {
			str = str[1:]
		}
	}

	//处理数
	res := 0
	for i := 0; i < len(str) && str[i] <= '9' && str[i] >= '0'; i++ {
		res = res*10 + int(str[i]-'0')
		//判断是否越界 res最大为2^32-1，但负数 最大为2^32，在java中即使是越界了，也不会变成小于0的
		if res > math.MaxInt32 && isZ == 1 {
			return math.MaxInt32
		}
		//2147483648
		if res > math.MaxInt32 && isZ == -1 {
			return math.MinInt32
		}
	}

	return res * isZ
}

//相同的元素 nums = [0,1,2,2,3,0,4,2]
func removeElement(nums []int, val int) int {

	i, j := 0, len(nums)-1
	//j指向最后一个不为val的数
	for ; j >= 0 && nums[j] == val; j-- {
	}

	for ; i < len(nums) && i < j; i++ {
		if nums[i] == val {
			//交换
			nums[i], nums[j] = nums[j], val
			//j需要往前走，指向最后一个不为val的数
			for ; j >= 0 && nums[j] == val; j-- {
			}
		}
	}
	return j
}
