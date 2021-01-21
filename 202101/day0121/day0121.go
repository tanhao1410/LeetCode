package main

func main() {

}

//623. 在二叉树中增加一行
func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d == 1 {
		return &TreeNode{
			Val:   v,
			Left:  root,
			Right: nil,
		}
	}
	//思路：先找到d-1层，然后将他们的
	for queue, queueLen, deep := []*TreeNode{root}, 1, 1; ; queueLen = len(queue) {
		//找到了它的上一层
		if deep == d-1 {
			for i := 0; i < queueLen; i++ {
				//创建两个
				left := &TreeNode{
					Val:   v,
					Left:  queue[i].Left,
					Right: nil,
				}
				right := &TreeNode{
					Val:   v,
					Left:  nil,
					Right: queue[i].Right,
				}
				queue[i].Left = left
				queue[i].Right = right
			}
			return root
		} else {
			deep++
			for i := 0; i < queueLen; i++ {
				if queue[i].Left != nil {
					queue = append(queue, queue[i].Left)
				}
				if queue[i].Right != nil {
					queue = append(queue, queue[i].Right)
				}
			}
			queue = queue[queueLen:]
		}
	}
}

//513. 找树左下角的值
func findBottomLeftValue(root *TreeNode) int {
	//思路：层次遍历最后一层的第一个数字即可。
	queue := []*TreeNode{root}
	res := 0
	for queueLen := len(queue); queueLen > 0; queueLen = len(queue) {
		//这一行的第一个数字
		res = queue[0].Val
		for i := 0; i < queueLen; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[queueLen:]
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
