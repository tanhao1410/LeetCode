package main

//环路检测
//给定一个链表，如果它是有环链表，实现一个算法返回环路的开头节点。
//有环链表的定义：在链表中某个节点的next元素指向在它前面出现过的节点，则表明该链表存在环路。
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	if fast == nil {
		return nil
	}

	for fast != nil {

		if fast.Next == nil {
			return nil
		}

		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			//说明走到一块了
			break
		}
	}

	if fast == nil {
		return nil
	}

	for slow = head; slow != fast; slow = slow.Next {
		fast = fast.Next
	}

	return slow
}

func main() {
	//fmt.Println(Atoi("123456")+6)
	//fmt.Printf(Itoa(0))

	var node1 *ListNode
	var node2 *ListNode
	node4 := &ListNode{-4, node2}
	node3 := &ListNode{0, node4}
	node2 = &ListNode{2, node3}
	node1 = &ListNode{3, node2}

	detectCycle(node1)

}

//杨辉三角
func generate(numRows int) [][]int {
	if numRows == 0 {

	}
	var res [][]int = [][]int{
		{1},
		{1, 1},
	}
	if numRows < 2 {
		return res[0:1]
	}

	for i := 2; i < numRows; i++ {
		var row []int = []int{1}
		for j := 1; j < i; j++ {
			row = append(row, res[i-1][j-1]+res[i-1][j])
		}
		row = append(row, 1)

		res = append(res, row)
	}

	return res
}

//历链表，将链表中所有小于x的节点抽离出来，通过头插法构造新链表，然后将该新链表连接到已被处理过的原链表左边
func partition2(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil { //没有节点或就一个节点的时候直接返回即可
		return head
	}

	point := head
	var pointPrev *ListNode = nil
	var temp *ListNode = nil

	//再创建一个
	var ultra *ListNode = nil
	var ultraHead *ListNode = nil

	var lessFirst *ListNode = nil

	for point != nil {
		temp = point.Next
		if point.Val >= x {
			//将这个节点去掉。加入到一个临时的节点。
			if pointPrev != nil {
				pointPrev.Next = point.Next
				point.Next = nil
				//加入到另一个队列中
				if ultra == nil {
					ultraHead, ultra = point, point
				} else {
					ultra.Next = point
					ultra = point
				}
			} else {
				point.Next = nil
				if ultra == nil {
					ultraHead, ultra = point, point
				} else {
					ultra.Next = point
					ultra = point //ultra需要往前走
				}
			}

		} else {
			if lessFirst == nil {
				lessFirst = point
			}
			pointPrev = point //若删了该节点，则prev是无变化的
		}

		point = temp

	}

	if pointPrev == nil {
		return ultraHead
	}
	pointPrev.Next = ultraHead

	return lessFirst //此处用head有问题！需要找到小于x的第一个节点。

}

//编写程序以 x 为基准分割链表，使得所有小于 x 的节点排在大于或等于 x 的节点之前。
//如果链表中包含 x，x 只需出现在小于 x 的元素之后(如下所示)。分割元素 x 只需处于“右半部分”即可，其不需要被置于左右两部分之间。
func partition(head *ListNode, x int) *ListNode {
	//先想思路：一次遍历，第一个指针先走，遇到比x小的元素，直接跳过，遇到比x大的或相等的，停下来
	//再来一个指针，在上面指针停下后，往前走，直到遇到一个比X小的元素，然后交换这两个元素，如果走到尾都没遇到，说明结束啦

	if head == nil {
		return head
	}

	//先找到第一个比X大的或等于的元素
	var firstPrev, first *ListNode = nil, nil
	//fmt.Println(firstPrev)
	for first = head; first != nil && first.Val < x; first = first.Next {
		firstPrev = first
	}

	if first == nil {
		return head
	}
	var temp *ListNode = nil
	for second := first.Next; second != nil; {
		temp = second.Next
		if second.Val < x {
			//遇到了比x小的了，可以插入前面了
			if firstPrev != nil {
				firstPrev.Next = second
				second.Next = first
				firstPrev = second
			} else {
				head = second
				second.Next = first
				firstPrev = second
			}
		}
		second = temp
	}
	return head
}

func Itoa(num int) string {
	res := []byte{}
	if num == 0 {
		return "0"
	}

	for num != 0 {
		v := num % 10
		s := byte(v + '0')
		res = append(res, s)
		num = num / 10
	}

	return string(res)
}

func Atoi(s string) int {

	res := 0
	//先把string转成[]byte
	var bytes []byte = []byte(s)
	for i := 0; i < len(bytes); i++ {
		v := bytes[i] - '0'
		res *= 10
		res += int(v)
	}

	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。
//算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)
func oddEvenList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	//两个指针，一个每次走一步，一个每次走两步
	//一个指针永远指向1节点

	firstPrev, first := head, head.Next
	if first == nil {
		return head
	}

	second, secondPrev := first.Next, first
	for second != nil {
		//将second插入到first的前面。
		firstPrev.Next = second
		secondPrev.Next = second.Next
		second.Next = first

		//交换插入过后，first指针未有变化，
		firstPrev = second
		//first的前一个变为新插入的节点

		//只需要往前走一步即可到达应该到的位置，因为挪走了一个
		secondPrev = secondPrev.Next
		if secondPrev == nil {
			break
		}
		second = secondPrev.Next

	}

	return head
}
