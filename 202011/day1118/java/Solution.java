import java.util.Stack;

/**
 * @author tanhao
 * @date 2020/11/18 08:59
 */
public class Solution {
    //剑指 Offer 36. 二叉搜索树与双向链表
    public Node treeToDoublyList(Node root) {
        if (root == null) {
            return null;
        }
        Node head = null;
        //思路：中序遍历二叉搜索树即为有序，
        //非递归中序遍历
        Stack<Node> stack = new Stack<>();
        //左子树进栈
        while (root != null) {
            stack.push(root);
            root = root.left;
        }

        Node pre = null;
        while (!stack.isEmpty()) {
            Node pop = stack.pop();
            if (pre != null){
                pre.right = pop;
                pop.left = pre;
            }else{
                head = pop;
            }
            pre = pop;
            Node popRight = pop.right;
            System.out.println(pop.val);

            while (popRight != null) {
                stack.push(popRight);
                popRight = popRight.left;
            }

        }

        //最后结束的时候，最后一个需要指向头，头的left指向最后一个
        head.left = pre;
        pre.right = head;

        return head;
    }
}

class Node {
    public int val;
    public Node left;
    public Node right;

    public Node() {
    }

    public Node(int _val) {
        val = _val;
    }

    public Node(int _val, Node _left, Node _right) {
        val = _val;
        left = _left;
        right = _right;
    }
}