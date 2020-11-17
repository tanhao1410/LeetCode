/**
 * @author tanhao
 * @date 2020/11/17 09:55
 */
public class Solution {

    //剑指 Offer 37. 序列化二叉树
    public String serialize(TreeNode root) {
        StringBuilder res = new StringBuilder("[");
        //层次遍历
        List<TreeNode> queue = new ArrayList<TreeNode>();
        queue.add(root);
        while (!queue.isEmpty()) {
            //出队
            TreeNode first = queue.remove(0);
            if (first != null) {
                res.append(first.val).append(",");
                queue.add(first.left);
                queue.add(first.right);
            } else {
                res.append("null").append(",");
            }
        }

        //最后多余的null需要处理下
        int endIndex = res.length() - 1;
        while ("null,".contains(res.charAt(endIndex) + "")) {
            endIndex -= 1;
        }
        if (endIndex == 0) {
            endIndex = res.length() - 2;
        }

        return res.substring(0, endIndex + 1) + "]";
    }

    public TreeNode createNode(String value) {
        if ("null".equals(value)) {
            return null;
        }
        return new TreeNode(Integer.parseInt(value));
    }

    // Decodes your encoded data to tree.[1,2,3,null,null,4,5]
    public TreeNode deserialize(String data) {
        data = data.substring(1, data.length() - 1);
        String[] values = data.split(",");
        TreeNode root = createNode(values[0]);
        List<TreeNode> queue = new ArrayList<TreeNode>();

        if (root != null) {
            queue.add(root);
        }
        int i = 1;
        while (!queue.isEmpty()) {
            //从队列中取出这一排的所有的
            int queueSize = queue.size();
            //依次取出下一层的二倍数量的字符串
            for (int j = 0; j < queueSize; j++) {
                TreeNode head = queue.remove(0);
                if (i < values.length) {
                    TreeNode left = createNode(values[i++]);
                    head.left = left;
                    if (left != null) {
                        queue.add(left);
                    }
                }
                if (i < values.length) {
                    TreeNode right = createNode(values[i++]);
                    head.right = right;
                    if (right != null) {
                        queue.add(right);
                    }
                }
            }
        }
        return root;
    }

    public static class Result{
        public int code;
        public TreeNode res;
    }

    //剑指 Offer 68 - I. 二叉搜索树的最近公共祖先
    public Result visit(TreeNode root,int v0,int v1){
        Result res = new Result();
        res.code = -1;
        if(root != null){
            Result l = this.visit(root.left,v0,v1);
            Result r = this.visit(root.right,v0,v1);
            if (l.code == 2){
                res.code = 2;
                res.res = l.res;
            }else if (r.code == 2){
                res.code = 2;
                res.res = r.res;
            }else if ((root.val == v0 && (l.code == 1 || r.code ==1 ))||(l.code == 1 && r.code == 0) ||(l.code == 0 && r.code == 1) ||(root.val == v1 && (l.code == 0 || r.code ==0)) ){
                res.code = 2;
                res.res = root;
            }else if (root.val == v0 || l.code == 0 || r.code == 0){
                res.code = 0;
            }else if(root.val == v1 || l.code == 1|| r.code == 1){
                res.code = 1;
            }else{
                res.code = -1;
            }
        }
        return res;
    }

    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        return this.visit(root,p.val,q.val).res;
    }

    //剑指 Offer 31. 栈的压入、弹出序列
    public boolean validateStackSequences(int[] pushed, int[] popped) {
        Stack<Integer> stack = new Stack();
        int i = 0,j = 0;
        while (i < popped.length && (j < pushed.length || stack.size() > 0)){
            if (stack.size() == 0){
                stack.push(pushed[j++]);
            }
            while (stack.peek() != popped[i] && j < popped.length){
                stack.push(pushed[j++]);
            }
            if (stack.peek() == popped[i]){
               stack.pop();
               i ++;
            }else{
                return false;
            }
        }
        return i == popped.length;
    }

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