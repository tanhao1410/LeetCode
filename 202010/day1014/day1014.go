package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//fmt.Println(commonChars([]string{"cool", "lock", "cook"}))
	//fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
	fmt.Println(getPermutation(3,5))
}

//60.第k个排列
func getPermutation(n int, k int) string {
	res := ""
	k--//从0开始排序的，而k是从1开始数的
	canSelected := make([]int, n)
	for i := 0; i < n; i++ {
		canSelected[i] = i + 1
	}
	for i := 0; i < n; i++ {
		index := getStartNum(canSelected, k)
		k = k % getN(len(canSelected)-1)
		res += strconv.Itoa(canSelected[index])
		for i :=0 ;i < len(canSelected)-1;i ++{
			if i < index{
				continue
			}
			canSelected[i] = canSelected[i+1]
		}
		canSelected = canSelected[:len(canSelected)-1]
	}

	return res
}

func getStartNum(canSelected []int, k int) int {
	index := k / getN(len(canSelected)-1)
	return index
}

func getN(n int) int {
	if n <= 1 {
		return 1
	}
	return getN(n-1) * n
}

//80.删除排序数组中的重复项 II
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//最多出现2次
	//记录前面的出现的次数，当出现2次了，和当前又相同，说明，该数需要删除
	//需要双指针了
	flag := math.MaxInt32
	j, i := 1, 0
	for ; j < len(nums); {
		if nums[j] != nums[i] {
			flag = nums[i]
			nums[i+1] = nums[j]
			j++
			i++
		} else if nums[i] == nums[j] && flag != nums[j] {
			flag = nums[i]
			nums[i+1] = nums[j]
			j++
			i++
		} else {
			j++
		}
	}
	return i + 1
}

//1002.查找常用字符
func commonChars(A []string) []string {
	res := []string{}

	countByte := make([]int8, 26)
	for i := 0; i < len(A[0]); i++ {
		countByte[A[0][i]-'a'] += 1
	}

	for i := 1; i < len(A); i++ {
		tempBytes := make([]int8, 26)
		for j := 0; j < len(A[i]); j++ {
			if countByte[A[i][j]-'a'] >= 1 {
				//说明该字母之前有
				tempBytes[A[i][j]-'a'] += 1
			}
		}

		for j := 0; j < 26; j++ {
			if tempBytes[j] < countByte[j] {
				countByte[j] = tempBytes[j]
			}
		}
	}

	for i := 0; i < 26; i++ {
		for j := countByte[i]; j > 0; j-- {
			res = append(res, string(i+'a'))
		}
	}

	return res
}
