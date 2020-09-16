package day0817

//判断是否是等差数列，可以重新排序
func canMakeArithmeticProgression(arr []int) bool {

	//不采用排序的方法下，怎么写呢？
	//想法，找出最小的两项，则，可以算出，等差数列的差，再依次判断是否存在对应的数，不存在，false
	m := make(map[int]bool, 100)
	//注意，如果有相等的数的话，则所有数必须全都相等
	s := [2]int{1000001, 1000001} //存放最小的两个数

	var flag bool = false
	for _, v := range arr {
		//存放最小的两个数
		if v < s[0] || v < s[1] {
			if s[0] < s[1] {
				s[1] = v
			} else {
				s[0] = v
			}
		}
		if _, ok := m[v]; ok {
			//说明存在相同的数了
			flag = true
		}
		m[v] = true
	}

	var x int = 0
	if s[0] > s[1] {
		x = s[0] - s[1]
	} else {
		x = s[1] - s[0]
		s[1], s[0] = s[0], s[1]
	}
	//s[0]>=s[1]的

	for i := 1; i < len(arr); i++ {
		b := m[s[1]+x*i] //
		if !b {          //如果不存在该数的话
			return false
		}
		//必须所有的数都相等
		if flag && s[1] != arr[i] {
			return false
		}
	}

	return true
}

//给定两个（单向）链表，判定它们是否相交并返回交点。
type ListNode struct {
	Val  int
	Next *ListNode
}

//时间复杂度o(n),空间复杂度O(1)，两个链表仍须保持原有的结构
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//思路：最简单的是暴力双层循环。不符合要求
	//map的话，空间复杂度不符合要求。
	//思考，如果两个链表是相交的话，则后面的全部都一样了。
	//思路：各自先走完，循环一遍，记录两个链表的长度，并判断最后一个是否相等。
	//大的那个，在第二次循环时，先走几步
	if headA == nil || headB == nil {
		return nil
	}
	aLen, bLen := 0, 0
	preA, preB := headA, headB
	for aPoint, bPoint := headA, headB; aPoint != nil || bPoint != nil; {
		if aPoint != nil {
			aLen++
			preA = aPoint
			aPoint = aPoint.Next
		}
		if bPoint != nil {
			bLen++
			preB = bPoint
			bPoint = bPoint.Next
		}
	}
	if preA != preB {
		return nil
	}

	//比较两个链表的大小
	x := aLen - bLen
	for aPoint, bPoint := headA, headB; aPoint != nil; {
		if x > 0 {
			x--
			aPoint = aPoint.Next
		} else if x < 0 {
			x++
			bPoint = bPoint.Next
		} else {
			if aPoint == bPoint {
				return aPoint
			}
			aPoint = aPoint.Next
			bPoint = bPoint.Next
		}
	}

	return nil
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//深拷貝
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	m := make(map[*Node]*Node, 500)
	//采用map的方式

	//先复制所有节点
	for point := head; point != nil; point = point.Next {
		//创建
		newNode := new(Node)
		newNode.Val = point.Val
		m[point] = newNode //加入map中
	}

	newHead, newPoint := m[head], m[head]

	//再复制指针
	for head = head; head != nil; head = head.Next {

		if head.Next != nil {
			newPoint.Next = m[head.Next]
		}
		if head.Random != nil {
			newPoint.Random = m[head.Random]
		}

		newPoint = newPoint.Next
	}

	return newHead
}
