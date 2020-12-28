package main

func main() {

}

//每日一题：188. 买卖股票的最佳时机 IV
func maxProfit(k int, prices []int) int {
	//思路：1.先划分区间，每个区间都是的最小值都是开头，最大值是结尾
	//1.1,合并部分区间，合并规则，若起点大于等于前一个的终点，应该合并
	//2.若区间的数量小于k，则取最大的几个即可。
	//3.若大于k,合并一部分区间，合并规则：若

	return 0
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//面试题 04.05. 合法二叉搜索树
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	nums := []int{}
	var dfs func(root *TreeNode)
	//中序遍历，左，中，右
	dfs = func(root *TreeNode) {
		if root != nil {
			dfs(root.Left)
			nums = append(nums, root.Val)
			dfs(root.Right)
		}
	}
	dfs(root)

	for pre, i := nums[0], 1; i < len(nums); i++ {
		if nums[i] <= pre {
			return false
		}
		pre = nums[i]
	}

	return true
}

//面试题 04.01. 节点间通路
func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	if n == 0 {
		return false
	}
	//思路：用一个集合来记录start能够到达的所有点，结束条件为该集合加入了targe或加入不了新的节点了
	m := make(map[int]map[int]bool) //val -->(target,bool)
	for _, v := range graph {
		if mm, ok := m[v[0]]; ok {
			mm[v[1]] = true
		} else {
			m[v[0]] = make(map[int]bool)
			m[v[0]][v[1]] = true
		}
	}
	//一趟遍历过后，就知道了那些节点可以直接到哪些节点
	//从start开始
	mAlready := make(map[int]bool)
	if set, ok := m[start]; ok {
		for flag := true; flag; {
			flag = false
			//只要有新的加入进来，就把false改为true
			for k, _ := range set {
				//已经加入过的不用重复再加了
				//把它所能到达的都加入进来
				if kset, kok := m[k]; !mAlready[k] && kok {
					mAlready[k] = true
					for kk, _ := range kset {
						if !set[kk] {
							set[kk] = true
							flag = true
						}
					}
				}
			}
			if set[target] {
				return true
			}
		}
	}
	return false
}
