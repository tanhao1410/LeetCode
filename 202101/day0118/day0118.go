package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(accountsMerge([][]string{{"David", "David0@m.co", "David1@m.co"}, {"David", "David3@m.co",
		"David4@m.co"}, {"David", "David4@m.co", "David5@m.co"}, {"David", "David2@m.co", "David3@m.co"}, {"David", "David1@m.co", "David2@m.co"}}))
}

//654. 最大二叉树
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	//先找最大值
	max := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[max] {
			max = i
		}
	}

	root := TreeNode{
		Val:   nums[max],
		Left:  constructMaximumBinaryTree(nums[:max]),
		Right: constructMaximumBinaryTree(nums[max+1:]),
	}

	return &root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//每日一题：721. 账户合并
func accountsMerge(accounts [][]string) [][]string {

	//思路：第一，name-->set(mail,mail...),遍历下一个用户，可能存在重名的情况。直接用list存，碰到相同名称的，判断邮箱是否是重叠，有，加入其中。
	//问题：前面遍历的时候虽然没有重叠，但随着后面的加入，前面的发生了重叠现象。
	//思路要改了
	l := []map[string]bool{}
	lName := []string{}

	for i := 0; i < len(accounts); i++ {
		//取得当前用户
		curAccount := accounts[i]
		isSame := -1
	OUTER:
		for j := 0; j < len(lName); j++ {
			if lName[j] == curAccount[0] {
				//找到相同名称的了，判断是否有重叠
				for k := 1; k < len(curAccount); k++ {
					if l[j][curAccount[k]] {
						isSame = j
						break OUTER
					}
				}
			}
		}

		if isSame != -1 {
			//把自己的都加入进来
			for k := 1; k < len(curAccount); k++ {
				l[isSame][curAccount[k]] = true
			}
		} else {
			//重新建立一个账户
			lName = append(lName, curAccount[0])
			emails := make(map[string]bool)
			for k := 1; k < len(curAccount); k++ {
				emails[curAccount[k]] = true
			}
			l = append(l, emails)
		}
	}

	//上面形成的集合中还可能存在重叠，需要继续合并

	dels := make([]bool, len(l))
	//只要产生了归并，就继续下去
	for noJoin := false; !noJoin; {
		noJoin = true

	OUTER2:
		for i := 0; i < len(dels); i++ {
			//该位置没被合并
			if !dels[i] {
				for j := 0; j < len(dels); j++ {
					if i != j && lName[i] == lName[j] && !dels[j] {
						//判断是否有重叠
						for k, _ := range l[j] {
							if l[i][k] {
								noJoin = false
								//合并重叠
								dels[j] = true
								for kk, _ := range l[j] {
									l[i][kk] = true
								}
								break OUTER2
							}
						}
					}
				}
			}
		}
	}

	//把结果整理成返回结果
	res := [][]string{}
	for i := 0; i < len(l); i++ {
		if !dels[i] {
			item := []string{}
			for k, _ := range l[i] {
				item = append(item, k)
			}
			sort.Strings(item)
			res = append(res, append([]string{lName[i]}, item...))
		}
	}
	return res
}
