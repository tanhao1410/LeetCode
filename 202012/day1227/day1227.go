package main

func main() {

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
