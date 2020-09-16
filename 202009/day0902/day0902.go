package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isNumber("12e+5.4"))
}

//表示数值的字符串：采用最原始的方法
func isNumber(s string) bool {
	s = strings.ReplaceAll(s, "E", "e")
	s = strings.Trim(s, " ")

	//思路：1.先看e/E,如果有它的话，后面必须要跟一个整数 测试 e+2 属于,e0属于。前面必须要有数，不能为空 0e2属于。
	eIndex := strings.Index(s, "e")
	if eIndex == len(s)-1 {
		return false
	}

	if eIndex != -1 {
		eAferS := s[eIndex+1:]
		if !isZhengShu(eAferS, true) {
			return false //e后面必须跟个数 可以带正负号
		}
		s = s[:eIndex]
	}

	//2.看点：在去掉e之后的所有的字符中查看。.前后必须至少一个有数。判断.之后，只能有整数或空，注意 空的话，.的前面必须有数字
	//看是否存在小数点
	pointIndex := strings.Index(s, ".")
	pointPreHaveNum := false
	if pointIndex != -1 {
		//判断.后的情况
		pointAfterS := s[pointIndex+1:]
		if pointAfterS == "" {
			//点之后后为空，那么点之前必须有数字
			pointPreHaveNum = true
		} else {
			if !isZhengShu(pointAfterS, false) {
				return false
			}
		}
		s = s[:pointIndex]
	} else {
		//没有点
		pointPreHaveNum = true //没有点的话，前面也必须有数字了
	}

	//3.去掉.之后，看前面是否符合整数或空。空的话，若点后有数，可以的。
	if s == "" || s == "+" || s == "-" {
		if pointPreHaveNum {
			//如果.之前要求必须有数，但没有
			return false
		} else {
			return true
		}
	}

	if !isZhengShu(s, true) {
		return false // 如果前面不为数的话，返回false
	}
	return true
}

//是否包括isContainSign符号位
func isZhengShu(s string, isContainSign bool) bool {
	if len(s) == 0 {
		return false
	}
	if isContainSign && (s[0] == '-' || s[0] == '+') { //可能带有符号位
		//看第一位是否是+-号
		if len(s) == 1 {
			return false
		}
		s = s[1:]
	}

	for _, v := range s {
		if v > '9' || v < '0' {
			return false
		}
	}
	return true
}
