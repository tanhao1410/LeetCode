package day0821

type ListNode struct {
	Val  int
	Next *ListNode
}

//判断链表是否是回文链表
func isPalindrome(head *ListNode) bool {

	if head == nil {
		return false
	}
	//先得到链表的长度
	length := 0
	for p := head; p != nil; p = p.Next { //不能通过head来走，因为后面还要用呢
		length++
	}
	if length < 2 {
		return true
	}

	//反转链表，仅反转到中间节点。然后再进行判断。

	var pre *ListNode
	for middle := length / 2; head != nil && middle > 0; middle-- {
		pre, head, head.Next = head, head.Next, pre //交换
	}

	//奇书个节点 >=3
	if length%2 != 0 {
		head = head.Next
	}
	//pre往前走，head往后走，直到都走完并且都相等的话，说明是回文
	for ; head != nil && pre != nil; head = head.Next {
		if head.Val != pre.Val {
			return false
		}
		pre = pre.Next
	}

	return true
}

//0 <= nums[i] <= 100
func smallerNumbersThanCurrent(nums []int) []int {
	res := make([]int, len(nums))
	//先来个暴力解决法
	for k, v := range nums {
		count := 0
		for _, v2 := range nums {
			if v2 < v {
				count++
			}
		}
		res[k] = count
	}
	return res
}

//0 <= nums[i] <= 100
func smallerNumbersThanCurrent2(nums []int) []int {
	//想想高效的思路
	//因为nums[i]=<100
	m := make([]int, 101) //记录每一个数出现的次数，下标为数，内容为次数
	//遍历一次数组
	for _, v := range nums {
		m[v] = m[v] + 1
	}
	//遍历m数组，是m数组达到这种状态，下标为数，内容为小于自己下标的数有多少个
	preNumCount := m[0] //前一个数字出现的次数
	m[0] = 0
	for i := 1; i < 101; i++ {
		preNumCount, m[i] = m[i], (m[i-1] + preNumCount)
	}

	res := make([]int, len(nums))
	for k, v := range nums {
		res[k] = m[v]
	}

	return res

}

//交换和
//分析：dis =2(x-y)可以满足
func findSwapValues(array1 []int, array2 []int) []int {

	res := make([]int, 0)

	m := make(map[int]int) //记录array1，中某数字出现了，并记录其下标

	sum1 := 0
	for _, v := range array1 {

		sum1 += v
	}

	sum2 := 0
	for k, v := range array2 {
		m[v] = k
		sum2 += v

	}

	dis := sum1 - sum2
	if dis%2 != 0 {
		return res //一定不存在这样的交换
	}
	//现在开始找，array[x] - array[y] = dis/2
	//array[y] = array[x]-dis/2
	for _, v := range array1 {
		y := v - dis/2
		if _, ok := m[y]; ok {
			res = append(res, v, y)
			return res
		}
	}
	return res

}
