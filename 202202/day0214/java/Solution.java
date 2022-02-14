class Solution {
    //572. 另一棵树的子树
    public boolean isSubtree(TreeNode root, TreeNode subRoot) {
        //递归思路：写一个方法，判断是否是包含另一个树。则是否存在子树的条件变成了:某一个节点与subRoot的根相同，且其子树包含其左右子树
        if (root == null) return false;
        if (subRoot == null) return true;
        if (root.val == subRoot.val){
            if (isContains(root.left,subRoot.left) && isContains(root.right,subRoot.right)){
                return true;
            }
        }
        return isSubtree(root.left,subRoot) || isSubtree(root.right,subRoot);

    }

    public boolean isContains(TreeNode root,TreeNode subRoot){
        if (subRoot == null) return root == null;
        if (root == null) return false;
        if (root.val != subRoot.val) return false;
        return isContains(root.left,subRoot.left)&& isContains(root.right,subRoot.right);
    }

    //117. 填充每个节点的下一个右侧节点指针 II
    public Node connect(Node root) {
        //思路：采用层序遍历
        LinkedList<Node> queue = new LinkedList();
        if (root == null){
            return root;
        }
        queue.add(root);
        int size = queue.size();
        while (size > 0){
            Node pre = queue.remove(0);
            if (pre.left != null) queue.add(pre.left);
            if (pre.right != null) queue.add(pre.right);
            for (int i = 1;i < size;i ++){
                Node node = queue.remove(0);
                if (node.left != null) queue.add(node.left);
                if (node.right != null) queue.add(node.right);
                pre.next = node;
                pre = node;
            }
            size = queue.size();
        }
        return root;
    }
}