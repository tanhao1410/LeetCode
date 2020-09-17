package main

import "fmt"

func main() {
	edges := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}}
	fmt.Print(findRedundantDirectedConnection(edges))
}

//冗余连接2
//[[1,2], [1,3], [2,3]]，一个节点只能有一个前驱，只能有一个根，根节点无父节点
func findRedundantDirectedConnection(edges [][]int) []int {

	if len(edges) == 3 {
		return edges[2]
	}

	//如果一个节点存在两个父节点，那么记录这两条线，其中必有一个是添加的
	line1 := []int{0, 0}
	line2 := []int{0, 0}

	var line2_index int
	//先找有两个父节点的节点
	m := make(map[int]int)
	for i := 0; i < len(edges); i++ {
		if p1, ok := m[edges[i][1]]; ok {
			//之前出现过了，现在又出现，说明它有两个父
			line1[0], line1[1] = p1, edges[i][1]
			line2[0], line2[1] = edges[i][0], edges[i][1]
			line2_index = i
		} else {
			m[edges[i][1]] = edges[i][0]
		}
	}

	if line1[0] != 0 {
		//去除哪一个呢？1.如果没成环，去除line2
		if !isCycle(edges) {
			return line2
		}
		//2.如果成环了，需要去除一个使它不能成环的
		//先去line2_看能否去掉环
		edges[line2_index][0] = 0
		if isCycle(edges) {
			return line1
		} else {
			return line2
		}
	} else {

		for i := len(edges) - 1; i > 0; i-- {
			temp := edges[i][0]
			edges[i][0] = 0
			if !isCycle(edges) {
				edges[i][0] = temp
				return edges[i]
			} else {
				edges[i][0] = temp
			}
		}
	}

	//3.两个line不存在，说明成环了，需要去除
	return nil
}

//判断有没有环
func isCycle(edges [][]int) bool {

	allNodes := make(map[int]bool) //所有的节点

	map_node_childs := make(map[int][]int)
	for i := 0; i < len(edges); i++ {

		if edges[i][0] == 0 {
			continue
		}

		if _, ok := map_node_childs[edges[i][1]]; !ok {
			map_node_childs[edges[i][1]] = []int{}
		}

		if v, ok := map_node_childs[edges[i][0]]; ok { //
			v = append(v, edges[i][1])
			map_node_childs[edges[i][0]] = v
			//fmt.Println(len(v))
		} else {
			map_node_childs[edges[i][0]] = []int{edges[i][1]}
		}
		allNodes[edges[i][1]] = true
		allNodes[edges[i][0]] = true
	}

	alreadyDeleteValue := make(map[int]bool)
	alreadyDeleteValue[0] = true

	for {
		size := len(alreadyDeleteValue)
		for k, v := range map_node_childs {
			//如果说没有子节点就删除

			haveChild := false
			for _, child := range v {
				if !alreadyDeleteValue[child] {
					haveChild = true
				}
			}

			if !haveChild {
				alreadyDeleteValue[k] = true
			}

		}
		if size == len(alreadyDeleteValue) {
			break
		}
	}
	return (len(alreadyDeleteValue) - 1) != len(allNodes)
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
