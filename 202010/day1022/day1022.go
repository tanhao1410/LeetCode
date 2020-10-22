package main

func main() {
	n1 := Node{1, []*Node{}}
	n2 := Node{2, []*Node{}}
	n3 := Node{3, []*Node{}}
	n4 := Node{4, []*Node{}}

	n1.Neighbors = append(n1.Neighbors, &n2, &n4)
	n2.Neighbors = append(n2.Neighbors, &n1, &n3)
	n3.Neighbors = append(n3.Neighbors, &n2, &n4)
	n4.Neighbors = append(n4.Neighbors, &n1, &n3)

	cloneGraph(&n1)
}

type Node struct {
	Val       int
	Neighbors []*Node
}

//133.克隆图
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	//采用hashmap ，广度优先遍历的方式
	m := make(map[int]*Node)
	queue := []*Node{}
	have := make(map[int]bool)

	copyNode := func(node *Node) *Node {

		var newNode *Node
		if k, ok := m[node.Val]; ok {
			newNode = k
		} else {
			newNode = &Node{node.Val, []*Node{}}
			m[newNode.Val] = newNode
		}

		if node.Neighbors == nil {
			return newNode
		}
		//复制兄弟
		neighbors := node.Neighbors
		for i := 0; i < len(neighbors); i++ {
			//从m中先获取，没有就创建
			if k, ok := m[neighbors[i].Val]; ok {
				newNode.Neighbors = append(newNode.Neighbors, k)
			} else {
				newNeighbors := Node{neighbors[i].Val, []*Node{}}
				m[neighbors[i].Val] = &newNeighbors
				newNode.Neighbors = append(newNode.Neighbors, &newNeighbors)
			}
		}
		return newNode
	}

	//复制第一个节点
	newHead := copyNode(node)
	//m[newHead.Val] = newHead
	have[node.Val] = true

	if node.Neighbors == nil {
		return newHead
	}
	queue = append(queue, node.Neighbors...)

	for len(queue) > 0 {
		queueLen := len(queue)
		for i := 0; i < queueLen; i++ {
			//如果该节点访问，跳过
			if have[queue[i].Val] {
				continue
			}
			copyNode(queue[i])
			have[queue[i].Val] = true
			if queue[i].Neighbors != nil {
				queue = append(queue, queue[i].Neighbors...)
			}
		}
		queue = queue[queueLen:]
	}
	return newHead
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//107.二叉树的层次遍历 II
func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}

	queueLen := len(queue)
	for queueLen > 0 {
		item := make([]int, queueLen)
		for i := 0; i < queueLen; i++ {
			item[i] = queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[queueLen:]
		queueLen = len(queue)
		res = append(res, item)
	}
	//逆置res
	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i, j = i+1, j-1
	}

	return res
}

//每日一题：763.划分字母区间
func partitionLabels(s string) []int {
	//思路：用一个数组（26）记录每一个字母在字符串中的最后的位置
	res := []int{}
	m := make([]int, 26)
	for i := len(s) - 1; i > 0; i-- {
		if m[s[i]-'a'] == 0 {
			m[s[i]-'a'] = i
		}
	}
	//开始构造结果集
	for i := 0; i < len(s); {
		//1.先获取该字母最后一个位置
		location := m[s[i]-'a']
		//2.在此位置区间的字母，是否有超出此空间的
		for j := i; j < location; j++ {
			if m[s[j]-'a'] > location {
				location = m[s[j]-'a']
			}
		}
		//3.lcation即终点
		res = append(res, location-i+1)
		i = location + 1
	}
	return res
}
