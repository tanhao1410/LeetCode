fn main() {
    println!("Hello, world!");
}


impl Solution {
    //1022. 从根到叶的二进制数之和
    pub fn sum_root_to_leaf(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        Self::sum_root_to_leaf_pre(root.as_ref().unwrap(), 0)
    }
    fn sum_root_to_leaf_pre(root: &Rc<RefCell<TreeNode>>, pre_value: i32) -> i32 {
        let root = root.borrow();
        let pre_value = pre_value * 2 + root.val;
        if root.left.is_none() && root.right.is_none() {
            return pre_value;
        }
        let mut res = 0;
        if root.left.is_some() {
            res += Self::sum_root_to_leaf_pre(root.left.as_ref().unwrap(), pre_value);
        }
        if root.right.is_some() {
            res += Self::sum_root_to_leaf_pre(root.right.as_ref().unwrap(), pre_value)
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