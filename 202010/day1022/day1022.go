package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//107.二叉树的层次遍历 II
func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil{
		return res
	}
	queue := []*TreeNode{root}

	queueLen := len(queue)
	for queueLen > 0{
		item := make([]int,queueLen)
		for i := 0;i < queueLen;i ++{
			item[i] = queue[i].Val
			if queue[i].Left != nil{
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil{
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[queueLen:]
		queueLen = len(queue)
		res = append(res, item)
	}
	//逆置res
	for i,j:=0,len(res) -1;i < j;{
		res[i],res[j] = res[j],res[i]
		i,j = i + 1,j -1
	}

	return res
}

//每日一题：763.划分字母区间
func partitionLabels(s string) []int {
	//思路：用一个数组（26）记录每一个字母在字符串中的最后的位置
	res := []int{}
	m := make([]int, 26)
	for i := len(s) - 1; i > 0; i-- {
		if m[s[i]-'a'] == 0 {
			m[s[i]-'a'] = i
		}
	}
	//开始构造结果集
	for i := 0; i < len(s); {
		//1.先获取该字母最后一个位置
		location := m[s[i]-'a']
		//2.在此位置区间的字母，是否有超出此空间的
		for j := i; j < location; j++ {
			if m[s[j]-'a'] > location {
				location = m[s[j]-'a']
			}
		}
		//3.lcation即终点
		res = append(res, location-i+1)
		i = location + 1
	}
	return res
}
