class Solution {

    //124. 二叉树中的最大路径和
    public int maxPathSum(TreeNode root) {
        return rootStartMaxAndMaxValue(root)[1];
    }

    // int[0] 是以root开始的最大的值，int[1] 是整个的最大值。
    private int[] rootStartMaxAndMaxValue(TreeNode root){
        //以根节点开始
        if(root == null){
            return new int[]{0,-1001};
        }
        int[] left = rootStartMaxAndMaxValue(root.left);
        int[] right = rootStartMaxAndMaxValue(root.right);
        int[] res = new int[2];
        //以root 为开始的最大值，要么连接左边，要么连接右边，要么自己就行
        res[0] = Math.max(Math.max(0,left[0]),right[0]) + root.val;
        //对于结果的最大值，要么是包括root本身，要么是root作为中间，要么是来自与子树
        res[1] = Math.max(Math.max(res[0],root.val + left[0] + right[0]),Math.max(left[1],right[1]));
        return res;
    }
    public class TreeNode {
         int val;
         TreeNode left;
         TreeNode right;
         TreeNode() {}
         TreeNode(int val) { this.val = val; }
         TreeNode(int val, TreeNode left, TreeNode right) {
             this.val = val;
             this.left = left;
             this.right = right;
         }
     }
}