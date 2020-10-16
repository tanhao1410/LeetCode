package main

func main() {

}

//977.有序数组的平方
func sortedSquares(A []int) []int {
	res := make([]int,len(A))
	for start,end ,i:= 0,len(A) -1,len(A) -1;start <= end;i--{
		start2 := A[start] * A[start]
		end2 := A[end] * A[end]
		if start2 > end2{
			res[i] = start2
			start++
		}else{
			res[i] = end2
			end--
		}
	}
	return res
}