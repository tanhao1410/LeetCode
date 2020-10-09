package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println(multiply("9", "9"))
	fmt.Println(search([]int{0, 1, 2}, 3))
}

//搜索旋转排序数组，O(log n) 级别，数组中不存在重复元素
func search(nums []int, target int) int {
	//思路：先找到旋转的地方，然后再进行二分查找即可。
	if len(nums) == 0 {
		return -1
	}

	head, tail := 0, len(nums)-1
	middle := (head + tail) / 2
	//是否经过旋转
	flag := false
	for nums[head] > nums[tail] {
		flag = true
		//比较head和middle
		if nums[middle] > nums[head] {
			head = middle
		} else {
			tail = middle
		}
		middle = (head + tail) / 2
	}
	if flag {
		head += 1
	}

	//开始二分查找
	//先判断最后一个数与target的关系
	if nums[len(nums)-1] > target {
		//在head 与 结尾之间找
		head, tail = head, len(nums)-1
	} else if nums[len(nums)-1] < target {
		head, tail = 0, head
	} else {
		return len(nums) - 1
	}
	middle = (head + tail) / 2
	for head <= tail {
		if nums[middle] == target {
			return middle
		}
		if nums[middle] < target {
			head = middle + 1
		} else {
			tail = middle - 1
		}
		middle = (head + tail) / 2
	}
	return -1
}

//字符串相乘
func multiply(num1 string, num2 string) string {
	res := ""
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	//1.先得到二维数组，num2的每一位与num1相乘得到一个数组,

	resMatrix := [][]int{}
	singleMultiply := func(singleNum int, index int) []int {
		res := []int{}
		//2.补0
		for i := 0; i < index; i++ {
			res = append(res, 0)
		}
		//从低位开始乘
		//进位标识
		flag := 0
		for i := len(num1) - 1; i >= 0; i-- {
			num := int(num1[i] - '0')
			item := num*singleNum + flag
			flag = item / 10
			res = append(res, item%10)
		}
		if flag != 0 {
			res = append(res, flag)
		}
		return res
	}

	for i := len(num2) - 1; i >= 0; i-- {
		resMatrix = append(resMatrix, singleMultiply(int(num2[i]-'0'), len(num2)-i-1))
	}

	//2.从二维数组中得出最后的结果
	//进位标识
	flag := 0
	for i := 0; i < len(resMatrix[len(resMatrix)-1]); i++ {
		num := flag
		for j := len(resMatrix) - 1; j >= 0 && len(resMatrix[j]) > i; j-- {
			num += resMatrix[j][i]
		}
		flag, num = num/10, num%10
		res = strconv.Itoa(num) + res
	}

	if flag != 0 {
		res = strconv.Itoa(flag) + res
	}

	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//判断是否是环行链表
func hasCycle(head *ListNode) bool {
	//双指针
	if head == nil {
		return false
	}

	for slow, fast := head, head.Next; fast != nil; {
		if slow == fast {
			return true
		}
		if fast.Next == nil || fast.Next.Next == nil {
			return false
		}
		slow, fast = slow.Next, fast.Next.Next
	}

	return false
}
