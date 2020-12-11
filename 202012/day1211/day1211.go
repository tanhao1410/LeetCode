package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(predictPartyVictory("DDRRRR"))
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
