package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(smallestStringWithSwaps("dcab", [][]int{{0, 3}, {1, 2}, {0, 2}}))
}

//每日一题：1202. 交换字符串中的元素
func smallestStringWithSwaps(s string, pairs [][]int) string {
	res := []byte(s)
	//思路：寻找可以互相交换的位置集合。集合内先最小，然后组合起来即可，类似于图了
	c := make([]map[int]bool, 0)

	for _, pair := range pairs {
		noOne := true
		firstC := -1
		for i := 0; i < len(c); i++ {

			//如果又在另一个集合中了，说明两者可以合并
			if firstC != -1 && (c[i][pair[0]] || c[i][pair[1]]) {
				//该位置可以合并了。
				for k, _ := range c[i] {
					c[firstC][k] = true
					delete(c[i], k)
				}
				break
			}

			//有一个在集合中，都加入集合
			if c[i][pair[0]] || c[i][pair[1]] {
				c[i][pair[0]], c[i][pair[1]] = true, true
				firstC = i
				noOne = false
			}

		}

		if noOne {
			m := make(map[int]bool)
			m[pair[0]], m[pair[1]] = true, true
			c = append(c, m)
		}
	}
	//最后形成的集合中存有重复的，需要去重

	//c的数量即最后的可以互换的数量。集合内先统计各字母的数量，然后拼接形成最小的。
	for i := 0; i < len(c); i++ {
		byteN := make([]int, 26)
		sortIndex := []int{}
		for k, _ := range c[i] {
			byteN[s[k]-'a']++
			sortIndex = append(sortIndex, k)
		}

		temp := make([]byte, len(sortIndex))
		t := 0
		for j := 0; j < 26; j++ {
			for k := 0; k < byteN[j]; k++ {
				temp[k+t] = byte(j + 'a')
			}
			t += byteN[j]
		}

		//排序过后的下标集合。
		sort.Ints(sortIndex)
		for j := 0; j < len(sortIndex); j++ {
			res[sortIndex[j]] = temp[j]
		}

	}

	return string(res)
}
