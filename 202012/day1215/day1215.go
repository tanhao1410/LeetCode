package main

import "fmt"

func main() {
	fmt.Println(monotoneIncreasingDigits(100))
	fmt.Println(validUtf8([]int{255}))
}

//393. UTF-8 编码验证
func validUtf8(data []int) bool {
	//只看最低八位。若第一位为0，则但字节，11->双字节，而且后面的是10开头。111->三字节，1111->四字节
	countOne := func(num int) int {
		if num&0b1111_1000 == 0b1111_1000 {
			return 5
		}
		if num&0b1111_1000 == 0b1111_0000 {
			return 4
		}
		if num&0b1111_0000 == 0b1110_0000 {
			return 3
		}
		if num&0b1110_0000 == 0b1100_0000 {
			return 2
		}
		if num&0b1100_0000 == 0b1000_0000 {
			return 1
		}
		return 0
	}

	for i := 0; i < len(data); i++ {
		count := countOne(data[i])
		if count == 1 || count == 5 {
			return false
		}
		for j := 1; j <= count-1; j++ {
			if i+j >= len(data) || countOne(data[i+j]) != 1 {
				return false
			}
		}
		if count > 1 {
			i += count - 1
		}
	}

	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//1669. 合并两个链表
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {

	start, startPre := list1, list1
	//走a步
	for i := 0; i < a; i++ {
		start, startPre = start.Next, start
	}
	//开始拼接即可。
	startPre.Next = list2

	//找到要删除的结束位置
	for i := 0; i < b-a; i++ {
		start = start.Next
	}

	//list2走至结尾，开始拼接
	for ; list2.Next != nil; list2 = list2.Next {
	}
	list2.Next = start.Next
	return list1
}

//每日一题：738. 单调递增的数字
func monotoneIncreasingDigits(N int) int {

	//先将数字转化为数组形式更方便
	//低位在前，高位在后
	createNums := func(n int) []int {
		res := []int{}
		for ; n > 0; n = n / 10 {
			//低位
			res = append(res, n%10)
		}
		return res
	}

	createNum := func(nums []int) int {
		res := 0
		for i := len(nums) - 1; i >= 0; i-- {
			res = res*10 + nums[i]
		}
		return res
	}
	nums := createNums(N)
	//如果不属于，确认在哪个位置出问题了538--> 499。5538-->4999,10-->9,216
	//找到在哪个位置开始出现偏差的，即数组递增发生了
	index := len(nums) - 1
	for ; index > 0; index-- {
		if nums[index] > nums[index-1] {
			//比后一位大了
			for ; index+1 < len(nums) && nums[index] == nums[index+1]; index++ {
			}

			nums[index]--
			for i := index - 1; i >= 0; i-- {
				nums[i] = 9
			}
			return createNum(nums)
		}
	}
	return N
}
