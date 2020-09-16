package main

import (
	"fmt"
)

//任务，自己写一个简单的链表，并翻转

type Node struct {
	Val  int
	Next *Node
}

type BiTree struct {
	Val   int
	Left  *BiTree
	Right *BiTree
}

//递归法遍历
func PrintBitree(root *BiTree) {
	if root != nil {
		fmt.Println(root.Val)
		PrintBitree(root.Left)
		PrintBitree(root.Right)
	}
}

func main() {
	//创建一个链表
	//head := &Node{1, nil}
	//head := &Node{2,n1}
	//n3 := &Node{3,n2}
	//head := &Node{4,n3}
	//var head *Node= nil
	//
	//head.PrintNode()
	//
	//head = SwitchList(head)
	//
	//head.PrintNode()
	a, b := 1, 2
	a, b = b, a
	fmt.Println(a, b)

}

//好数对的数目
//满足 nums[i] == nums[j] 且 i < j ，就可以认为这是一组 好数对
func numIdenticalPairs(nums []int) int {
	res := 0
	//想到了map，每一次判断后，将该数放入到map中，
	//问题？相同的数的话，只算了一次。将bool改成int，用来记录该数出现的次数，下一次结果不是直接++，而是加上这个数，并让该数增加1
	var m map[int]int = make(map[int]int, 10)
	for i := 0; i < len(nums); i++ {
		if v, ok := m[nums[i]]; ok {
			res += v
			m[nums[i]] = v + 1
		}
		m[nums[i]] = 1
	}

	return res
}

func shuffle(nums []int, n int) []int {

	//使用原来的切片
	//互换，怎么个互换法呢，2 4 6 ---- n+1 n+2 n+3....
	//但由于数组是从零开始算起的，所以上面变成了 1 3 5 ---- n n+1 n+2....n+n-1
	//temp:= 0
	//for i := 1; i < n; i++ {
	//	temp = nums[i*2-1]
	//	nums[i*2-1] = nums[n+i-1]
	//	nums[n+i-1] = temp
	//}

	//理解错了，不仅仅是互换，更像是插入
	//[x1,x2,x3][y1,y2,y3]--->[x1,y1,x2,y2,x3,y3]
	//练习数组的使用，暴力求法
	temp := 0
	for i := n; i < 2*n-1; i++ {
		temp = nums[2*(i-n)+1]
		nums[2*(i-n)+1] = nums[i]
		//所有数开始往后退
		for j := i; j > 2*(i-n)+1; j-- {
			nums[j] = nums[j-1]
		}
		nums[2*(i-n+1)] = temp
	}

	return nums
}

func runningSum2(nums []int) []int {

	if nums == nil {
		return nil
	}
	if len(nums) == 0 {
		return make([]int, 1)
	}
	var res []int = nil
	res = append(res, nums[0])
	for i := 1; i < len(nums); i++ {
		sum := nums[i] + res[i-1]
		res = append(res, sum)
	}
	return res
}

func runningSum(nums []int) []int {

	var arr1 [10]int = [10]int{1, 2, 3, 4}

	fmt.Println(arr1)

	//多维数组的写法
	var arr2 [3][4]int = [3][4]int{{1, 2, 4}, {1, 2}}

	fmt.Println(arr2)
	//需要注意的是数组是值类型

	var slice1 []int = arr1[:3] //如何实现从数组转变成slice

	//var slice2 [][]int = arr2[:2]
	//golang的多维数组切片？
	var slice3 [][]int = [][]int{{0}, {0}}

	fmt.Println(slice3)
	fmt.Println(slice1)

	//if nums == nil {
	//	return nil
	//}
	//if len(nums) == 0{
	//	return make([]int,1)
	//}
	//var res []int = make([]int, len(nums))
	//为了空间，可以利用原来的空间来存放
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}

	return nums
}

//判断链表是否有换
func HaveCycle(head *Node) bool {
	//两个指针，一个往前

	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。
func kthToLast(head *ListNode, k int) int {
	//倒数第几个节点？双指针，一个先走k,
	first, second, res := head, head, head
	for ; first != nil && k > 1; k-- { //比如返回最后一个，则first与second应该走的步数相同
		first = first.Next
	}

	for ; first != nil && second != nil; first = first.Next {
		res, first, second = second, first.Next, second.Next
	}

	return res.Val
}

//反转链表2
func SwitchList2(head *Node) *Node {
	if head == nil {
		return nil
	}
	first, second := head, head.Next
	var temp *Node
	for first != nil && second != nil {
		temp = second.Next
		second.Next = first
		first.Next = temp
		first = first.Next
		second = first.Next
	}
	if head.Next != nil {
		return head.Next
	}
	return head
}

//反转链表
func SwitchList(head *Node) *Node {
	if head == nil {
		return nil
	}

	//原来头结点的next需要置位nil,否则无限循环了。
	var tail *Node = head
	next := head.Next
	var temp *Node
	for next != nil {
		//head 与next 指针互换
		temp = next.Next
		next.Next = head

		//head 和next 指针都往前走
		head = next
		next = temp
	}
	tail.Next = nil
	return head
}

//打印链表
func (head *Node) PrintNode() {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
