/**
 * @author tanhao
 * @date 2020/11/17 09:55
 */
public class Solution {

    //剑指 Offer 26. 树的子结构
    public boolean isSubStructure(TreeNode A, TreeNode B) {
        if (A == null || B == null) {
            return false;
        }
        if (A.val == B.val && isEqual(A.left, B.left) && isEqual(A.right, B.right)) {
            return true;
        }
        return isSubStructure(A.right, B) || isSubStructure(A.left, B);
    }

    public boolean isEqual(TreeNode a, TreeNode b) {
        if (b == null) {
            return true;
        }
        if (a == null) {
            return false;
        }
        if (a.val == b.val) {
            return isEqual(a.left, b.left) && isEqual(a.right, b.right);
        }
        return false;
    }
}

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode(int x) {
        val = x;
    }
}