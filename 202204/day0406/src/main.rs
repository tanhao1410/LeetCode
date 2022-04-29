fn main() {
    println!("Hello, world!");
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

impl Solution {
    //951. 翻转等价二叉树
    pub fn flip_equiv(root1: &Option<Rc<RefCell<TreeNode>>>, root2: &Option<Rc<RefCell<TreeNode>>>) -> bool {
        if root1.is_none() {
            return root2.is_none();
        }
        if root2.is_none() {
            return root1.is_none();
        }
        if let Some(node1) = root1.as_ref() {
            if let Some(node2) = root2.as_ref() {
                if node1.borrow().val == node2.borrow().val {
                    return (Self::flip_equiv(&node1.borrow().left, &node2.borrow().left)
                        && Self::flip_equiv(&node1.borrow().right, &node2.borrow().right))
                        || (Self::flip_equiv(&node1.borrow().right, &node2.borrow().left)
                        && Self::flip_equiv(&node1.borrow().left, &node2.borrow().right));
                }
            }
        }
        false
    }
}