class Solution {
    //剑指 Offer II 028. 展平多级双向链表
    public Node flatten(Node head) {
        if (head == null) return null;
        //优先将child放进来
        if (head.child != null){
            //处理head.child
            Node next = head.next;
            Node child = flatten(head.child);
            head.child = null;

            Node childTail = tail(child);
            head.next = child;
            if (child != null) child.prev = head;

            childTail.next = flatten(next);
            if (next != null) next.prev = childTail;
        }else{
            flatten(head.next);
        }
        return head;
    }
    private Node tail(Node head){
        Node tail = head;
        while(tail.next != null){
            tail = tail.next;
        }
        return tail;
    }
    class Node {
        public int val;
        public Node prev;
        public Node next;
        public Node child;
    };
    //1376. 通知所有员工所需的时间
    public int numOfMinutes(int n, int headID, int[] manager, int[] informTime) {
        //思路：用一个数组，表示员工的下级都有谁
        ArrayList<Integer>[] subs = new ArrayList[n];
        for(int i = 0;i < n;i ++) subs[i] = new ArrayList();
        for(int i = 0;i < manager.length;i ++){
            if (manager[i] >= 0 )subs[manager[i]].add(i);
        }
        //从headID开始广度优先遍历，需要记录到达每一个员工的时间
        int[] times = new int[n];
        Stack<Integer> stack = new Stack();
        stack.push(headID);
        while(stack.size() > 0){
            //从中弹出一个人
            int cur = stack.pop();
            //通知到它的时候已经花费的时间为
            int alreadyTime = times[cur];
            //它的下级
            for(Integer sub : subs[cur]){
                times[sub] = alreadyTime + informTime[cur];
                stack.push(sub);
            }
        }
        int res = 0;
        for(int i = 0;i < n;i ++){
            res = Math.max(res,times[i]);
        }
        return res;
    }
    //49. 字母异位词分组
    public List<List<String>> groupAnagrams(String[] strs) {
        //hashmap ,key 为字母排序后的
        Map<String,List<String>> map = new HashMap();
        for(String str: strs){
            String sortedStr =sortStr(str);
            if(map.containsKey(sortedStr)){
                List<String> item = map.get(sortedStr);
                item.add(str);
            }else{
                List<String> item = new ArrayList();
                item.add(str);
                map.put(sortedStr,item);
            }
        }
        List<List<String>> res = new ArrayList();
        for(Map.Entry<String,List<String>> entry:map.entrySet()){
            res.add(entry.getValue());
        }
        return res;
    }

    private String sortStr(String str){
        byte[] bytes = str.getBytes();
        Arrays.sort(bytes);
        return new String(bytes);
    }
}