package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println(calculate("3 *1 /2"))
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 180}))
}

//300. 最长上升子序列
func lengthOfLIS(nums []int) int {
	//以自己为结尾的最大子序列
	res := 0
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		if res < dp[i] {
			res = dp[i]
		}
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				if dp[i] > res {
					res = dp[i]
				}
			}
		}
	}
	return res
}

//227. 基本计算器 II
func calculate(s string) int {
	//仅包含+-*/
	//思路：注意顺序即可
	//去除所有空格，
	s1 := strings.ReplaceAll(s, " ", "")
	ss := []string{}
	start := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] == '+' || s1[i] == '-' || s1[i] == '*' || s1[i] == '/' {
			//需要截取数字了
			ss = append(ss, s1[start:i])
			ss = append(ss, s1[i:i+1])
			start = i + 1
		}
	}
	ss = append(ss, s1[start:])
	stack := []string{ss[0]}
	for i := 1; i < len(ss); i++ {
		cur := ss[i]
		//看栈顶的情况
		top := stack[len(stack)-1]
		if top == "*" {
			//取出来栈的两个元素，第一个为*号，第二为应有的数字
			num, _ := strconv.Atoi(stack[len(stack)-2])
			curNum, _ := strconv.Atoi(cur)
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = strconv.Itoa(num * curNum)
			//结果入栈
		} else if top == "/" {
			num, _ := strconv.Atoi(stack[len(stack)-2])
			curNum, _ := strconv.Atoi(cur)
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = strconv.Itoa(num / curNum)
		} else {
			stack = append(stack, cur)
		}
	}

	res, _ := strconv.Atoi(stack[0])
	//循环结束后，栈中只剩下+-了
	for i := 1; i < len(stack); i++ {
		if stack[i] == "+" {
			nextNum, _ := strconv.Atoi(stack[i+1])
			res += nextNum
		} else if stack[i] == "-" {
			nextNum, _ := strconv.Atoi(stack[i+1])
			res -= nextNum
		}
	}
	return res
}

//每日一题：283. 移动零
func moveZeroes(nums []int) {
	//思路：遇到不是0的数就往前移动，移动到前面第一个0的位置处用一个变量来记录
	next := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 && next > -1 {
			nums[next], nums[i] = nums[i], 0
			next += 1
		} else if nums[i] == 0 && next == -1 {
			next = i
		}
	}
}
