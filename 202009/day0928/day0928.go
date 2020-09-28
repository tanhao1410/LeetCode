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
	if root == nil {
		return root
	}
	connect2(root.Left, root.Right, nil)
	return root
}

func connect2(left, right, brother *Node) {
	if left != nil { //左孩子不为空
		if right != nil { //右也不为空

			left.Next = right
			for nextBrother := brother; nextBrother != nil; nextBrother = nextBrother.Next {

				if nextBrother.Left != nil {
					right.Next = nextBrother.Left
					break
				}
				if nextBrother.Right != nil {
					right.Next = nextBrother.Right
					break
				}
			}

			connect2(right.Left, right.Right, right.Next)
		} else { //左不空，右空

			//brother虽然没有子节点，但是，它的下一个可能有子节点也算！
			for nextBrother := brother; nextBrother != nil; nextBrother = nextBrother.Next {

				if nextBrother.Left != nil {
					left.Next = nextBrother.Left
					break
				}
				if nextBrother.Right != nil {
					left.Next = nextBrother.Right
					break
				}
			}
		}
		//递归处理
		connect2(left.Left, left.Right, left.Next)

	} else {              //左孩子为空
		if right != nil { //右也不空
			for nextBrother := brother; nextBrother != nil; nextBrother = nextBrother.Next {

				if nextBrother.Left != nil {
					right.Next = nextBrother.Left
					break
				}
				if nextBrother.Right != nil {
					right.Next = nextBrother.Right
					break
				}
			}
			connect2(right.Left, right.Right, right.Next)
		}
	}
}
