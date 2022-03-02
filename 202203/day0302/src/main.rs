fn main() {
    println!("Hello, world!");
}


impl Solution {
    //199. 二叉树的右视图
    pub fn right_side_view(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        use std::collections::VecDeque;
        let mut res = vec![];
        let mut queue = VecDeque::new();
        if root.is_some() {
            queue.push_back(root.as_ref().unwrap().clone());
        }
        while !queue.is_empty() {
            let len = queue.len();
            for i in 0..len {
                let cur = queue.pop_front().unwrap();
                if i == len - 1 {
                    res.push(cur.borrow().val);
                }
                if cur.borrow().left.is_some() {
                    queue.push_back(cur.borrow().left.as_ref().unwrap().clone());
                }
                if cur.borrow().right.is_some() {
                    queue.push_back(cur.borrow().right.as_ref().unwrap().clone());
                }
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

struct Solution;

use std::rc::Rc;
use std::cell::RefCell;