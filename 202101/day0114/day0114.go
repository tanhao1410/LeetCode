package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
}

//540. 有序数组中的单一元素
func singleNonDuplicate(nums []int) int {
	//思路：二分法，每个数字都出现两次，唯独一个是单个的，
	//从中间划开，若中间的数和左边相等，说明在左边，若和右边相等，说明在右边，否则返回该数
	start, end := 0, len(nums)-1
	for middle := end / 2; ; middle = (start + end) / 2 {

		//第一，看右边
		//如果右边的数个数是偶数个，且middle != middle + 1，那么肯定不会在这边的了
		//如果右边的数个数是奇书个，且middle == middle + 1,那么肯定不会在这边

		//先判断是否和两边都不相等
		if middle+1 < len(nums) && nums[middle] == nums[middle+1] {
			//判断后面数的个数
			count := end - middle - 1
			if count%2 == 0 {
				//在左边
				end = middle - 1
			} else {
				start = middle + 2
			}
		} else if middle-1 >= 0 && nums[middle] == nums[middle-1] {
			count := middle - 1
			if count%2 == 0 {
				//在右边
				start = middle + 1
			} else {
				end = middle - 2
			}
		} else {
			return nums[middle]
		}
	}
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
