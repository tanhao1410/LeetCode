package main

import "fmt"

func main() {
	fmt.Println(findKthNumber(13123, 2123))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//437. 路径总和 III
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	var pathSum2 func(root *TreeNode, preNum int) int
	pathSum2 = func(root *TreeNode, preNum int) int {
		if root == nil {
			return 0
		}
		//有三种情形，以自己结束的 ,以自己开头的传给左边，以自己开头给右边
		res := 0
		if preNum+root.Val == sum {
			res += 1
		}

		res += pathSum2(root.Left, root.Val+preNum)
		res += pathSum2(root.Right, root.Val+preNum)
		return res
	}

	res := pathSum2(root, 0)

	if root.Left != nil {
		res += pathSum(root.Left, sum)
	}
	if root.Right != nil {
		res += pathSum(root.Right, sum)
	}

	return res
}

//440. 字典序的第K小数字
func findKthNumber(n int, k int) int {

	//求n是几位数
	getBitCount := func(num int) int {
		res := 0
		for ; num != 0; num /= 10 {
			res++
		}
		return res
	}

	//得到10的n次方
	getTenN := func(n int) int {
		res := 1
		for i := 0; i < n; i++ {
			res *= 10
		}
		return res
	}
	//求N是几开头的
	getBit := func(num int) int {
		for ; num > 10; num /= 10 {
		}
		return num
	}
	//先确定有没有在余数中，然后再求就简化了，如果在余数中，递归，不在余数中，求

	//求n的整部分
	bigPart := getTenN(getBitCount(n)-1) * getBit(n)
	if bigPart < k {
		return findKthNumber(n-bigPart, k-bigPart)
	} else if bigPart == k {
		return bigPart
	}

	//求几位数开头的数 即 1 + 10 + 100 + 1000
	getCount := func(num int) int {
		res := 1
		for i, ten := 0, 10; i < num; i++ {
			res += ten
			ten *= 10
		}
		return res
	}

	//以 1,2，。。开头的数的数量，都相等
	count := getCount(getBitCount(bigPart-1) - 1)

	//确定第一位
	first := k/count + 1

	return first

}
