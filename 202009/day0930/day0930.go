package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println(countAndSay(5))
	fmt.Println(permute([]int{1, 2, 3}))
}

//rust实现全排列
//pub fn permute(nums: Vec<i32>) -> Vec<Vec<i32>> {
//	let mut res:Vec<Vec<i32>> = vec![];
//
//	fn next_num(already: Vec<i32>, remain: Vec<i32>, res: &mut Vec<Vec<i32>>) {
//		if remain.is_empty() {
//			res.push(already);
//			return;
//		}
//
//		for i in remain.iter() {
//			let mut newAlready = vec![];
//			let mut newRemain = vec![];
//
//			for j in already.iter() {
//				newAlready.push(*j)
//			}
//			newAlready.push(*i);
//
//			for j in remain.iter(){
//				if *j == *i{
//					continue;
//				}
//				newRemain.push(*j);
//			}
//			next_num(newAlready, newRemain, res);
//		}
//	}
//
//	next_num(vec![], nums,&mut res);
//	return res;
//}

//全排列，nums 中无重复
func permute(nums []int) [][]int {
	res := [][]int{}

	var nextNum func(already, remain []int)
	//递归，每一轮插入一个数
	nextNum = func(already []int, remain []int) {
		if len(remain) == 0 {
			res = append(res, already)
			return
		}
		for i := 0; i < len(remain); i++ {
			newNums := make([]int, len(already))
			copy(newNums, already)
			newNums = append(newNums, remain[i])

			newRemain := []int{}
			for j := 0; j < len(remain); j++ {
				if j == i {
					continue
				}
				newRemain = append(newRemain, remain[j])
			}
			nextNum(newNums, newRemain)
		}
	}
	nextNum([]int{}, nums)
	return res
}

//外观数列
func countAndSay(n int) string {
	m := [][]int8{{1}, {1, 1}}
	for i := 2; i < n; i++ {
		item, pre := []int8{}, m[i-1]
		//需要看它的前一项的情况
		var count int8 = 1
		var num int8 = pre[0]
		for j := 1; j < len(pre)-1; j++ {
			if pre[j] == num {
				count++
			} else {
				item = append(item, count, num)
				num, count = pre[j], 1
			}
		}
		item = append(item, count, num)
		m = append(m, item)
	}

	res := ""
	for i := 0; i < len(m[n-1]); i++ {
		res += strconv.Itoa(int(m[n-1][i]))
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
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	if root.Val > val {

		//比根小，说明应该插入到左边，或往下继续传递
		if root.Left == nil {
			newNode := &TreeNode{val, nil, nil}
			root.Left = newNode
		} else {
			insertIntoBST(root.Left, val)
		}
	} else {
		if root.Right == nil {
			newNode := &TreeNode{val, nil, nil}
			root.Right = newNode
		} else {
			insertIntoBST(root.Right, val)
		}
	}

	return root
}
