package main

func main() {

}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//复杂链表的复制
func copyRandomList(head *Node) *Node {

	if head == nil {
		return nil
	}
	m := make(map[*Node]*Node)
	var res *Node = &Node{head.Val, nil, nil}
	m[head] = res
	//先复制，再填充
	for p, resP := head.Next, res; p != nil; p = p.Next {
		node := &Node{p.Val, nil, nil}
		m[p] = node
		resP.Next = node
		resP = node
	}

	//填充random？用一个map作对应关系
	for p, resP := head, res; p != nil; {
		if p.Random != nil {
			resP.Random = m[p.Random]
		}
		p, resP = p.Next, resP.Next
	}

	return res
}

//type Node struct {
//	Val   int
//	Prev  *Node
//	Next  *Node
//	Child *Node
//}
//
////扁平化多级双向链表
//func flatten(root *Node) *Node {
//	//1.沿着child一直往后走，遇到child就把它接到next上，原来的next头用一个数组来保存。
//	//2.空间改进：把child与next互换，走到最后之后，往前走，看谁有child，有的话，就接到尾部。
//	if root == nil {
//		return root
//	}
//
//	var tail *Node = nil // 用于指向first的前一个节点,最后会指向尾
//	for first := root; first != nil; {
//		//如果有子节点，那么就将子节点和next互换
//		if first.Child != nil {
//			first.Next, first.Child = first.Child, first.Next
//			first.Next.Prev = first
//		}
//		first, tail = first.Next, first
//	}
//
//	//问题？新的尾部中，里面结构也有尾部？？用递归的方式
//
//	//此时first指向了尾部
//	//往前走，谁有child，就讲child进行递归，然后返回的结果接到尾部
//	for second := tail; second != nil; second = second.Prev {
//		if second.Child != nil {
//			tail.Next, second.Child = flatten(second.Child), nil
//			tail.Next.Prev = tail
//			for first := tail.Next; first != nil; {
//				tail, first = first, first.Next
//			}
//		}
//	}
//	return root
//}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//左右子树调换
func invertTree(root *TreeNode) *TreeNode {
	//思路：值调换
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
