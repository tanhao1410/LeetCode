package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{8},8))
}

//13.罗马数字转整数 "MMMXLV"
func romanToInt(s string) int {

	res := 0
	//采用hash的算法吧  I， V， X， L，C，D 和 M。
	m := map[string]int{"I":1,"IV":4,"V":5,"IX":9,"X":10,
		"XL":40,"L":50,"XC":90,"C":100,
		"CD":400,"D":500,"CM":900,"M":1000}
	//切割，
	for i :=0;i < len(s);i ++{
		if i < len(s) -1  && m[string(s[i])] < m[string(s[i + 1])]{
			//说明是单字母的
			res += m[string(s[i:i + 2])]
			i++
		}else{
			res += m[string(s[i])]
		}
	}

	return res
}

//34. 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}

	//-1代表没有，其他的说明为下标
	haveNum := func(start, end int) int {
		for start <= end {
			middle := (start + end) / 2
			if nums[middle] == target {
				return middle
			} else if nums[middle] < target {
				start = middle + 1
			} else {
				end = middle - 1
			}
		}
		return -1
	}

	//思路：先二分法查找，没找到，直接返回
	middleIndex :=  haveNum(0,len(nums)-1)
	if -1 == middleIndex {
		return res
	}

	//找到了，再在前面二分查找，没找到，说明，该次找到的为头，找到了，继续往前二分查找头。尾同样的思路
	for pre,end := middleIndex,middleIndex;;{
		pre = end
		end = haveNum(0, end)
		if end == -1 || end == pre{
			res[0] = pre
			break;
		}
	}

	for pre,end := middleIndex,middleIndex;;{
		pre = end
		end = haveNum(end + 1, len(nums) -1)
		if end == -1 || end == pre{
			res[1] = pre
			break;
		}
	}

	return res
}

//数字转罗马数字 1-3999
func intToRoman(num int) string {

	processNum := func(count int, nine, five, four, one string) string {
		res := ""
		for count > 0 {
			if count == 9 {
				res += nine
				count -= 9
			}

			if count >= 5 {
				res += five
				count -= 5
			}

			if count == 4 {
				res += four
				count -= 4
			}

			for count < 4 && count > 0 {
				res += one
				count--
			}
		}
		return res
	}

	//1.先确实是几千
	res := processNum(num/1000, "", "", "", "M")
	num %= 1000
	//2.确定是几百
	res += processNum(num/100, "CM", "D", "CD", "C")
	num %= 100
	//3.确定是几十
	res += processNum(num/10, "XC", "L", "XL", "X")
	num %= 10
	//4.确定是几
	res += processNum(num, "IX", "V", "IV", "I")

	return res
}
