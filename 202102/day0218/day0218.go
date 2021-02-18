package main

import (
	"math"
	"math/rand"
	"time"
)

func main() {

}

//496. 下一个更大元素 I
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	//先用一个map记录所有的位置
	m := make(map[int]int)
	for k, v := range nums2 {
		m[v] = k
	}
	//记录所有数的下一个更大元素
	dp := make([]int, len(nums2))
	for i := 0; i < len(nums2); i++ {
		if i > 0 && nums2[i] > nums2[i-1] && dp[i-1] > nums2[i] {
			dp[i] = dp[i-1]
		} else {
			for j := i + 1; j < len(nums2); j++ {
				if nums2[j] > nums2[i] {
					dp[i] = nums2[j]
					break
				}
			}
			if dp[i] == 0 {
				dp[i] = -1
			}
		}
	}

	res := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		res[i] = dp[m[nums1[i]]]
	}

	return res
}

//478. 在圆内随机生成点
type Solution struct {
	x float64
	y float64
	r float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{
		x: x_center,
		y: y_center,
		r: radius,
	}
}

func (this *Solution) RandPoint() []float64 {
	rand.Seed(time.Now().UnixNano())
	//随机角度 0,2pi //随机长度
	f := rand.Float64() * 2 * math.Pi
	//l := rand.Float64() * this.r
	l := this.r * float64(rand.Intn(math.MaxInt32)) / (math.MaxInt32 - 1)
	//f := 2 * math.Pi * float64(rand.Intn(math.MaxInt32))/(math.MaxInt32 - 1)

	res := make([]float64, 2)
	res[0] = l*math.Cos(f) + this.x
	res[1] = l*math.Sin(f) + this.y
	return res
}
