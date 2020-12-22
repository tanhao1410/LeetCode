package main

func main() {

}

//每日一题：387. 字符串中的第一个唯一字符
func firstUniqChar(s string) int {
	m := make([]int, 26)
	for i := len(s) - 1; i >= 0; i-- {
		m[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if m[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
