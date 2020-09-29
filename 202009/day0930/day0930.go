package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(5))
}

//外观数列
func countAndSay(n int) string {
	m := [][]int8{{1},{1,1}}
	for i := 2;i < n ;i ++{
		item,pre := []int8{}, m [i -1]
		//需要看它的前一项的情况
		var count int8 = 1
		var num int8 = pre[0]
		for j := 1;j < len(pre) - 1;j ++{
			if pre[j] == num {
				count ++
			}else{
				item = append(item, count,num)
				num,count = pre[j],1
			}
		}
		item = append(item, count,num)
		m = append(m, item)
	}

	res := ""
	for i := 0 ;i < len(m[n -1]);i ++{
		res += strconv.Itoa(int(m[n -1][i]))
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//二叉搜索树中的插入操作
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	//思路：递归法，如果比root大，往右边走，右边为空，插入到右边，否则递归
	if root == nil{
		return &TreeNode{val,nil,nil}
	}
	if root.Val > val{

		//比根小，说明应该插入到左边，或往下继续传递
		if root.Left == nil{
			newNode := &TreeNode{val,nil,nil}
			root.Left = newNode
		}else{
			insertIntoBST(root.Left,val)
		}
	}else{
		if root.Right == nil{
			newNode := &TreeNode{val,nil,nil}
			root.Right = newNode
		}else{
			insertIntoBST(root.Right,val)
		}
	}

	return root
}
