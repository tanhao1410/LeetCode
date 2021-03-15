package main

import (
	"fmt"
	"sort"
)

func main() {
	strings := permutation("aabc")
	for _, s := range strings {
		fmt.Println(s)
	}
}

//面试题 08.08. 有重复字符串的排列组合
func permutation(S string) []string {
	//1.hash去重，2.字符串先进行排序。
	bytes := []byte(S)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})

	res := []string{}
	nextLetter("", bytes, &res)

	return res
}

func nextLetter(s string, letters []byte, res *[]string) {

	if len(letters) == 0 {
		*res = append(*res, s)
	}

	//从剩下的字母中选中一个
	for i := 0; i < len(letters); i++ {

		if i > 0 && letters[i] == letters[i-1] {
			continue
		}

		nextLetters := []byte{}
		//下一个字母应该要少一个
		for j := 0; j < len(letters); j++ {
			if j != i {
				nextLetters = append(nextLetters, letters[j])
			}
		}
		nextLetter(s+string(letters[i]), nextLetters, res)

	}

}

//面试题 05.01. 插入
func insertBits(N int, M int, i int, j int) int {
	//i - j 位置改为0
	num := 0
	for k := 0; k < 32; k++ {
		if k > j || k < i {
			num += 1 << k
		}
	}
	num &= N
	//求插入的数M，i-0代表M应该右移的位置数
	m := M << i
	res := int32(num | m)
	return int(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//面试题 04.12. 求和路径
func pathSum(root *TreeNode, sum int) int {
	//递归，每次返回到一个节点时，会传过来一个已有的值。如果加上val == sum，那么结果+1，传递给下一层的时候，加上自身
	res := 0
	pathSumContainParents(root, sum, 0, &res)

	//求它的左右节点的
	if root.Left != nil {
		res += pathSum(root.Left, sum)
	}
	if root.Right != nil {
		res += pathSum(root.Right, sum)
	}

	return res
}

func pathSumContainParents(root *TreeNode, sum int, already int, count *int) {
	if root == nil {
		return
	}

	newAlready := already + root.Val
	if newAlready == sum {
		*count += 1
	}
	pathSumContainParents(root.Left, sum, newAlready, count)
	pathSumContainParents(root.Right, sum, newAlready, count)

}
