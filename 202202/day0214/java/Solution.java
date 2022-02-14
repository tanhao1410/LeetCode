class Solution {
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