package main

func main() {
	
}
//每日一题：242. 有效的字母异位词
func isAnagram(s string, t string) bool {
	s1,s2 := make([]int,27),make([]int,27)
	for _,v := range s{
		s1[v - 'a'] += 1
	}
	for _,v := range t{
		s2[v - 'a'] += 1
	}
	for i :=0;i < 27;i ++{
		if s1[i] != s2[i]{
			return false
		}
	}
	return true
}
