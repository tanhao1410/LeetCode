package main

func main() {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//116.填充每个节点的下一个右侧节点指针
func connect(root *Node) *Node {
	//思路：左孩子-->右孩子，右孩子指向父节点右边的左孩子（没有？指向它的右）
	if root == nil || (root.Right == nil && root.Left == nil){
		return root
	}
	if root.Left != nil{
		if root.Right != nil{
			root.Left.Next = root.Right
			if root.Next != nil{
				if root.Next.Left != nil{
					root.Right.Next = root.Next.Left
				}else if root.Next.Right != nil{
					root.Right.Next = root.Next.Right
				}
			}
		}else{
			if root.Next != nil{
				if root.Next.Left != nil{
					root.Left.Next = root.Next.Left
				}else if root.Next.Right != nil{
					root.Left.Next = root.Next.Right
				}
			}
		}
	}else if root.Right != nil{
		if root.Next != nil{
			if root.Next.Left != nil{
				root.Right.Next = root.Next.Left
			}else if root.Next.Right != nil{
				root.Right.Next = root.Next.Right
			}
		}
	}
	connect(root.Right)
	connect(root.Left)
	return root
}
