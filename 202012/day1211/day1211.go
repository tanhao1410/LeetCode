package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(predictPartyVictory("DDRRRR"))
	fmt.Println(isScramble("dbdac", "abcdd"))
}

//87. 扰乱字符串
func isScramble(s1 string, s2 string) bool {
	//思路：各分成连续的两段，其中需要s1中的 一段对应s2中的一段，保证相同（不要求顺序相同），递归，如果都是可以的，那么总的就是可以的
	//分成两段
	//如果长度小于4，并且字母相同，那么肯定属于扰乱字符串
	if !isSame(s1, s2) {
		return false
	}
	if isSame(s1, s2) && len(s1) < 4 {
		return true
	}
	i := 1
	for ; i < len(s1); i++ {
		if isSame(s1[:i], s2[:i]) && isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) { //可以交叉相等的
			return true
		}
		if isSame(s1[:i], s2[len(s2)-i:]) && isScramble(s1[:i], s2[len(s2)-i:]) && isScramble(s1[i:],
			s2[:len(s2)-i]) {
			return true
		}
	}
	return false
}

//判断是否是相同字符串，不要求顺序
func isSame(s1 string, s2 string) bool {
	m := make([]int, 26)
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		m[s1[i]-'a'] += 1
		m[s2[i]-'a'] -= 1
	}
	for i := 0; i < 26; i++ {
		if m[i] > 0 {
			return false
		}
	}
	return true
}

//391. 完美矩形
func isRectangleCover(rectangles [][]int) bool {
	//思路：找 合在一起的左下和右上，计算面积
	a, b, c, d := math.MaxInt32, math.MaxInt32, 0, 0
	for _, v := range rectangles {
		if v[0] < a {
			a = v[0]
		}
		if v[1] < b {
			b = v[1]
		}
		if v[2] > c {
			c = v[2]
		}
		if v[3] > d {
			d = v[3]
		}
	}
	all := (d - b) * (c - a)
	//计算分别的面积。若不相同，肯定不是的
	for _, v := range rectangles {
		all -= (v[3] - v[1]) * (v[2] - v[0])
	}
	if all != 0 {
		return false
	}
	//若相同，判断是否存在重叠的，有重叠的，返回false
	//某个矩形的左下角，或左上角落在了别的矩形范围内（不包括右下边界）
	for k, v := range rectangles {
		for kk, vv := range rectangles {
			if kk == k {
				continue
			}
			//v的左上角为 v[0],v[3] 落在了vv的矩形范围内
			if v[0] >= vv[0] && v[0] < vv[2] && v[3] <= vv[3] && v[3] > vv[1] {
				return false
			}
			//v的左下角为 v[0],v[1]
			if v[0] >= vv[0] && v[0] < vv[2] && v[1] >= vv[1] && v[1] < vv[3] {
				return false
			}
			//完全包含
			if v[0] >= vv[0] && v[1] >= vv[1] && v[2] <= vv[2] && v[3] <= vv[3] {
				return false
			}
		}
	}
	return true
}

//每日一题：649. Dota2 参议院
func predictPartyVictory(senate string) string {
	//RD
	r, d := 0, 0
	remain := ""
	for _, k := range senate {
		if k == 'R' {
			//如果是R,先看这个R是否被取消资格了，如果被取消资格了，直接跳过即可。
			//d记录的是被取消资格的R的量
			if d > 0 {
				d--
			} else {
				r++
				//取消一个D的资格
				remain += "R"
			}
		} else {
			if r > 0 {
				r--
			} else {
				d++
				remain += "D"
			}
		}
	}
	//删除r个D
	remain = strings.Replace(remain, "D", "", r)
	remain = strings.Replace(remain, "R", "", d)
	if strings.Contains(remain, "R") && !strings.Contains(remain, "D") {
		return "R"
	} else if !strings.Contains(remain, "R") && strings.Contains(remain, "D") {
		return "D"
	}
	return predictPartyVictory(remain)
}
