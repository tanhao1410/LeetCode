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
    //19. 删除链表的倒数第 N 个结点
    pub fn remove_nth_from_end(mut head: Option<Box<ListNode>>, n: i32) -> Option<Box<ListNode>> {
        //最原始思路，遍历一遍，得到链表长度。如何一次遍历即可呢。双指针，一个在前，一个在后，在前的先走n步，
        let mut fast = head.as_ref().unwrap();
        let mut slow = head.as_ref().unwrap();
        //快的先走 n步，然后，慢的指向的就是要删的。
        for _ in 0..n {
            if let Some(next) = fast.next.as_ref() {
                fast = next;
            } else {
                //说明走到头了，即删除的是第一个元素。或长度本来就是1
                return head.unwrap().next.take();
            }
        }

        while let Some(next) = fast.next.as_ref() {
            fast = next;
            slow = slow.next.as_ref().unwrap();
        }
        // 采用unsafe来删除呢？
        let remove = slow.next.as_ref().unwrap() as *const _ as *mut Box<ListNode>;
        let slow = slow as *const _ as *mut Box<ListNode>;
        unsafe {
            let remain = (*remove).next.take();
            (*slow).next = remain;
        }
        head
    }

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
