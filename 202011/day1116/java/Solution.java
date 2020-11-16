/**
 * @author tanhao
 * @date 2020/11/16 14:13
 */


class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode(int x) {
        val = x;
    }
}

class Solution {

    //112. 路径总和
    public boolean hasPathSum(TreeNode root, int sum) {
        boolean res = false;
        if (root == null){
            return false;
        }else{
            //是否是叶子节点
            if(root.left == null && root.right == null){
                res = sum == root.val;
            }else{
                if (root.left != null){
                    res = hasPathSum(root.left,sum - root.val);
                }
                if (root.right != null){
                    res = res || hasPathSum(root.right,sum - root.val);
                }
            }
        }
        return res;
    }
}
