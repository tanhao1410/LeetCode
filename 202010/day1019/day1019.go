package main

func main() {

}

//每日一题：844.比较含退格的字符串
func backspaceCompare(S string, T string) bool {
	//思路1：都转换成字符串，判断是否相等即可。
	getRealString := func(s string)  string{
		s1 := []byte{}
		for i := 0; i < len(s); i++ {
			switch s[i] {
			case '#':
				if len(s1) > 0 {
					s1 = s1[:len(s1)-1]
				}
			default:
				s1 = append(s1, s[i])
			}
		}
		return string(s1)
	}

	return getRealString(S) == getRealString(T)
}
