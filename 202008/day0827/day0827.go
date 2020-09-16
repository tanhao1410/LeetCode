package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

//验证有效数独
func isValidSudoku2(board [][]byte) bool {
	//只是验证
	//思路：1.每一行都不重，2.每一列都不重，3.每一个9方格不重

	//验证行
	for i := 0; i < 9; i++ {
		nums := [10]int{0}
		//验证一行
		for _, v := range board[i] {
			if v == '.' {
				continue
			}
			if nums[v-'0'] != 0 { //说明重复了
				return false
			}
			nums[v-'0'] = 1
		}
	}

	//验证列
	for i := 0; i < 9; i++ {
		nums := [10]int{0}
		for k := 0; k < 9; k++ {
			if board[k][i] == '.' {
				continue
			}
			if nums[board[k][i]-'0'] != 0 {
				return false
			}
			nums[board[k][i]-'0'] = 1
		}
	}
	//验证9个9方格
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			nums := [10]int{0}
			for k := 0; k < 9; k++ {
				n := board[i*3+j][(k%3)+j*3]
				if n == '.' {
					continue
				}
				if nums[n-'0'] != 0 {
					return false
				}
				nums[n-'0'] = 1
			}
		}
	}

	return true
}

func isValidSudoku(board [][]byte) bool {
	//只是验证
	//思路：1.每一行都不重，2.每一列都不重，3.每一个9方格不重

	//验证行
	for i := 0; i < 9; i++ {
		nums := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		//验证一行
		for _, v := range board[i] {
			if !Check(v, nums) {
				return false
			}
		}
	}

	//验证列
	for i := 0; i < 9; i++ {
		nums := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		for k := 0; k < 9; k++ {
			if !Check(board[k][i], nums) {
				return false
			}
		}
	}
	//验证9个9方格
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			nums := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
			for k := 0; k < 9; k++ {
				n := board[i*3+k/3][(k%3)+j*3]
				if !Check(n, nums) {
					return false
				}
			}
		}
	}

	return true
}

func Check(b byte, nums []int) bool {
	if b == '.' {
		return true
	}
	if nums[b-'0'] != 0 {
		fmt.Println(b - '0')
		return false
	}
	nums[b-'0'] = 1
	return true
}

func main() {
	fmt.Println(divide(1110, -1))
}

//两数相除，将两数相除，要求不使用乘法、除法和 mod 运算符。
func divide(dividend int, divisor int) int {

	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	//先不考虑正负
	flag := true
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		flag = false
	}

	//都转成整数
	dendS := strings.ReplaceAll(strconv.Itoa(dividend), "-", "")
	sorS := strings.ReplaceAll(strconv.Itoa(divisor), "-", "")
	dividend, _ = strconv.Atoi(dendS)
	divisor, _ = strconv.Atoi(sorS)

	if len(sorS) > len(dendS) {
		return 0
	}

	shang := ""
	yushu, _ := strconv.Atoi(dendS[0 : len(sorS)-1]) //第一次的时候取同样长度的数
	for i := 0; len(sorS)+i <= len(dendS); i++ {

		//要考虑上一步留下来的余数
		if yushu != 0 {
			yushu, _ = strconv.Atoi(strconv.Itoa(yushu) + dendS[len(sorS)-1+i:len(sorS)+i])
		} else {
			yushu, _ = strconv.Atoi(dendS[len(sorS)-1+i : len(sorS)+i])
		}

		//先算高位,如果高位大于除数，则商1--->不一定是1
		if yushu >= divisor {
			j := 0
			for ; yushu >= divisor; j++ {
				yushu -= divisor //每次减
			}
			//h为余数，参与下次运算，j为商
			shang = shang + strconv.Itoa(j)
		} else {
			if len(shang) > 0 { //说明 不是第一次了
				shang = shang + "0"
			}
		}
	}

	if len(shang) == 0 {
		return 0
	}

	if !flag {
		shang = "-" + shang
	}
	res, _ := strconv.Atoi(shang)
	return res

}

//有效的括号'('，')'，'{'，'}'，'['，']'
func isValid(s string) bool {
	//初始化一个栈
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			//进栈
			stack = append(stack, s[i])
		} else {
			//出栈
			if len(stack) == 0 {
				return false
			}
			switch s[i] {
			case ')':
				if stack[len(stack)-1] != '(' {
					return false
				}
				stack = stack[:len(stack)-1]
			case '}':
				if stack[len(stack)-1] != '{' {
					return false
				}
				stack = stack[:len(stack)-1]
			case ']':
				if stack[len(stack)-1] != '[' {
					return false
				}
				stack = stack[:len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}
