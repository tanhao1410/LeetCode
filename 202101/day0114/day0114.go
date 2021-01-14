package main

func main() {

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
