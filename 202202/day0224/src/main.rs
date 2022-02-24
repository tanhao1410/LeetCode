fn main() {
    println!("Hello, world!");
}

impl Solution {
    //2. 两数相加
    pub fn add_two_numbers(mut l1: Option<Box<ListNode>>, mut l2: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        //每一次创建一个
        let mut head = Some(Box::new(ListNode::new(0)));
        let mut p = head.as_ref().unwrap().as_ref();
        let mut flag = 0;
        while l1.is_some() && l2.is_some() {
            let val = l1.as_ref().unwrap().val + l2.as_ref().unwrap().val + flag;
            flag = val / 10;
            let new_node = Some(Box::new(ListNode::new(val % 10)));
            unsafe {
                let p_ptr = p as *const ListNode as *mut ListNode;
                (*p_ptr).next = new_node;
                p = (*p_ptr).next.as_ref().unwrap().as_ref();
            }
            l1 = l1.as_mut().unwrap().next.take();
            l2 = l2.as_mut().unwrap().next.take();
        }
        let mut l3 = None;
        if l2.is_some() {
            l3 = l2;
        }
        if l1.is_some() {
            l3 = l1;
        }
        while let Some(mut node) = l3 {
            let val = node.val + flag;
            flag = val / 10;
            let new_node = Some(Box::new(ListNode::new(val % 10)));
            unsafe {
                let p_ptr = p as *const ListNode as *mut ListNode;
                (*p_ptr).next = new_node;
                p = (*p_ptr).next.as_ref().unwrap().as_ref();
            }
            l3 = node.next.take();
        }
        if flag > 0 {
            let new_node = Some(Box::new(ListNode::new(flag)));
            unsafe {
                let p_ptr = p as *const ListNode as *mut ListNode;
                (*p_ptr).next = new_node;
            }
        }
        head.as_mut().unwrap().next.take()
    }

    //20. 有效的括号
    pub fn is_valid(s: String) -> bool {
        let mut stack = vec![];
        for s in s.chars() {
            match s {
                '(' | '[' | '{' => stack.push(s),
                _ => {
                    if stack.is_empty() {
                        let pop = stack.pop().unwrap();
                        if (pop == '(' && s != ')') || (pop == '{' && s != '}') || (pop == '[' && s != ']') {
                            return false;
                        }
                    } else {
                        return false;
                    }
                }
            }
        }
        stack.len() == 0
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