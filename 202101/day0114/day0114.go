package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

}

//537. 复数乘法
func complexNumberMultiply(a string, b string) string {
	//(x+yi)*(c*di)
	indexAPlus := strings.Index(a, "+")
	indexAI := strings.Index(a, "i")
	indexBPlus := strings.Index(b, "+")
	indexBI := strings.Index(b, "i")
	x, _ := strconv.Atoi(a[:indexAPlus])
	y, _ := strconv.Atoi(a[indexAPlus+1 : indexAI])
	c, _ := strconv.Atoi(b[:indexBPlus])
	d, _ := strconv.Atoi(b[indexBPlus+1 : indexBI])
	return fmt.Sprint((x*c - y*d), "+", (y*c + x*d), "i")
}

//每日一题：1018. 可被 5 整除的二进制前缀
func prefixesDivBy5(A []int) []bool {
	res := make([]bool, len(A))
	preNum := 0
	for i := 0; i < len(A); i++ {
		curNum := preNum<<1 + A[i]
		preNum = curNum % 5
		if preNum == 0 {
			res[i] = true
		}
	}
	return res
}
