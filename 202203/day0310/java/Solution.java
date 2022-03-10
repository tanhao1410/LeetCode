class Solution {
    //138. 复制带随机指针的链表
    public Node copyRandomList(Node head) {
        //先复制节点，再复制random；用一个hashmap记录。old-new;
        HashMap<Node,Node> map = new HashMap();
        Node newHead = null;
        Node p = null;
        if(head != null){
            newHead = new Node(head.val);
            newHead.random = head.random;
            map.put(head,newHead);
            head = head.next;
            p = newHead;
        }
        while (head != null){
            Node node = new Node(head.val);
            node.random = head.random;
            map.put(head,node);
            head = head.next;
            p.next = node;
            p = p.next;
        }
        //处理random
        p = newHead;
        while (p != null){
            p.random = map.get(p.random);
            p = p.next;
        }
        return newHead;
    }
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
        public Node random;

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