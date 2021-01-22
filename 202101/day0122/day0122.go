package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findLongestWord("aaa", []string{"abc", "bcd", "abd", "dafae", "dafee"}))

}

//524. 通过删除字母匹配到字典里最长单词
func findLongestWord(s string, d []string) string {
	//思路：字典中字符串先按长度排序，同长度按字母序，按顺序一个个判断是否能取出来
	stractStr := func(s2 string) bool {
		for i, j := 0, 0; j < len(s2); j++ {
			curLetter := s2[j]
			for ; i < len(s) && s[i] != curLetter; i++ {
			}
			if i >= len(s) {
				return false
			}
			i++
		}
		return true
	}

	sort.Slice(d, func(i, j int) bool {
		//先比较长度
		if len(d[i]) != len(d[j]) {
			return len(d[i]) > len(d[j])
		} else {
			for k := 0; k < len(d[i]); k++ {
				if d[i][k] != d[j][k] {
					return d[i][k] < d[j][k]
				}
			}
			return true
		}
	})

	for i := 0; i < len(d); i++ {
		if stractStr(d[i]) {
			return d[i]
		}
	}
	return ""
}

//每日一题：989. 数组形式的整数加法
func addToArrayForm(A []int, K int) []int {
	//思路：低位和K相加，看是否有进位，如果有进位，就处理。没有，直接处理后面的即可
	num2Array := func(num int) []int {
		if num == 0 {
			return []int{0}
		}
		res := []int{}

		for num != 0 {
			res = append(res, num%10)
			num = num / 10
		}

		//逆转
		for i := 0; i < len(res)/2; i++ {
			res[i], res[len(res)-1-i] = res[len(res)-i-1], res[i]
		}

		return res
	}

	//先取五位
	if len(A) >= 5 {
		lowNum := A[len(A)-5]*10000 + A[len(A)-4]*1000 + A[len(A)-3]*100 + A[len(A)-2]*10 + A[len(A)-1] + K

		A[len(A)-5] = lowNum / 10000
		A[len(A)-4] = (lowNum % 10000) / 1000
		A[len(A)-3] = (lowNum % 1000) / 100
		A[len(A)-2] = (lowNum % 100) / 10
		A[len(A)-1] = lowNum % 10

		if lowNum >= 100000 { //说明产生了进位
			i := len(A) - 6
			for ; i >= 0; i-- {
				if A[i] == 9 {
					A[i] = 0
				} else {
					A[i]++
					break
				}
			}
			if i < 0 {
				//说明总的已经超过位数
				return append([]int{1}, A...)
			}
		}
		return A
	} else {

		num := 0
		for i := 0; i < len(A); i++ {
			num = num*10 + A[i]
		}
		return num2Array(num + K)

	}

}
