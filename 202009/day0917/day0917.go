package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//两两交换其中相邻的节点，并返回交换后的链表。
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换
//1->2->3->4, 你应该返回 2->1->4->3.
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	res := head.Next

	var pre *ListNode
	//还需要前一个
	for first, second := head, head.Next; second != nil; {
		first.Next, second.Next = second.Next, first
		if pre != nil {
			pre.Next = second
		}
		//指针往前走
		if first.Next != nil {
			pre = first
			first, second = first.Next, first.Next.Next
		} else {
			break
		}
	}

	return res
}

//删除中间的节点，你只能访问该节点，不是第一个，也不是最后一个
func deleteNode(node *ListNode) {
	//思路：删除的这个节点不知道它的前驱？可以是假删除，将后一个的数据给它，然后删除最后一个
	pre := node
	for next := node.Next; next != nil; {
		node.Val = next.Val
		pre = node
		node = node.Next
		next = next.Next
	}
	pre.Next = nil
}

//冗余连接2
//[[1,2], [1,3], [2,3]]，一个节点只能有一个前驱，只能有一个根，根节点无父节点
func findRedundantDirectedConnection(edges [][]int) []int {

	if len(edges) < 1 {
		return nil
	}

	//思路：第一，根据一个节点只能有一个父节点，如果有个节点含两个前驱，删除一个即可。
	//第二，如果不存在这样的，可能是因为，存在环了，即有节点指向了根节点。问题就是找根节点了
	//删除哪一个呢？任意删除一个，只要不成环即可，即删除环路上的任意一个。
	m := make(map[int]int) //记录的是它的父节点
	node_childsCount := make(map[int]int)
	for i := 0; i < len(edges); i++ {
		if _, ok := m[edges[i][1]]; ok {
			//说明这个节点在前面已经有节点指向了，这个为重复指向的
			return edges[i]
		} else {
			m[edges[i][0]] = edges[i][0]
		}
		if _, ok := node_childsCount[edges[i][0]]; ok {
			node_childsCount[edges[i][0]]++
		} else {
			node_childsCount[edges[i][0]] = 1
		}

	}

	//说明是因为环路的存在了
	//问题就是找环路了。切枝，去掉没有子节点的节点，因为他们肯定不会是环路上的，递归去，最后剩的就是环路。
	//没有子节点，即没有出现在edges[i][0]的位置。所有的节点已经保存在了m中。
	for deleteComplete := false; !deleteComplete; {
		deleteComplete = true
		for k, v := range m {
			if counts, ok := node_childsCount[k]; ok {
				if counts < 0 {
					delete(node_childsCount, k)
					node_childsCount[v]--
					delete(m, k)
					deleteComplete = false
				}
				//这个节点有子节点
			} else {
				//把m去掉时，它的父节点的孩子数应该-1
				node_childsCount[v]--
				delete(m, k)
				//只要删除了一个，就得再循环删一遍
				deleteComplete = false
			}
		}
	}
	//现在m中剩下的就是都是有子节点的了
	for i := len(edges) - 1; i >= 0; i-- {
		if _, ok := m[edges[i][0]]; ok {
			if _, ok2 := m[edges[i][0]]; ok2 {
				return edges[i]
			}
		}
	}

	return nil
}
