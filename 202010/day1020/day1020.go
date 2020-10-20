package main

func main() {

}

//56.合并区间
func merge(intervals [][]int) [][]int {

	flag := true
A:
	for flag {
		for i := 0; i < len(intervals)-1; i++ {
			interval := intervals[i]
			min, max := interval[0], interval[1]
			//和它后面的进行比较
			for j := i + 1; j < len(intervals); j++ {
				interval2 := intervals[j]
				min2, max2 := interval2[j], interval2[1]
				if min2 > max || min > max2 {
					//没重合
				} else {
					//重合了
					newInterval := []int{}
					if min < min2 {
						newInterval = append(newInterval, min)
					} else {
						newInterval = append(newInterval, min2)
					}
					if max > max2 {
						newInterval = append(newInterval, max)
					} else {
						newInterval = append(newInterval, max2)
					}

					newIntervals := make([][]int, len(intervals)-1)
					for k, l := 0, 0; k < len(intervals); k++ {
						if k != i && k != j {
							newIntervals[l] = intervals[k]
							l++
						}
					}
					newIntervals[len(intervals)-2] = newInterval
					intervals = newIntervals
					continue A
				}
			}
		}
		flag = false
	}

	return intervals
}

//100.相同的树
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		return isSameTree(p.Right, q.Right) && isSameTree(p.Left, q.Left)
	} else if p == nil && q == nil {
		return true
	} else {
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
