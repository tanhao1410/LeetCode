class Solution {
    //707. 设计链表
    class MyLinkedList {

        Node head;
        Node tail;
        int  len = 0;

        class Node{
            public int val;
            public Node next;
            public Node pre;

            public Node(int val){
                this.val = val;
            }
        }

        public MyLinkedList() {

        }

        public int get(int index) {
            if (index >= this.len) return -1;
            Node node = this.head;
            for(int i = 0;i < index;i ++){
                node = node.next;
            }
            return node.val;
        }

        public void addAtHead(int val) {
            Node node = new Node(val);
            node.next = this.head;
            if(this.head != null){
                this.head.pre = node;
            }else{
                this.tail = node;
            }
            this.head = node;
            this.len ++;
        }

        public void addAtTail(int val) {
            this.len ++;
            Node node = new Node(val);
            if(this.tail == null){
                this.head = node;
                this.tail = node;
            }else{
                node.pre = this.tail;
                this.tail.next = node;
                this.tail = node;
            }
        }

        public void addAtIndex(int index, int val) {
            if(index == this.len){
                addAtTail(val);
            }else if (index == 0){
                addAtHead(val);
            }else if (index < this.len){
                //先找index -1个位置
                Node node = this.head;
                for(int i = 0;i < index - 1;i ++){
                    node = node.next;
                }

                Node newNode = new Node(val);
                Node next = node.next;

                node.next = newNode;
                newNode.pre = node;

                newNode.next = next;
                next.pre = newNode;

                this.len ++;
            }
        }

        public void deleteAtIndex(int index) {
            if(index < this.len){
                //只有一个元素的情况
                if (this.len == 1){
                    this.head = null;
                    this.tail = null;
                }else if (index == 0){
                    this.head = this.head.next;
                    this.head.pre = null;
                }else if (index == this.len - 1){
                    this.tail = this.tail.pre;
                    this.tail.next = null;
                }else{
                    //删除中间的节点。
                    Node node = this.head;
                    for(int i = 0;i < index - 1;i ++){
                        node = node.next;
                    }
                    Node delNode = node.next;
                    node.next = delNode.next;
                    node.next.pre = node;

                    delNode.next = null;
                    delNode.pre = null;
                }
                this.len --;
            }
        }
    }

    //599. 两个列表的最小索引总和
    public String[] findRestaurant(String[] list1, String[] list2) {
        Map<String,Integer> map = new HashMap();
        for(int i = 0;i < list1.length;i ++){
            map.put(list1[i],i);
        }
        int min = list1.length + list2.length;
        int count = 0;
        for(int i = 0;i < list2.length;i ++){
            if (map.containsKey(list2[i])){
                if (min > i + map.get(list2[i])){
                    min = i + map.get(list2[i]);
                    count = 1;
                }else if (min == i + map.get(list2[i])){
                    count ++;
                }
            }
        }
        String[] res = new String[count];
        int index = 0;
        for(int i = 0;i < list2.length;i ++){
            if (map.containsKey(list2[i]) && i + map.get(list2[i]) == min){
                res[index++] = list2[i];
            }
        }
        return res;
    }
}