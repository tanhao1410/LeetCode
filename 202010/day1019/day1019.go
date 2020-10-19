package main

import "fmt"

func main() {
	fmt.Println(numTrees(100))
}

//96.不同的二叉搜索树
func numTrees(n int) int {
	//思路：递归的方式，只有一个数的时候，生成的树只有一种。选择以某个树为根，则剩余的子节点能形成的树，左，右，应该是相乘的关系，时间超时。
	//nums := make([]int,n)
	//for i:=0;i < n;i ++{
	//	nums[i] = i +1
	//}
	//return numsTrees(nums)

	//if n < 2 {
	//	return 1
	//}
	//
	//res := 0
	//for i := 1; i <= n; i++ {
	//	left := numTrees(i - 1)
	//	right := numTrees(n - i)
	//	res += (left * right)
	//}
	//return res

	dp:= []int{1,1}
	for i:=2;i <= n;i ++{
		res :=0
		for j:=1;j <= i;j ++{
			left := dp[j-1]
			right := dp[i-j]
			res += (left * right)
		}
		dp = append(dp, res)
	}
	return dp[n]
}

//采用dp方式
func numsTrees2(nums []int) int {
	if len(nums) < 2 {
		return 1
	}
	res := 0
	//有n个数，可以有n个循环
	for i := 0; i < len(nums); i++ {
		leftNum := numsTrees(nums[:i])
		right := numsTrees(nums[i+1:])
		res += (leftNum) * right
	}

	return res
}

//nums是已经排好序的--该方法时间超时，
func numsTrees(nums []int) int {
	if len(nums) < 2 {
		return 1
	}
	res := 0
	//有n个数，可以有n个循环
	for i := 0; i < len(nums); i++ {
		leftNum := numsTrees(nums[:i])
		right := numsTrees(nums[i+1:])
		res += (leftNum) * right
	}

	return res
}

//每日一题：844.比较含退格的字符串
func backspaceCompare(S string, T string) bool {
	//思路1：都转换成字符串，判断是否相等即可。
	getRealString := func(s string) string {
		s1 := []byte{}
		for i := 0; i < len(s); i++ {
			switch s[i] {
			case '#':
				if len(s1) > 0 {
					s1 = s1[:len(s1)-1]
				}
			default:
				s1 = append(s1, s[i])
			}
		}
		return string(s1)
	}

	return getRealString(S) == getRealString(T)
}
