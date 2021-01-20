package main

func main() {

}

//每日一题：628. 三个数的最大乘积
func maximumProduct(nums []int) int {

	//全正数-最大的三个数，全负数-最大的三个数，一个正数其他为负-最小的两个负数*最大的正数，就一个负数-最大的三个数，多个正多个负-最小的两个负*最大的正或 最大的三个数
	//分析可知，记录最大的三个数，最小的两个数，共五个数
	fir, sec, thi := -10000, -10000, -10000
	min, min2 := 10000, 10000
	for _, v := range nums {
		//记录最大的三个数
		if v > fir {
			fir, sec, thi = v, fir, sec
		} else if v > sec {
			sec, thi = v, sec
		} else if v > thi {
			thi = v
		}

		//记录最小的两个数
		if v < min {
			min, min2 = v, min
		} else if v < min2 {
			min2 = v
		}

	}
	max1 := fir * sec * thi
	max2 := fir * min * min2
	if max1 > max2 {
		return max1
	}
	return max2
}
