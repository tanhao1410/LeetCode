package main

func main() {

}

//100.相同的树
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil{
		if p.Val != q.Val{
			return false
		}
		return isSameTree(p.Right,q.Right)&&isSameTree(p.Left,q.Left)
	}else if p == nil && q == nil{
		return true
	}else{
		return false
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//103.二叉树的锯齿形层次遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	//用队列,进行层次遍历，先进后出，奇书层 左右，偶数层，右左
	res := [][]int{}
	queue := []*TreeNode{}
	if root != nil {
		queue = append(queue, root)
	}

	flag := true
	for len(queue) != 0 {
		queueLen := len(queue)
		item := make([]int, queueLen)
		for i := 0; i < queueLen; i++ {

			if flag {
				item[i] = queue[i].Val
			} else {
				item[i] = queue[queueLen-1-i].Val
			}

			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}

		}
		flag = !flag
		queue = queue[queueLen:]
		res = append(res, item)
	}
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//每日一题：143.重排链表
func reorderList(head *ListNode) {
	//求总数
	count := 0
	for p := head; p != nil; p = p.Next {
		count++
	}
	if count <= 2 {
		return
	}
	//逆置后面的
	count = count / 2
	middle, pre := head, head
	for ; count >= 0; count-- {
		pre = middle
		middle = middle.Next
	}
	pre.Next = nil

	//逆置middle与end之间的链表
	if middle.Next != nil {

		m, n := middle, middle.Next
		for ; n != nil; {
			n.Next, m, n = m, n, n.Next
		}
		middle.Next = nil
		middle = m
		//middle为新头
		for pp := head; pp != nil && middle != nil; {
			pp.Next, middle.Next, middle, pp = middle, pp.Next, middle.Next, pp.Next
		}
	} else {
		//即middle 就一个数
		head.Next, middle.Next = middle, head.Next
	}
}
