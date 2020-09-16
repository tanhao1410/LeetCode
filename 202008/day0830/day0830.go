package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println(reverseWords2("ta"), "--", "_")
	//nums := [][]int{{5, 5}, {9, 4}, {9, 7}, {6, 4}, {7, 0},{9, 5}, {10, 7}, {1, 1}, {7, 5}}
	//command := "RRU"
	//fmt.Println(robot(command,nums,1486,743))
	fmt.Println(primePalindrome(100000000))
	fmt.Println(count)
	//fmt.Println(IsPalinrome2(22))
}

//判断是否存在连续三个数都是奇数
func threeConsecutiveOdds(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 1 {
			//是奇数
			if count++; count == 3 {
				return true
			}
		} else {
			count = 0
		}
	}
	return false
}

//验证IP地址IPv4 IPv6 Neither
func validIPAddress(IP string) string {

	return ""
}

//大于等等N的最小回文素数
func primePalindrome(N int) int {
	//先判断是否是回文，再判断是否是素数
	for ; ; N++ {
		if IsPalinrome2(N) && IsSushu(N) {
			return N
		}
	}
}

//找出下一个回文数
func NextPalinrome(n int) int {
	//看最低位与高位，如果小于，则，直接增长最低位即可。
	//大于等于的话，说明只能通过增长十位来解决了
	//都不行的话，n需要扩大一位即1000001
	ns := []byte(strconv.Itoa(n))
	for i, j := 0, len(ns)-1; i < j; {
		//最高位的数 ns[i]
		if ns[i] > ns[j] {
			ns[j] = ns[i]
			res, _ := strconv.Atoi(string(ns))
			return res
		} else {
			ns[j] = 0
			ns[j-1]++
		}
	}

	res := 1
	for i := 0; i < len(ns); i++ {
		res = res * 10
	}
	return 1
}

func IsSushu(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//判断是否是回文数
func IsPalinrome(n int) bool {
	//1.简单思路，转成字符串来判断
	ns := strconv.Itoa(n)
	for i, j := 0, len(ns)-1; i < j; i++ {
		if ns[i] != ns[j] {
			return false
		}
		j--
	}
	return true
	//2.

}

var count int = 0

func IsPalinrome2(n int) bool {
	//2.判断是否是回文数，
	res := 0
	for i := n; i > 0; i = i / 10 {
		res = res*10 + i%10
	}
	count++
	return res == n

}

//表示数值的字符串 e12 false,1. true .3true .e3false .1e3 true
func isNumber(s string) bool {
	//1.开头可以是数字，后面必须都是数字，直到小数点或E为止,
	//2.正负号，后面必须都是数字，直到小数点或E为止
	//3.. 后面必须是数，e/E,后面必须是正负整数
	if len(s) == 0 {
		return false
	}

	//判断第一位
	if s[0] == '.' {
		//第一位是. 在遇到e之前一定是个 正整数
		for i := 1; i < len(s); i++ {

		}
	} else if s[0] == '+' || s[0] == '-' {
		//第一位是+、-
		//在遇到e之前一定是个 无符号数

	} else if IsNum(s[0]) {
		//第一位是个数字，在遇到e之前一定是个 无符号数
	}
	return false
}

//判断是不是数
func IsNum(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}

//判断是不是数或正负号
func IsNumOrSign(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	if b == '+' || b == '-' {
		return true
	}
	return false
}

//判断是不是整数
func IsZhengshu(s string) bool {
	if len(s) == 0 || IsNumOrSign(s[0]) { //判断第一位
		return false
	}

	for i := 1; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

//判断是否是没有+-号的整数
func IsZhengZhengShu(s string) bool {

	for _, v := range s {
		if !IsNum(byte(v)) {
			return false
		}
	}

	return true
}

func robot2(command string, obstacles [][]int, x int, y int) bool {

	return false
}

//反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序
func reverseWords(s string) string {
	//思路：简单思路，按空格切分
	ss := strings.Split(s, " ")
	//反转每一个字符串，
	res := ""
	for _, v := range ss {
		for i := len(v) - 1; i >= 0; i-- {
			res += string(v[i])
		}
		res += " "
	}
	return strings.TrimRight(res, " ")
}

func reverseWords2(s string) string {
	//思路：双指针思路，一个指向单词开头，一个指向结尾，然后交换位置。开始下一个单词
	ss := []byte(s)

	for i, j := 0, 1; i < len(s); {
		//i不动，j向后走至空格或结尾处
		for ; j < len(s) && s[j] != ' '; j++ {
		}
		temp := j //空格的位置,或结束的位置
		//开始交换位置
		for i < j && j <= len(s) {
			ss[i], ss[j-1] = s[j-1], s[i]
			i, j = i+1, j-1
		}
		//交换完成后，i需要跳转至
		i, j = temp+1, temp+2
	}
	return string(ss)
}

//"RRU"
//[[5, 5], [9, 4], [9, 7], [6, 4], [7, 0], [9, 5], [10, 7], [1, 1], [7, 5]]
//1486
//743
//UR:上和右移动，给定终点x,y能否完整的，指令是无限循环指令
func robot(command string, obstacles [][]int, x int, y int) bool {
	//思路：每走一步，得到一个坐标，先看有没有超出目标地或到达目标地，然后看坐标是否在障碍物中，
	//思路2，走过一个循环后，其实后面的路线都算很明确了

	//这个map不行，
	//m:= make(map[int]int, len(obstacles))
	//for i := 0;i < len(obstacles);i++{
	//	m[obstacles[i][0]] = obstacles[i][1]
	//}
	m := make(map[int][]int)
	for i := 0; i < len(obstacles); i++ {
		if ms, ok := m[obstacles[i][0]]; ok {
			//把新的添加进来
			ms = append(ms, obstacles[i][1])
		} else {
			m[obstacles[i][0]] = []int{obstacles[i][1]}
		}
	}
	//fmt.Println(obstacles[0][1])
	x1, y1 := 0, 0 //目前的位置
	for {
		for i := 0; i <= len(command)-1; i++ {
			if command[i] == 'U' { //向上
				y1++
			} else { //向下
				x1++
			}
			if x1 > x || y1 > y {
				return false
			}
			if x1 == x && y1 == y {
				return true
			}
			//判断x1,y1,是否在障碍物中
			if obstaclesY, ok := m[x1]; ok {
				for _, v := range obstaclesY {
					if v == y1 {
						return false
					}
				}
			}
		}
	}
}
