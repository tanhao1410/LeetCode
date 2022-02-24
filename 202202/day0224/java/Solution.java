class Solution {
    //142. 环形链表 II
    public ListNode detectCycle(ListNode head) {
        if (head == null || head.next == null) return null;
        //快慢指针先找=判断是否有换，
        ListNode fast = head.next;
        ListNode slow = head;
        while(fast != null && slow != null){
            fast = fast.next;
            if (fast == null) return null;
            fast = fast.next;
            slow = slow.next;
            if (fast == slow){
                //成环了
                //此时，fast指针跳到head，一步一步走即可
                fast = head;
                slow = slow.next;
                while(fast != slow){
                    slow = slow.next;
                    fast = fast.next;
                }
                return fast;
            }
        }
        return null;
    }
    //2. 两数相加
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        //用l1返回结果
        ListNode res = l1;
        int flag = 0;
        ListNode pre = null;
        while(l1 != null && l2 != null){
            l1.val += l2.val + flag;
            flag = l1.val/10;
            l1.val %= 10;
            pre = l1;
            l1 = l1.next;
            l2 = l2.next;
        }
        if (l1 == null) {
            pre.next = l2;
            l1 = pre.next;
        }

        while(l1 != null){
            l1.val += flag;
            flag = l1.val/10;
            l1.val %= 10;
            pre = l1;
            l1 = l1.next;
        }

        if (flag > 0){
            pre.next = new ListNode(flag);
        }

        return res;
    }
    //20. 有效的括号
    public boolean isValid(String s) {
        Stack<Character> stack = new Stack();
        for(int i = 0;i < s.length();i ++){
            char cur = s.charAt(i);
            if(cur == '(' || cur =='{' || cur == '['){
                stack.push(cur);
            }else if (stack.size() > 0){
                Character pop = stack.pop();
                if (cur == ')' && pop != '(') return false;
                if (cur == ']' && pop != '[') return false;
                if (cur == '}' && pop != '{') return false;
            }else{
                return false;
            }
        }
        return stack.size() == 0;
    }
    class ListNode {
          int val;
          ListNode next;
          ListNode() {}
          ListNode(int val) { this.val = val; }
          ListNode(int val, ListNode next) { this.val = val; this.next = next; }
    }
}