package main

import "fmt"

func main() {
	//height := []int{65,70,56,75,60,68}
	//weight := []int{100,150,90,190,95,110}
	//fmt.Println(bestSeqAtIndex(height,weight))

	//fmt.Println(oneEditAway("mart","karma"))
	fmt.Println(reverseWords("the sky is bule"))
}

//翻转单词顺序
func reverseWords(s string) string {
	resBytes := make([]byte, 0)
	for i := len(s) - 1; i >= 0; {

		if s[i] == ' ' {
			//为空格的几个情形：1.末尾的空格，2.单词的间隔3.单词的多余空格4.开头的空格
			i--
			continue
		}
		j := i
		for ; j >= 0 && s[j] != ' '; j-- {

		}
		//走到这说明j-i之间夹了一个单词了，j指向空格或j已经<0
		resBytes = append(resBytes, s[j+1:i+1]...)

		resBytes = append(resBytes, ' ')

		i = j
	}

	if len(resBytes) == 0 {
		return ""
	}
	return string(resBytes[:len(resBytes)-1])
}

//一次编辑，插入，删除，替换
func oneEditAway(first string, second string) bool {
	if len(second)-len(first) > 1 || len(second)-len(first) < -1 {
		return false
	}
	//替换一个相同的情况,字符串长度相等
	if len(second) == len(first) {
		//从第一个开始比较，如果只有一个不同，返回true
		count := 0
		for i := 0; i < len(second); i++ {
			if second[i] != first[i] {
				count++
				if count > 1 {
					return false
				}
			}
		}
		if count > 1 {
			return false
		} else {
			return true
		}
	}

	//思路：两者相差一个，从开始相等的位置开始找
	//判断第一位是否相等，如果相等，从首位对齐，不然，短的添加一个，首位对其
	//让second为长的那个
	if len(first) > len(second) {
		first, second = second, first
	}
	//首位对其
	count := 0
	for i, j := 0, 0; i < len(first) && j < len(second); {
		if first[i] != second[j] {
			//遇到不相等的了，则需要插入一个,即j往前走，i不变，相等于在first上插入一个字母。
			count++
			j++
			if count > 1 {
				return false
			}
		} else {
			j++
			i++
		}
	}
	if count > 1 {
		return false
	}

	return true
}

//在上面的人要比下面的人矮一点且轻一点。height.length == weight.length <= 10000
func bestSeqAtIndex(height []int, weight []int) int {
	//思路：主要是选哪一个，可能有些人矮点，但是重，
	if len(height) < 2 {
		return len(height)
	}

	//体重按从高到低排序
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			if height[j] > height[i] {
				height[i], height[j] = height[j], height[i]
				weight[i], weight[j] = weight[j], weight[i]
			}
		}
	}
	//动态规划呢？dp[i]表示i放在最上面，若i=1能放在它上面，则dp[i+1]=dp[i]+1
	//若不能，从前面找到能放在它上面的，都不能，则为1，但是这样的话，体重需要排序。从高往低排。
	dp := make([]int, len(height))
	res := 1
	for i := 0; i < len(height); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			//需要找最大的dp[i]而不是前面的哪个就好了
			if height[j] > height[i] && weight[j] > weight[i] && dp[j] >= dp[i] {
				dp[i] = dp[j] + 1 //即可以放在前面的上面
			}
		}

		//找最大的结果返回
		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}
