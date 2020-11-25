package main

func main() {

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
