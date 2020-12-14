package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

//429. N 叉树的层序遍历
func levelOrder(root *Node) [][]int {
	res := [][]int{}

	if root == nil {
		return res
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		item := []int{}
		itemLen := len(queue)
		for i := 0; i < itemLen; i++ {
			item = append(item, queue[i].Val)
			if queue[i].Children != nil {
				queue = append(queue, queue[i].Children...)
			}
		}
		queue = queue[itemLen:]
		res = append(res, item)
	}
	return res
}
