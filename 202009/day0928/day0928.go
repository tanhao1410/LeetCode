package main

func main() {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	connect3(root)

	//
	return root
}

func connect3(root *Node) *Node {
	//思路：根节点的左指向它的右，递归？
	if root == nil{
		return root
	}
	connect2(root.Left,root.Right)
	return root
}

func connect2(left,right *Node) {
	if left != nil{
		left.Next = right
	}
	//右边的返回结果，是它的最左边的，然后，介入到左子树的最后一个。
	connect3(right)
	connect3(left)
}
