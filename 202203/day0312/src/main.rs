fn main() {
    println!("Hello, world!");
}


impl Solution {
    //61. 旋转链表
    pub fn rotate_right(mut head: Option<Box<ListNode>>, k: i32) -> Option<Box<ListNode>> {
        if head.is_none() {
            return None;
        }
        let list_len = |mut head: &Option<Box<ListNode>>| {
            let mut res = 0;
            while let Some(node) = head {
                res += 1;
                head = &node.next;
            }
            res
        };
        let list_append = |mut head: Option<Box<ListNode>>, tail: Option<Box<ListNode>>| {
            let mut p = head.as_mut();
            while let Some(node) = p {
                if node.next.is_none() {
                    node.next = tail;
                    return head;
                }
                p = node.next.as_mut();
            }
            tail
        };
        let len = list_len(&head);
        let mut k = len - (k % len);
        if k == 0 {
            return head;
        }
        let mut p = head.as_mut().unwrap();
        while k > 1 {
            p = p.next.as_mut().unwrap();
            k -= 1;
        }
        //需要断开
        list_append(p.next.take(), head)
    }
    //860. 柠檬水找零
    pub fn lemonade_change(bills: Vec<i32>) -> bool {
        //优先找零10元的。
        let mut five = 0;
        let mut ten = 0;
        for i in bills {
            match i {
                5 => five += 1,
                10 => {
                    five -= 1;
                    ten += 1;
                }
                _ => {
                    //需要找零15元
                    if ten > 0 {
                        ten -= 1;
                        five -= 1;
                    } else {
                        five -= 3;
                    }
                }
            }
            if five < 0 {
                return false;
            }
        }
        true
    }
}

struct Solution;

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

use std::collections::BinaryHeap;

//1845. 座位预约管理系统
struct SeatManager {
    queue: BinaryHeap<i32>,
    n: i32,
}

impl SeatManager {
    fn new(n: i32) -> Self {
        let mut queue = BinaryHeap::new();
        for i in 1..=n {
            queue.push(i);
        }
        Self { queue, n }
    }

    fn reserve(&mut self) -> i32 {
        self.n - self.queue.pop().unwrap() + 1
    }

    fn unreserve(&mut self, seat_number: i32) {
        self.queue.push(self.n + 1 - seat_number);
    }
}