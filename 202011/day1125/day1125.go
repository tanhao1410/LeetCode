package main

import "fmt"

func main() {

}

//299. 猜数字游戏
func getHint(secret string, guess string) string {
	//思路：先记录全对的个数。然后统计其他数字的个数，找共同的个数即可
	a, b := 0, 0
	m1, m2 := make([]int, 10), make([]int, 10)
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			a++
		} else {
			m1[secret[i]-'0'] += 1
			m2[guess[i]-'0'] += 1
		}
	}
	for i := 0; i < 10; i++ {
		if m1[i] < m2[i] {
			b += m1[i]
		} else {
			b += m2[i]
		}
	}

	return fmt.Sprintf("%s%s%s%s", a, "A", b, "B")
}

//每日一题：1370. 上升下降字符串
func sortString(s string) string {
	res := ""
	//思路：用一个26个大小的数组，记录所有的字母的数量
	m := make([]int, 26)
	for _, v := range s {
		m[v-'a'] += 1
	}
	//开始拼接字符串
	//从中选择最小的
	selectMin := func() {
		for i := 0; i < 26; i++ {
			if m[i] > 0 {
				m[i]--
				res += string('a' + i)
			}
		}
	}
	selectMax := func() {
		for i := 25; i >= 0; i-- {
			if m[i] > 0 {
				m[i]--
				res += string('a' + i)
			}
		}
	}
	for len(res) < len(s) {
		//先选最小的
		selectMin()
		selectMax()
	}
	return res
}
