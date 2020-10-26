package main

func main() {
	//nums := []int{5, 0, 10, 0, 10, 6}
	//print(smallerNumbersThanCurrent(nums))
	matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(matrix)
}

//48.旋转图像
func rotate(matrix [][]int) {
	//先旋转外层，再递归旋转里层，直到里层为空或1
	l := len(matrix)
	if l < 2 {
		return
	}

	//列比行多
	if l < len(matrix[0]) {

		//行和原来一样，取列的时候，都加上一个数，
		n := (len(matrix[0]) - l) / 2
		for i := 0; i < l-1; i++ {
			matrix[0][i+n], matrix[l-1-i][n], matrix[l-1][l-1-i+n], matrix[i][l-1+n] =
				matrix[l-1-i][n], matrix[l-1][l-1-i+n], matrix[i][l-1+n], matrix[0][i+n]
		}
	} else {
		//旋转外层
		for i := 0; i < l-1; i++ {
			matrix[0][i], matrix[l-1-i][0], matrix[l-1][l-1-i], matrix[i][l-1] =
				matrix[l-1-i][0], matrix[l-1][l-1-i], matrix[i][l-1], matrix[0][i]
		}
	}

	//旋转内层
	if l > 3 {
		//抽取行
		miniMatrix := matrix[1 : l-1]
		rotate(miniMatrix)
	}

}

//1365.有多少小于当前数字的数字
func smallerNumbersThanCurrent(nums []int) []int {
	//思路，用一个数组统计每个数字出现的次数，
	//从小到大，一次遍历，得到每个数字前面有多少小于自己的数的数组
	//改变入参的数组，返回即可
	numCount := make([]int, 101)
	for _, v := range nums {
		numCount[v] += 1
	}

	//比当前数小的=比前一个数小的数量+前一个数的数量
	preCount := numCount[0]
	numCount[0] = 0
	for i := 1; i < 101; i++ {
		numCount[i], preCount = numCount[i-1]+preCount, numCount[i]
	}

	//改变参数，返回结果
	for i := 0; i < len(nums); i++ {
		nums[i] = numCount[nums[i]]
	}

	return nums
}
