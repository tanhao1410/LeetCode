package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Println(readBinaryWatch(7))
}

//401. 二进制手表
func readBinaryWatch(num int) []string {
	//上面最多亮三个灯
	//下面最多亮5个
	//一个都不亮的话，是0:00
	res := []string{}

	//记录从0到59所需要的灯的数量,最多也就5个灯
	m := make([][]int, 6)
	for i := 0; i < 60; i++ {
		count := 0
		for j := i; j > 0; j >>= 1 {
			if 1&j == 1 {
				count++
			}
		}
		m[count] = append(m[count], i)
	}

	//先求小时的可能性，然后求分钟的可能性
	for i := 0; i <= 3 && i <= num; i++ {
		//可能的小时数
		h := m[i]
		//可能的小时数
		if num-i <= 5 {
			mm := m[num-i]
			for j := 0; j < len(h); j++ {
				if h[j] > 11 {
					continue
				}
				for k := 0; k < len(mm); k++ {
					item := strconv.Itoa(h[j]) + ":"
					if mm[k] < 10 {
						item += "0"
					}
					item += strconv.Itoa(mm[k])
					res = append(res, item)
				}
			}
		}
	}
	return res
}

//492. 构造矩形
func constructRectangle(area int) []int {
	//从根号area 到1 开始 找数，直到1为止
	num := int(math.Sqrt(float64(area)))
	for i := num; ; i-- {
		if area%i == 0 {
			return []int{i, area / i}
		}
	}
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
