package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//每日一题：103. 二叉树的锯齿形层序遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	flag := true
	for queueLen := len(queue); queueLen > 0; queueLen = len(queue) {
		item := make([]int, queueLen)
		for i := 0; i < queueLen; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			if flag {
				item[i] = queue[i].Val
			} else {
				item[i] = queue[queueLen-1-i].Val
			}
		}
		res = append(res, item)
		flag = !flag
		queue = queue[queueLen:]
	}
	return res
}
