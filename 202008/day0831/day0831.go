package main

import (
	"fmt"
	"strconv"
)

func main() {
	//nums := [][]int{{1},{2},{3},{}}
	//fmt.Println(canVisitAllRooms(nums))
	//fmt.Println(compressString("aabcccccaaa"))
	fmt.Println(countSegments("a  b c"))
}

//统计单词数；连续的不是空格的字符
func countSegments(s string) int {

	//走到第一个非空格处
	i := 0
	for ; i < len(s) && s[i] == ' '; i++ {
	}

	count := 0
	flag := false //上一个是否是空格,true代表是空格，以方便走过连续空格位
	for ; i < len(s); i++ {
		if s[i] == ' ' {
			if flag {

			} else {
				count++
			}
			flag = true
		} else {
			flag = false
		}
	}
	//看最后一个字符是否是空格
	if len(s) > 0 && s[len(s)-1] != ' ' {
		count++
	}
	return count
}

//压缩字符串
func compressString(s string) string {

	if len(s) == 0 {
		return s
	}
	letters := []byte{s[0]}
	nums := []int16{}
	var count int16 = 0
	for i, j := 0, 0; i < len(s); i++ {
		if s[i] != s[j] {
			//letters = append(letters, byte(count))//超出byte的范围？？
			nums = append(nums, count)
			j, count = i, 1
			letters = append(letters, s[i])
		} else {
			count++
		}
	}
	nums = append(nums, count)

	if len(letters)*2 < len(s) {
		res := ""
		for i := 0; i < len(letters); i++ {
			res += string(letters[i])
			res += strconv.Itoa(int(nums[i]))
		}
		return res
	}
	return s
}

//钥匙和房间
func canVisitAllRooms(rooms [][]int) bool {
	//思路：用一个数组，表示一个方面能否被打开

	m := make(map[int]bool) //下标-->false 可进入，但还未进入，true，已访问，不存在，说明无法进入
	for i := 0; ; {
		m[i] = true //代表访问过了
		room := rooms[i]
		for _, v := range room {
			if _, ok := m[v]; !ok { //说明在map中不存在
				m[v] = false
			}
		}
		//需要进入下一个可访问的房间，需要注意遍历过的房间不再遍历了
		canButNotCount := len(m)
		for k, v := range m {
			if !v {
				i = k
				break
			}
			canButNotCount--
			if canButNotCount == 0 {
				for i := 1; i < len(rooms); i++ {
					if _, ok := m[i]; !ok {
						return false
					}
				}
				return true
			}
		}
	}

}
