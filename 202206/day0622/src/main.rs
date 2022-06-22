fn main() {
    println!("Hello, world!");
}

impl Solution {
    //1315. 祖父节点值为偶数的节点和
    pub fn sum_even_grandparent(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let sum_son = |root: Rc<RefCell<TreeNode>>| {
            let mut res = 0;
            let node_ref = root.borrow();
            res += node_ref.left.as_ref().map_or(0, |e| e.borrow().val);
            res += node_ref.right.as_ref().map_or(0, |e| e.borrow().val);
            res
        };

        let sum_grand_son = |root: Rc<RefCell<TreeNode>>| {
            let mut res = 0;
            let node_ref = root.borrow();
            if node_ref.left.is_some() {
                res += sum_son(node_ref.left.as_ref().unwrap().clone());
            }
            if node_ref.right.is_some() {
                res += sum_son(node_ref.right.as_ref().unwrap().clone());
            }
            res
        };

        let mut res = 0;
        if let Some(node) = root {
            if node.borrow().val % 2 == 0 {
                res += sum_grand_son(node.clone());
            }

            res += Self::sum_even_grandparent(node.borrow_mut().left.take());
            res += Self::sum_even_grandparent(node.borrow_mut().right.take());
        }
        res
    }

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