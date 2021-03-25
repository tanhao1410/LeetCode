package main

func main() {

}

//面试题 08.06. 汉诺塔问题
func hanota(A []int, B []int, C []int) []int {
	//思路：当盘子中就一个时，直接放，如果是两个，先先把所有的，处最后一个外，放在中间的，剩下的这个放到C里
	if len(A) == 0 {
		return C
	}

	for len(A) > 1 {
		B = append(B, A[len(A)-1])
		A = A[:len(A)-1]
	}
	C = append(C, A[0])
	A = A[:0]
	for len(B) > 0 {
		A = append(A, B[len(B)-1])
		B = B[:len(B)-1]
	}
	return hanota(A, B, C)
}
