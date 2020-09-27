package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(calculate("14/3-2"))
}

type ListNode struct {
	Val  int
	Next *ListNode
}
//删除指定的节点
func deleteNode(head *ListNode, val int) *ListNode {
	//找到该节点，让它的前一个节点指向它的后一个节点即可
	if head == nil{
		return head
	}
	if head.Val == val{
		return head.Next
	}

	pre := head
	for p := head;p != nil;{
		pre,p = p,p.Next
		if p.Val == val{
			pre.Next = p.Next
			return head
		}
	}
	return head
}

//计算器 + - * / 3+2*2
func calculate(s string) int {
	//思路：难点在于优先级。一个栈用来存数，先读一个数进栈，再读一个符号，如果是+-,那么，数接着 入站，再读一个 符号，
	//若是+-，则，从栈中取出两数，运算再放进去。
	//若是*/，则，接着读下一个 数，并 将运算结果放进去。

	stack := []string{}

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			j := i
			for ; j < len(s) && s[j] >= '0' && s[j] <= '9'; j++ {
			}
			num := s[i:j]
			i = j - 1

			if len(stack) > 0 && (stack[len(stack)-1] == "*" || stack[len(stack)-1] == "/") {
				//从栈中取出一个数和*/号，和当前数进行运算，再放进去
				preNum, _ := strconv.Atoi(stack[len(stack)-2])
				curNum, _ := strconv.Atoi(num)
				if stack[len(stack)-1] == "*" {
					res := preNum * curNum
					//stack = append(stack, strconv.Itoa(res))
					stack[len(stack)-2] = strconv.Itoa(res)
					stack = stack[:len(stack)-1]
				} else {
					res := preNum / curNum
					stack[len(stack)-2] = strconv.Itoa(res)
					stack = stack[:len(stack)-1]
				}
			} else {
				stack = append(stack, num)
			}
		} else if s[i] == ' ' {
			continue
		} else {
			stack = append(stack, string(s[i]))
		}

	}

	res, _ := strconv.Atoi(stack[0])
	//处理完之后的栈中，肯定都是数和+-号了
	for i := 1; i < len(stack); i++ {
		if stack[i] == "+" {
			num, _ := strconv.Atoi(stack[i+1])
			res += num
			i++
		} else if stack[i] == "-" {
			num, _ := strconv.Atoi(stack[i+1])
			res -= num
			i++
		}
	}

	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//二叉搜索树的最近公共节点
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//如果root 介于两者之间的话，那么root就是最近的公共节点了，如果小于，那么，公共节点在左/右孩子
	//if root.Val == p.Val || root.Val == q.Val{
	//	return root
	//}

	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	return root
}
