package main

func main() {

}

//链表中更大的结点
func nextLargerNodes(head *ListNode) []int {
	res := []int{}

	if head == nil {
		return res
	}
	//从后往前比较容易，将list转化成数组
	for ; head != nil; head = head.Next {
		res = append(res, head.Val)
	}

	temp := res[len(res)-1] //最后一个数的值
	res[len(res)-1] = 0
	for i := len(res) - 2; i >= 0; i-- {

		//如果temp不比res[i]大，而且res[i+1]也不比它大，就为0,
		if temp > res[i] {
			res[i], temp = temp, res[i]
		} else if res[i+1] > res[i] {
			res[i], temp = res[i+1], res[i]
		} else {
			//还是无法保证结果的。
			temp = res[i]
			flag := false
			for j := i + 1; j < len(res); j++ {
				if res[j] > res[i] {
					res[i] = res[j]
					flag = true
					break
				}
			}
			if !flag {
				res[i] = 0
			}

		}
	}

	return res
}

func isSubPath2(head *ListNode, root *TreeNode) bool {
	return isSubPathFunc2(head, head, root)
}

func isSubPathFunc2(primitive, head *ListNode, root *TreeNode) bool {
	//从root开始，找和head相等的，然后，递归调用，root向后走，head 向后走，
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}

	//每一次断了后，需要重新从头结点开始比较，而不能拿剩余的结点去比较了。
	if head.Val == root.Val {
		return isSubPathFunc2(primitive, head.Next, root.Left) ||
			isSubPathFunc2(primitive, head.Next, root.Right)
	}

	//下一个结点不相等的话，新链表从头开始比较的时候，跳过了该节点。
	if head == primitive {
		return isSubPathFunc2(primitive, primitive, root.Left) ||
			isSubPathFunc2(primitive, primitive, root.Right)
	} else {
		return isSubPathFunc2(primitive, primitive, root) ||
			isSubPathFunc2(primitive, primitive, root)
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//合并二叉树
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	//明显的递归的方式，左子树与右子树分别与t2的左右子树合并
	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	t1.Val += t2.Val
	//递归调用
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)

	return t1
}
