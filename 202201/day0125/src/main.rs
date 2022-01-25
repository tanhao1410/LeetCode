fn main() {
    println!("Hello, world!");
    let mut head = ListNode::new(1);
    let node1 = ListNode::new(2);
    head.next = Some(Box::new(node1));
    let head = Some(Box::new(head));
    println!("{:?}", Solution::middle_node(head));
}

#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode {
            next: None,
            val,
        }
    }
}

struct Solution;

impl Solution {
    //876. 链表的中间结点
    pub fn middle_node(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        //一个一次走两步，一个一次走一步，快的到终点时，慢的即中间节点（如果，快的只走一步就到终点，则慢的再走一步）
        let mut fast = head.as_ref().unwrap();
        let mut slow = head.as_ref().unwrap();

        //快的先走一步，如果没有了，则慢的就是中间
        while let Some(next) = fast.next.as_ref() {
            //慢的走一步
            slow = slow.next.as_ref().unwrap();
            //快的走两步
            if let Some(next) = next.next.as_ref() {
                fast = next;
            } else {
                fast = next;
            }
        }
        Some(slow.clone())
    }
}

//1688. 比赛中的配对次数
pub fn number_of_matches(mut n: i32) -> i32 {
    let mut res = 0;
    while n > 1 {
        res += n / 2;
        n = n - n / 2;
    }
    res
}
