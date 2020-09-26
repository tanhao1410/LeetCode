package main

func main() {

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

	if root.Val < p.Val && root.Val < q.Val{
		return lowestCommonAncestor(root.Right,p,q)
	}

	if root.Val >p.Val && root.Val > q.Val{
		return lowestCommonAncestor(root.Left,p,q)
	}

	return root
}
