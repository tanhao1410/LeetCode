class Solution {

    //143. 重排链表
    public void reorderList(ListNode head) {
        //走到中间，倒置链表，合成插入。
        ListNode mid = getMid(head);
        ListNode revMid = reverseListNode(mid);
        while(revMid != null){
            ListNode temp = head.next;
            head.next = revMid;
            revMid = revMid.next;
            head.next.next = temp;
            head = temp;
        }
    }

    private ListNode getMid(ListNode head){
        if (head == null || head.next == null) return null;
        //双指针走法
        ListNode fast = head.next;
        ListNode slow = head;
        while(fast != null && fast.next != null){
            fast = fast.next;
            if (fast != null){
                fast = fast.next;
            }
            slow = slow.next;
        }
        ListNode res = slow.next;
        slow.next = null;
        return res;
    }

    private ListNode reverseListNode(ListNode head){
        if (head == null || head.next == null) return head;
        ListNode pre = null;
        while (head != null){
            ListNode temp = head.next;
            head.next = pre;
            pre = head;
            head = temp;
        }
        return pre;
    }

    public class ListNode {
        int val;
        ListNode next;
        ListNode() {}
        ListNode(int val) { this.val = val; }
        ListNode(int val, ListNode next) { this.val = val; this.next = next; }
    }
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