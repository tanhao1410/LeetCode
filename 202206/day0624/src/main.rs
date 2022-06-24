fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //515. 在每个树行中找最大值
    pub fn largest_values(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        use std::collections::VecDeque;
        let mut queue = VecDeque::new();
        let mut res = vec![];
        if let Some(node) = root {
            queue.push_back(node);
            while !queue.is_empty() {
                let queue_len = queue.len();
                let mut max = i32::MIN;
                for _ in 0..queue_len {
                    let head = queue.pop_front().unwrap();
                    let head_ref = head.borrow();
                    max = max.max(head_ref.val);
                    head_ref.left.as_ref().map_or((), |e| queue.push_back(e.clone()));
                    head_ref.right.as_ref().map_or((), |e| queue.push_back(e.clone()));
                }
                res.push(max);
            }
        }
        res
    }
}

#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}

use std::rc::Rc;
use std::cell::RefCell;