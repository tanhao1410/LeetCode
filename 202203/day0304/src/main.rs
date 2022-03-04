fn main() {
    println!("Hello, world!");
}

use std::rc::Rc;
use std::cell::RefCell;

impl Solution {
    //236. 二叉树的最近公共祖先
    pub fn lowest_common_ancestor(root: Option<Rc<RefCell<TreeNode>>>, p: Option<Rc<RefCell<TreeNode>>>, q: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        //思路：如果根与p,或q相等，则返回root。否则，可能在于root的左子树或右子树。如果左与右都没有，说明左右各一样，返回root。否则返回对应的即可。
        if root.is_none() {
            return None;
        }
        let root_val = root.as_ref().unwrap().borrow().val;
        if root_val == p.as_ref().unwrap().borrow().val || root_val == q.as_ref().unwrap().borrow().val {
            return root;
        }
        let pp = Some(p.as_ref().unwrap().clone());
        let qq = Some(q.as_ref().unwrap().clone());
        let left = Self::lowest_common_ancestor(root.as_ref().unwrap().borrow_mut().left.take(), pp, qq);
        let right = Self::lowest_common_ancestor(root.as_ref().unwrap().borrow_mut().right.take(), p, q);
        if left.is_some() && right.is_some() {
            return root;
        } else if left.is_none() {
            return right;
        } else {
            return left;
        }
    }
}

struct Solution;

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