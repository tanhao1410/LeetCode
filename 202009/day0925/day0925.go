package main

func main() {
	//matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//rotate(matrix)
	nums := []int{1, 3, 2}
	nextPermutation(nums)
}

//k个一组反转链表
func reverseKGroup(head *ListNode, k int) *ListNode {

	if k == 1 {
		return head
	}

	//k个一组进行反转链表
	newHead, tail := head, head
	//记录目前访问了第多少个元素了
	count := 0
	for p := head; p != nil; {
		count++
		if count == k {
			//第一次反转，返回的头节点为本次反转后的头
			temp := p.Next
			p, p.Next = p.Next, nil
			newHead, tail = reverseList(head)
			//下一次进行反转的头
			head = temp
		} else if count%k == 0 {
			//第二次及之后的反转
			temp := p.Next
			p, p.Next = p.Next, nil
			h, t := reverseList(head)
			//之前的尾部接上这次反转的链表
			tail.Next = h
			tail = t
			head = temp
		} else {
			p = p.Next
		}
	}

	if head != newHead{
		//说明反转了至少一次，而head中可能还有没被反转的剩余的节点，接上
		tail.Next = head
	}
	//如果相等说明没进行任何反转，直接返回头节点即可

	return newHead
}

//反转链表，返回头和尾
func reverseList(head *ListNode) (*ListNode, *ListNode) {
	temp := head
	var pre *ListNode
	for ; head != nil; {
		head.Next, head, pre = pre, head.Next, head
	}
	return pre, temp
}

//下一个排列
func nextPermutation(nums []int) {
	//思路：从后往前找一个 前面比它小的数，找不到的话，说明，是按降序 排列的，颠倒顺序返回即可。
	//若找到了，和前面的进行交换，然后，它后面的数字按照从小到大进行排序返回即可。

	//上述想法错误！不能是直接交换，而应该是选择一个比前面大，的最小的进行交换。
	if len(nums) < 2 {
		return
	}

	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			//说明找到了这个数了,交换这两个数
			//nums[i-1], nums[i] = nums[i], nums[i-1]

			exchange := i
			for min := i + 1; min < len(nums); min-- {
				if nums[min] > nums[i-1] && nums[min] < nums[exchange] {
					exchange = min
				}
			}
			nums[i-1], nums[exchange] = nums[exchange], nums[i-1]

			sortNum(nums[i:])
			return
		}
	}

	//如果都没有
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}

func sortNum(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {

	size := len(inorder)
	if size == 0 {
		return nil
	} else if size == 1 {
		return &TreeNode{inorder[0], nil, nil}
	}

	rootValue := postorder[size-1]
	//确立左右子树的中序和后序
	i := 0
	for ; i < size && inorder[i] != rootValue; i++ {
	}

	//根据postorder确立根节点
	root := &TreeNode{rootValue, nil, nil}

	if i != 0 {
		//有左子树
		leftChildInorder := inorder[:i]
		leftChildPostorder := postorder[:i]
		root.Left = buildTree(leftChildInorder, leftChildPostorder)
	}
	if i != size-1 {
		//有右子树
		rightChildInorder := inorder[i+1:]
		rightChildPostorder := postorder[size-1-len(rightChildInorder) : size-1]
		root.Right = buildTree(rightChildInorder, rightChildPostorder)
	}
	return root
}

//旋转矩阵，不占用额外空间
func rotate(matrix [][]int) {

	n := len(matrix)

	//先旋转外圈：
	for i := 0; i < n-1; i++ {
		matrix[i][n-1], matrix[n-1][n-1-i], matrix[n-1-i][0], matrix[0][i] =
			matrix[0][i], matrix[i][n-1], matrix[n-1][n-1-i], matrix[n-1-i][0]
	}

	////如何补内圈
	//for j := 0; j < n/2 -1; j++ {
	//
	//	for i := 0; i < n-j*2 -1; i++ {
	//		matrix[i][n-1], matrix[n-1][n-1-i], matrix[n-1-i][0], matrix[0][i] =
	//			matrix[j][i+j], matrix[i+j][n-j-1], matrix[n-j-1][n-j-1-i], matrix[n-i-1-i][j]
	//	}
	//
	//}
	//递归的方式旋转内圈吧
	if n > 3 {
		inner := make([][]int, n-2)
		for i := 0; i < n-2; i++ {
			inner[i] = (matrix[i+1])[1 : n-1]
		}
		rotate(inner)
		//问题？因为在这里进行了切割，导致少了一些东西
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//分隔链表
func partition(head *ListNode, x int) *ListNode {
	//思路:一个指针往前走，停在比x大的第一个位置的前面。
	//第二个指针往后走，遇到比x小的就插入到前面指针的后面即可
	if head == nil {
		return head
	}

	var first *ListNode // 指向第一个不小于x的前一个位置
	p := head           //指向第一个不小于x的节点
	pPre := head        // 指向p的前一个节点
	for ; p.Val < x && p != nil; p = p.Next {
		pPre = p
		first = p
	}

	for p != nil {
		if p.Val >= x {
			pPre = p
			p = p.Next
		} else {
			if first == nil {
				//first == nil ,
				p, first, pPre.Next = p.Next, p, p.Next
				first.Next = head
				head = first
			} else {
				p, pPre.Next, first.Next, p.Next = p.Next, p.Next, p, first.Next
				first = first.Next
			}

		}
	}

	return head

}

func removeZeroSumSublists(head *ListNode) *ListNode {

	return nil
}
