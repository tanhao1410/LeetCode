class Solution {
    //199. 二叉树的右视图
    public List<Integer> rightSideView(TreeNode root) {
        //层次遍历的最后一个。
        List<Integer> res = new ArrayList();
        LinkedList<TreeNode> queue = new LinkedList();
        if(root != null) queue.add(root);
        while(queue.size() > 0){
            int size = queue.size();
            for (int i = 0;i < size;i ++){
                TreeNode first = queue.remove();
                if (first.left != null) queue.add(first.left);
                if (first.right != null) queue.add(first.right);
                if(i == size - 1) res.add(first.val);
            }
        }
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
    //1886. 判断矩阵经轮转后是否一致
    public boolean findRotation(int[][] mat, int[][] target) {
        int n = mat.length;
        //是否是90度情况
        boolean b = true;
        for(int i = 0;i < n;i ++){
            for(int j = 0;j < n;j ++){
                if(mat[i][j] != target[j][n - 1 - i]) {
                    b = false;
                    break;
                }
            }
        }
        if (b) return true;
        //是否180度
        for(int i = 0;i < n;i ++){
            for(int j = 0;j < n;j ++){
                if(mat[i][j] != target[n - 1- i][n - 1 - j]){
                    b = false;
                    break;
                }else{
                    b = true;
                }
            }
            if(!b) break;
        }
        if(b) return true;
        for(int i = 0;i < n;i ++){
            for(int j = 0;j < n;j ++){
                if (target[i][j] != mat[j][n - 1 - i]){
                    b = false;
                    break;
                }else{
                    b = true;
                }
            }
            if(!b) break;
        }
        if(b) return true;
        //是否完全相等情况
        for(int i = 0;i < n;i ++){
            for(int j = 0;j < n ;j ++){
                if(mat[i][j] != target[i][j]) return false;
            }
        }
        return true;
        //0,0   0,0
        //0,1   1,0
    }
}