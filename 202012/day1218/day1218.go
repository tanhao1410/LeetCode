package main

func main() {

}

//每日一题：389. 找不同
func findTheDifference(s string, t string) byte {
	m := make([]int, 26)
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']--
		m[t[i]-'a']++
	}
	m[t[len(s)]-'a']++
	for i := 0; ; i++ {
		if m[i] > 0 {
			return byte(i + 'a')
		}
	}
}
