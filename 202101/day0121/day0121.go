package main

func main() {

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
