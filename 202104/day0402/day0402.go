package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 10000; i++ {
		fmt.Print(rand.Intn(1000), ",")
	}
}

//1797. 设计一个验证系统
type AuthenticationManager struct {
	tokens     map[string]int
	timeToLive int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		tokens:     make(map[string]int),
		timeToLive: timeToLive,
	}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.tokens[tokenId] = currentTime
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	//有的情况下更新
	if time, ok := this.tokens[tokenId]; ok {
		//没过期
		if currentTime < time+this.timeToLive {
			this.tokens[tokenId] = currentTime
		}
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	res := 0
	time := currentTime - this.timeToLive
	for k, v := range this.tokens {
		if v > time {
			res++
		} else {
			delete(this.tokens, k)
		}
	}
	return res
}

//面试题 17.21. 直方图的水量
func trap(height []int) int {
	//思路：访问到第一个高度不为0，以此为基点，往后走，直到找到比它的为止，如果没有，则下一个
	//问题？如果它很大，后面第二小的放在了最后面。跳过了，就不适合了/
	//找不到比它的大的话，可以找后面第最大的，求他们之间存的水量即可。然后从下一个继续
	//如果有比它大或相等的，则它所能存的水量可以求出来。然后在一次为基础，往后走。
	res := 0
	//求存的水量
	getTrap := func(start, end int) int {
		if start >= end-1 {
			return 0
		}
		res := 0
		if height[start] > height[end] {
			res = height[end] * (end - start - 1)
		} else {
			res = height[start] * (end - start - 1)
		}
		for i := start + 1; i < end; i++ {
			res -= height[i]
		}
		return res
	}

	for i := 0; i < len(height)-1; {

		if height[i] == 0 {
			i++
			continue
		}
		//找到后面最大的或比height[i]大于等于的
		maxIndex := i + 1
		for j := i + 1; j < len(height); j++ {
			//找后面最高的
			if height[j] >= height[maxIndex] {
				maxIndex = j
			}
			//找到比它高的了
			if height[j] >= height[i] {

				//求他们之间的水量
				res += getTrap(i, j)
				i = j
				break
			}
		}
		//即后面的所有都比i处要低
		if height[maxIndex] < height[i] {
			res += getTrap(i, maxIndex)
			i = maxIndex
		}
	}
	return res
}
