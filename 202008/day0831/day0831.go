package main

func main() {

}

//表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格
func calculate(s string) int {
	//先算 *  /,

	return 0
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//导数低第k个结点
func getKthFromEnd(head *ListNode, k int) *ListNode {

	if head == nil {
		return head
	}
	//first先走k-1步
	first := head
	for i := 1; i < k; i++ {
		first = first.Next
		if first == nil {
			return nil
		}
	}
	//第二个指针再开始走
	for ; first.Next != nil; first = first.Next {
		head = head.Next
	}
	return head
}

//判断它是否可以由它的一个子串重复多次构成
func repeatedSubstringPattern(s string) bool {
	//思路：重复多次的话，一个，二个，一直 到math.sqr(len)个
	//2.找到一个相同的

	return false
}

//拼出最多的串
func findMaxForm(strs []string, m int, n int) int {
	//思路：
	//dp[m][n] = dp[m-1]

	return 0
}
