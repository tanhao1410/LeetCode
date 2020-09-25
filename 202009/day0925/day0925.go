package main

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
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
		inner := make([][]int,n-2)
		for i :=0;i < n-2;i ++{
			inner[i] = (matrix[i+1])[1:n-1]
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
