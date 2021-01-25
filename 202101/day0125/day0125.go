package main

func main() {

}

//面试题 08.05. 递归乘法
func multiply(A int, B int) int {
	if A == 1 {
		return B
	}
	if B == 1 {
		return A
	}
	if A > B {
		return multiply(A, B-1) + A
	}
	return multiply(A-1, B) + B
}

//每日一题：959. 由斜杠划分区域
func regionsBySlashes(grid []string) int {

	return 0
}
