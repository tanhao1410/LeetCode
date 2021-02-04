package main

import "fmt"

func main() {
	for i := 1; i < 1000; i++ {
		fmt.Print(i, ",")
	}
	fmt.Println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}))
}

//819. 最常见的单词
func mostCommonWord(paragraph string, banned []string) string {

	paragraphs := []byte(paragraph)
	//大写变小写
	for i := 0; i < len(paragraphs); i++ {
		if paragraphs[i] <= 'Z' && paragraphs[i] >= 'A' {
			paragraphs[i] += ('a' - 'A')
		}
	}
	//分割单词
	words := []string{}
	start := -1
	for i := 0; i < len(paragraphs); i++ {
		//是字母
		if paragraphs[i] >= 'a' && paragraphs[i] <= 'z' {
			//记录一个新单词的开始
			if start == -1 {
				start = i
			}
		} else {
			//记录一个单词的结束
			if i > start && start != -1 {
				words = append(words, string(paragraphs[start:i]))
			}
			start = -1
		}
	}
	//最后一个单词
	if start != -1 {
		words = append(words, string(paragraphs[start:]))
	}

	m := make(map[string]int)
	//map处理
	for _, word := range words {
		if _, ok := m[word]; ok {
			m[word]++
		} else {
			m[word] = 1
		}
	}

	//取出ban
	for _, word := range banned {
		if _, ok := m[word]; ok {
			delete(m, word)
		}
	}

	max := 0
	res := ""
	for word, count := range m {
		if count > max {
			res = word
			max = count
		}
	}

	return res
}

//923. 三数之和的多种可能
func threeSumMulti(arr []int, target int) int {

	var twoSumMulti func(arr []int, target int) int
	twoSumMulti = func(arr []int, target int) int {
		m := make(map[int]int)
		for _, v := range arr {
			if _, ok := m[v]; ok {
				m[v]++
			} else {
				m[v] = 1
			}
		}

		res := 0

		already := make(map[int]bool)

		for k, count := range m {
			if !already[k] {
				if target-k != k {
					if oCount, ok := m[target-k]; ok {
						res += (oCount * count)
						already[target-k] = true
					}
				} else {
					//两个数相等的话
					res += (count * (count - 1) / 2)
				}
			}

		}
		return res
	}

	res := 0
	for i := 0; i < len(arr)-2; i++ {
		res += twoSumMulti(arr[i+1:], target-arr[i])
		res %= 1000000007
	}

	return res
}

//1396. 设计地铁系统
type UndergroundSystem struct {
	m                map[string][]float64 //站名+站名-->[总时间，次数]
	customsStartTime map[int]int          //顾客id-->上车时间
	customsStation   map[int]string       // 顾客id-->上车的站
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		m:                make(map[string][]float64),
		customsStartTime: map[int]int{},
		customsStation:   map[int]string{},
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	//顾客进站
	this.customsStation[id] = stationName
	this.customsStartTime[id] = t
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {

	//顾客的进站点
	startStation := this.customsStation[id]
	delete(this.customsStation, id)
	//计算本次时间
	time := t - this.customsStartTime[id]
	delete(this.customsStartTime, id)

	if pre, ok := this.m[startStation+stationName]; ok {
		res := []float64{pre[0] + float64(time), pre[1] + 1}
		this.m[startStation+stationName] = res
	} else {
		res := []float64{float64(time), 1}
		this.m[startStation+stationName] = res
	}

}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	if res, ok := this.m[startStation+endStation]; ok {
		return res[0] / res[1]
	}
	return 0.0
}

//873. 最长的斐波那契子序列的长度-时间超时。。。
func lenLongestFibSubseq2(arr []int) int {
	//
	res := 0
	for i := 0; i < len(arr)-2; i++ {

		//后面的数不用考虑了
		if len(arr)-i < res {
			break
		}
		for j := i + 1; j < len(arr)-1; j++ {

			size := 2
			fir, sec := arr[i], arr[j]
			for k := i + 2; k < len(arr); k++ {
				if arr[k] > fir+sec {
					break
				}
				if arr[k] == fir+sec {
					size++
					fir, sec = sec, arr[k]
				}
			}

			if size > 2 && size > res {
				res = size
			}
		}
	}

	return res
}

//每日一题：643. 子数组最大平均数 I
func findMaxAverage(nums []int, k int) float64 {

	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	res := sum

	for i := k; i < len(nums); i++ {

		sum = sum - nums[i-k] + nums[i]
		if sum > res {
			res = sum
		}

	}

	return float64(res) / float64(k)
}
