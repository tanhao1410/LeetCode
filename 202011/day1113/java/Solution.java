import java.util.List;

/**
 * @author tanhao
 * @date 2020/11/13 09:37
 */
public class Solution {

    public ListNode oddEvenList(ListNode head) {
        //链表为空或链表长度为1,2时直接返回
        if (head == null || head.next == null || head.next.next == null){
            return head;
        }
        //下一个奇数节点
        ListNode nextOdd = head.next.next;
        ListNode tailOdd = head;
        ListNode firstEven = head.next;
        ListNode tailEven = head.next;
        while (nextOdd != null){
            tailEven.next = nextOdd.next;
            tailOdd.next = nextOdd;
            nextOdd.next = firstEven;

            tailOdd = nextOdd;
            tailEven = tailEven.next;
            if (tailEven == null){
                break;
            }
            nextOdd = tailEven.next;
        }

        return head;
    }

    public class ListNode {
        int val;
        ListNode next;

        ListNode() {
        }

        ListNode(int val) {
            this.val = val;
        }

        ListNode(int val, ListNode next) {
            this.val = val;
            this.next = next;
        }
    }
}
