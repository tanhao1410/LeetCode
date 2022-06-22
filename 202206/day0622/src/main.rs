fn main() {
    println!("Hello, world!");
}

impl Solution {
    //513. 找树左下角的值
    pub fn find_bottom_left_value(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let mut queue = std::collections::VecDeque::new();
        queue.push_back(root.unwrap());
        let mut res = 0;
        while !queue.is_empty() {
            let queue_len = queue.len();
            res = queue.get(0).unwrap().borrow().val;
            for _ in 0..queue_len {
                let pop = queue.pop_front().unwrap();
                if pop.borrow().left.is_some() {
                    queue.push_back(pop.borrow().left.as_ref().unwrap().clone());
                }
                if pop.borrow().right.is_some() {
                    queue.push_back(pop.borrow().right.as_ref().unwrap().clone());
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

use std::rc::Rc;
use std::cell::RefCell;

struct Solution;