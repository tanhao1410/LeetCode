package main

func main() {

}

//青蛙跳台阶的问题
func numWays(n int) int {
	//dp算法
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2] // 可以从前一个台阶或前两个台阶跳过来
	}
	return dp[n]
}

//一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数
func singleNumber(nums []int) int {
	//思路：暴力方式一趟循环用map记录：
	//思路2：如果某个数已经出现了三次，那么应该把这个位置给删了。省下了点空间
	//
	m := make(map[int]int)
	for _, v := range nums {
		if count, ok := m[v]; ok {
			if count == 2 {
				//说明达到了三次了
				delete(m, v)
			}
		} else {
			m[v] = 1
		}
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return -1
}

//不能用乘除法，不能用for循环 1+2+。。。+n ，if
func sumNums(n int) int {
	//即只能用加法，能想到的有递归，或者用动态规划

	//dp := []int{}
	//dp[0] = 1
	//还是需要循环：dp[i] = i+1 + dp[i-1]

	return sumNums(n-1) + n //用到了if 也是不行
}

//递归或循环肯定是需要判断结束条件的，所以应该都不能用的。
func sumNums2(n int) int {
	//即只能用加法，能想到的有递归，或者用动态规划

	dp := make([]int, n)
	dp[0] = 1
	//还是需要循环：dp[i] = i+1 + dp[i-1]

	return dp[n-1] //用到了if 也是不行
}

func sumNums2_dp(n int, dp []int) {
	dp[n-1] = dp[n-2] + n
}

//中序遍历二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归法
func inorderTraversal(root *TreeNode) []int {
	res := &[]int{}
	inorderTraversalDi(root, res)
	return *res
}

func inorderTraversalDi(root *TreeNode, res *[]int) {
	if root != nil {
		//先遍历左
		inorderTraversalDi(root.Left, res)
		*res = append(*res, root.Val)
		inorderTraversalDi(root.Right, res)
	}
}
