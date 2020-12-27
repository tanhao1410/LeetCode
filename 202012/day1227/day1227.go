package main

func main() {

}

//509. 斐波那契数
func fib(n int) int {
	if n < 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

//520. 检测大写字母
func detectCapitalUse(word string) bool {

	//思路：先看第一个字母是否是大写，如果是，除非后面全大写，或全小写，或为空
	//如果第一个字母是小写，则后面必须是全小写
	firstIsSmall := word[0] >= 'a' && word[0] <= 'z'
	haveSmall, haveBig := false, false
	for i := 1; i < len(word); i++ {
		if word[i] >= 'a' && word[i] <= 'z' {
			haveSmall = true
		}
		if word[i] >= 'A' && word[i] <= 'Z' {
			haveBig = true
		}
		//第一个字母是小写，后面出现了大写
		if firstIsSmall && haveBig {
			return false
		}
		//第一个字母是大写，后面又出现小写，又出现大写
		if !firstIsSmall && haveBig && haveSmall {
			return false
		}

	}
	return true
}

//每日一题：205. 同构字符串
func isIsomorphic(s string, t string) bool {
	m1 := make(map[byte]byte)
	m2 := make(map[byte]byte)
	for k, v := range s {
		if vv, ok := m1[byte(v)]; ok {
			if vv != t[k] {
				return false
			}
		} else {
			//m1[byte(v)] = t[k]
			if _, ok := m2[t[k]]; ok {
				return false
			} else {
				m1[byte(v)] = t[k]
				m2[t[k]] = byte(v)
			}
		}
	}

	return true
}
