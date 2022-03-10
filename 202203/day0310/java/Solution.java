class Solution {
    //589. N 叉树的前序遍历
    public List<Integer> preorder(Node root) {
        //思路：先中再左
        List<Integer> res = new ArrayList();
        if(root == null) return res;
        res.add(root.val);
        for(Node node : root.children){
            res.addAll(preorder(node));
        }
        return res;
    }
    class Node {
        public int val;
        public List<Node> children;

        public Node() {}

        public Node(int _val) {
            val = _val;
        }

        public Node(int _val, List<Node> _children) {
            val = _val;
            children = _children;
        }
    };
}