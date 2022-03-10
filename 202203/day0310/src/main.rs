fn main() {
    println!("Hello, world!");
}

impl Solution {
    //445. 两数相加 II
    pub fn add_two_numbers(l1: Option<Box<ListNode>>, l2: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        //需要从最低位开始加起来。
        let list_to_vec = |mut list: Option<Box<ListNode>>| {
            let mut res = vec![];
            while let Some(mut node) = list {
                res.push(node.val);
                list = node.next.take();
            }
            res
        };
        let vec_to_list = |v: Vec<i32>| {
            let mut root = ListNode::new(0);
            let mut p = &mut root;
            for num in v.into_iter().rev() {
                let mut node = ListNode::new(num);
                p.next = Some(Box::new(node));
                p = p.next.as_mut().unwrap().as_mut();
            }
            root.next.take()
        };
        let v1 = list_to_vec(l1);
        let v2 = list_to_vec(l2);
        let mut index = 0;
        let mut res = vec![];
        let mut flag = 0;
        while index < v1.len() || index < v2.len() {
            let mut sum = flag;
            if index < v1.len() && index < v2.len() {
                sum += v1[v1.len() - 1 - index] + v2[v2.len() - 1 - index];
            } else if index < v1.len() {
                sum += v1[v1.len() - 1 - index];
            } else {
                sum += v2[v2.len() - 1 - index];
            }
            res.push(sum % 10);
            flag = sum / 10;
            index += 1;
        }
        if flag > 0 {
            res.push(1);
        }
        vec_to_list(res)
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
